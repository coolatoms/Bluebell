package logic

import (
	"studyWeb/Bluebell/dao/mysql"
	"studyWeb/Bluebell/models"
	"studyWeb/Bluebell/pkg/jwt"
	"studyWeb/Bluebell/pkg/snowflake"
)

// SingUp 注册逻辑
func SingUp(p *models.ParamSignUp) (err error) {
	//1，判断用户名是否存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	//2，生成UID
	userID := snowflake.GenID()
	//3.构造用户实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//4，保存进数据库s
	return mysql.InsertUser(user)
}

// Login 登录逻辑
func Login(login *models.ParamLogin) (user *models.User, err error) {
	//
	user = &models.User{
		Username: login.Username,
		Password: login.Password,
	}
	//登录查询 , 传递指针能拿到userID
	err = mysql.Login(user)
	if err != nil {
		return nil, err
	}

	//生成JWTToken
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
