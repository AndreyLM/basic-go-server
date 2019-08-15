package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andreylm/basic-go-server.git/pkg/jwt"
	"github.com/andreylm/basic-go-server.git/pkg/passwords"

	"github.com/andreylm/basic-go-server.git/pkg/db"
	"github.com/andreylm/basic-go-server.git/pkg/helpers"
	"github.com/andreylm/basic-go-server.git/pkg/server/handlers/forms"
	UserService "github.com/andreylm/basic-go-server.git/pkg/server/services/user"
)

// Login - login
func Login(db db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginForm := forms.Login{}
		if err := json.NewDecoder(r.Body).Decode(&loginForm); err != nil {
			http.Error(w, "Cannot decode request body", http.StatusBadRequest)
			return
		}

		if errors := UserService.ValidateLoginForm(&loginForm); len(errors) > 0 {
			helpers.JSONResponse(w, http.StatusUnauthorized, errors)
			return
		}

		user, err := UserService.GetUserByLoginForm(db, &loginForm)
		if err != nil {
			log.Println(err)
			helpers.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error getting user"})
			return
		}

		if user.ID == 0 || !passwords.IsValid(user.Password, loginForm.Password) {
			http.Error(w, "Invalid login/email or password", http.StatusUnauthorized)

			// helpers.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid login/email or password"})
			return
		}

		user.Password = ""
		helpers.JSONResponse(w, http.StatusOK, map[string]interface{}{
			"user":  user,
			"token": jwt.GetToken(int64(user.ID)),
		})
	}
}
