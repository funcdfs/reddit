package logic

import (
	"errors"
	"reddit/dao/mysql"
	"reddit/models"
	"reddit/pkg/gen"
	"reddit/pkg/jwt"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在
	err = mysql.CheckUserExists(p.UserName)
	if err != nil {
		return err // 数据库查询出错
	}
	// 生成 uid
	userID := gen.NewID()
	if err != nil {
		return errors.New("id generation failed")
	}
	// 构造一个 user 实例
	u := models.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: p.Password,
	}

	// 保存进 数据库
	return mysql.InsertUser(&u)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		return "error token", errors.New("login failed: " + err.Error())
	}
	// zap.L().Info("login successful!!!!!!")
	// generate jwt token
	return jwt.GenToken(user.UserID, user.UserName)
}
