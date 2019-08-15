package auth

import (
	"log"
	"net/http"

	"github.com/andreylm/basic-go-server.git/pkg/helpers"
	"github.com/andreylm/basic-go-server.git/pkg/server/models"

	"github.com/andreylm/basic-go-server.git/pkg/db"

	"github.com/andreylm/basic-go-server.git/pkg/jwt"
	UserService "github.com/andreylm/basic-go-server.git/pkg/server/services/user"
)

// CheckAuth - check auth handler
func CheckAuth(DB db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-token")

		if len(token) == 0 {
			http.Error(w, "You must provide token", http.StatusUnauthorized)
			return
		}
		id, err := jwt.IsTokenValid(token)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		user := &models.User{ID: int(id)}
		if err = UserService.GetUser(DB, user); err != nil {
			log.Println(err)
			http.Error(w, "Can't get user", http.StatusUnauthorized)
			return
		}

		user.Password = ""
		helpers.JSONResponse(w, http.StatusOK, map[string]interface{}{"user": user})
	}
}
