package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"studyWeb/Bluebell/models"
)

const secort = "qiuyunhui"

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) error {
	sqlstr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlstr, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// InsertUser 向数据库中插入用户记录
func InsertUser(user *models.User) (err error) {
	//加密密码
	password := encryptPassword(user.Password)
	//入库
	sqlstr := `insert into user(user_id,username,password)values(?,?,?) `
	_, err = db.Exec(sqlstr, user.UserID, user.Username, password)
	return err
}

// encryptPassword 加密明文密码
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secort))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

// Login 用户登录逻辑
func Login(login *models.User) (err error) {
	opassword := login.Password
	sqlstr := `select user_id,username,password from user where username=?`
	if err := db.Get(login, sqlstr, login.Username); err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(opassword)
	if password != login.Password {
		return ErrorInvalidPassword
	}
	return nil
}

// GetUserByID 通过用户名ID获取用户
func GetUserByID(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, id)
	return
}
