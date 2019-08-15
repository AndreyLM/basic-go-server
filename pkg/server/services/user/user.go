package user

import (
	"github.com/andreylm/basic-go-server.git/pkg/db"
	"github.com/andreylm/basic-go-server.git/pkg/passwords"
	"github.com/andreylm/basic-go-server.git/pkg/server/handlers/forms"
	"github.com/andreylm/basic-go-server.git/pkg/server/models"
	DefaultValidator "github.com/andreylm/basic-go-server.git/pkg/validator"
)

// Exist - checks if user exist
func Exist(db db.DB, u *models.User) (bool, error) {
	exist, err := db.Exists(&models.User{Login: u.Login})
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}

	return db.Exists(&models.User{Email: u.Email})
}

// CreateUser - user validations
func CreateUser(db db.DB, u *models.User) error {
	var err error
	u.Password, err = passwords.Encrypt(u.Password)
	if err != nil {
		return err
	}

	return db.Store(u)
}

// GetUserByLoginForm - gets user from login form
func GetUserByLoginForm(db db.DB, model *forms.Login) (*models.User, error) {
	user := &models.User{}

	// Check if field is email if error then user entered login otherwise email
	if err := DefaultValidator.ValidateVar(model.User, "email"); err != nil {
		user.Login = model.User
	} else {
		user.Email = model.User
	}
	err := db.Get(user)

	return user, err
}

// GetUser - gets user from login form
func GetUser(db db.DB, user *models.User) (err error) {
	err = db.Get(user)
	return
}

// ValidateAuthForm - validates user model
func ValidateAuthForm(u *forms.Auth) (errors map[string]string) {
	if errors = DefaultValidator.Validate(u); len(errors) > 0 {
		return
	}

	if u.Password != u.ConfimPassword {
		errors["ConfirmPassword"] = "Confirm password and password must be the same"
	}

	return
}

// ValidateLoginForm - validates user model
func ValidateLoginForm(form *forms.Login) (errors map[string]string) {
	errors = DefaultValidator.Validate(form)
	return
}
