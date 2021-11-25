package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {

	sqlStr := "select id,username,password,email from users where username = ? and pasword = ?"

	row := utils.Db.QueryRow(sqlStr, username, password)
	

	user := &model.User{}
	row.Scan(user.ID, user.Username, user.Password, user.Email)
	return user, nil
}

func CheckUserName(username string) (bool, error) {

	sqlStr := "select id,username,password,email from users where username = ?"

	row := utils.Db.QueryRow(sqlStr, username)

	user := &model.User{}
	row.Scan(user.ID, user.Username, user.Password, user.Email)
	return user, nil
}