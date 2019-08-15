package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andreylm/basic-go-server.git/pkg/jwt"

	"github.com/andreylm/basic-go-server.git/pkg/helpers"

	"github.com/andreylm/basic-go-server.git/pkg/server/handlers/forms"

	"github.com/andreylm/basic-go-server.git/pkg/db"

	UserService "github.com/andreylm/basic-go-server.git/pkg/server/services/user"
)

// Register - register new user
func Register(DB db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userForm := forms.Auth{}
		if err := json.NewDecoder(r.Body).Decode(&userForm); err != nil {
			http.Error(w, "Cannot decode request body", http.StatusUnauthorized)
			return
		}

		if userForm.ID > 0 {
			http.Error(w, "You cannot provide id property", http.StatusUnauthorized)
			return
		}

		if errors := UserService.ValidateAuthForm(&userForm); len(errors) > 0 {
			var res string
			for field, msg := range errors {
				res += fmt.Sprintf("%s: %s. ", field, msg)
			}
			http.Error(w, res, http.StatusUnauthorized)
			return
		}

		exist, err := UserService.Exist(DB, &userForm.User)
		if err != nil {
			http.Error(w, "Cannot check if  user exists", http.StatusBadRequest)
			return
		}

		if exist {
			http.Error(w, "User exists", http.StatusBadRequest)
			return
		}

		if err := UserService.CreateUser(DB, &userForm.User); err != nil {
			log.Println(err)
			http.Error(w, "Cannot create user", http.StatusBadRequest)
			return
		}
		userForm.Password = ""
		helpers.JSONResponse(w, http.StatusOK, map[string]interface{}{
			"user":  userForm.User,
			"token": jwt.GetToken(int64(userForm.ID)),
		})
	}
}
