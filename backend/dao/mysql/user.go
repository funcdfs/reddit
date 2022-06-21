package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
	"reddit/models"
)

var (
	ErrorUserExists   = errors.New("user already exists: ")
	ErrorUserNotExist = errors.New("user does not exist: ")
	ErrorPassword     = errors.New("password error: ")
	ErrorInsertFailed = errors.New("insert failed: ")
)

const secretKey = "https://github.com/fengwei2002"

// encryptPassword encrypt a password using the provided secret key
func encryptPassword(originPassword string) string {
	h := md5.New()
	h.Write([]byte(secretKey))
	return hex.EncodeToString(h.Sum([]byte(originPassword)))
}

// CheckUserExists checks if the user exists in the database
func CheckUserExists(username string) error {
	sqlStr := `select count(user_id) from user where user_name = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExists
	}
	return nil
}

// InsertUser insert user into mysql database using *models.User object
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	// run sql query
	sqlStr := `insert into user(user_id, user_name, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	if err != nil {
		// log error
		zap.L().Error(err.Error())
		return ErrorInsertFailed
	}
	return nil
}

// Login use *models.User to login
// return a ErrorUserNotExist error if the user is not found
// return a ErrorPassword error if the password is not valid
func Login(user *models.User) (err error) {
	originPassword := user.Password
	sqlStr := `select user_id, user_name, password from user where user_name=?`
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return errors.New("query sql error" + err.Error())
	}
	password := encryptPassword(originPassword)
	if password != user.Password {
		return ErrorPassword
	}
	return
}
