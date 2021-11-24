package model

import "web01_db/utils"

// User结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// 添加User的方法
func (user *User) AddUser() error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"

}
