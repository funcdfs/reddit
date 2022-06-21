package models

// User in models define the integral parameters in the user
type User struct {
	UserName string `db:"user_name"`
	Password string `db:"password"`
	UserID   int64  `db:"user_id"`
}
