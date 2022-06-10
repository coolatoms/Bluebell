package logic

import (
	"errors"
	"studyWeb/Bluebell/dao/mysql"
	"studyWeb/Bluebell/models"
	"studyWeb/Bluebell/pkg/snowflake"
)

func SingUp(p *models.ParamSignUp) (err error) {
	//1，判断用户是否存在
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("用户已存在")
	}
	//2，生成UID
	userID := snowflake.GenID()

	//3，密码加密
	//4，保存进数据库
	mysql.InsertUser()
}
