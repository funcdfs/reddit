package models

type User struct {
	UserName string `db:"user_name"`
	Password string `db:"password"`
	UserID   int64  `db:"user_id"`
}
