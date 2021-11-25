package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 检查用户名和密码
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	// 插入
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

// 检查用户名
func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username = ?"
	// 插入
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

// 插入用户信息
func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"
	// 插入
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
