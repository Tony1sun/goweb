package dao

import (
	"bookstore/model"
	"bookstore/utils"
)


func CheckUserNameAndPassword(username string, password string) (*model.User, error) {

	sqlStr := "select id,username,password,email from users where username = ? and pasword = ?"

	row := utils.Db.QueryRow(sqlStr, username, password)


}