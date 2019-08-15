package forms

import (
	"github.com/andreylm/basic-go-server.git/pkg/server/models"
)

// Auth - form for registration new user
type Auth struct {
	models.User
	ConfimPassword string `validate:"required,gte=6" json:"confirm_password"`
}

// Login - form for logginig
type Login struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}
