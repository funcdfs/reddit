package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"reddit/models"
)

const secretKey = "https://github.com/fengwei2002"

// CheckUserExists checks if the user exists in the database
func CheckUserExists(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

// InsertUser 向数据库中插入一个新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行 sql 语句入库
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.Password)
	return err
}

func encryptPassword(originPassword string) string {
	h := md5.New()
	h.Write([]byte(secretKey))
	return hex.EncodeToString(h.Sum([]byte(originPassword)))
}
