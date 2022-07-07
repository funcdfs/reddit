package models

// User in models define the integral parameters in the user
type User struct {
	UserName    string `db:"user_name"`
	Password    string `db:"password"`
	UserID      uint64 `json:"user_id,string" db:"user_id"`
	AccessToken string
	// RefreshToken   string
}
