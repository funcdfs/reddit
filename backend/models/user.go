package models

type User struct {
	UserName string `db:"user_name"`
	Password string `db:"password"`
	UserID   int8   `db:"user_id"`
}
