package logic

import (
	"errors"

	"reddit/dao/mysql"
	"reddit/logger"
	"reddit/models"
	"reddit/pkg/gen"
	"reddit/pkg/jwt"
)

// SignUp logic.Signup use models.ParamSignUp to call mysql.InsertUser
// check if user exists
// get a random user ID use snowflake generator
// use parameters to generate a user
// use mysql.InsertUser to insert this user
func SignUp(p *models.ParamSignUp) (err error) {
	err = mysql.CheckUserExists(p.UserName)
	if err != nil {
		return err // 数据库查询出错
	}
	userID := gen.NewID()
	if err != nil {
		return errors.New("id generation failed")
	}
	u := models.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: p.Password,
	}
	return mysql.InsertUser(&u)
}

// Login use models.ParamLogin to create a new user
// then use mysql.Login to login
// after mysql.Login is done you can get the userId int64
// use the userId and username to generate a jwt token
// return this token and error in GenToken process
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		return "error token", errors.New("login failed: " + err.Error())
	}
	logger.Blue("login successful!!!!!!")
	// generate jwt token and return to controller
	return jwt.GenToken(uint64(user.UserID), user.UserName)
	// return "", nil
}
