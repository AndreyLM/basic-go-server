package models

// User - user
type User struct {
	ID       int    `xorm:"'id' pk autoincr" json:"id" `
	Login    string `conform:"trim" validate:"required" xorm:"'login'" json:"login"`
	Email    string `conform:"email" validate:"required,email" xorm:"'email'" json:"email"`
	Password string `validate:"required,gte=6" xorm:"'password'" json:"password"`
}

// TableName - table name for xorm
func (u *User) TableName() string {
	return "users"
}
