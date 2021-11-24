package model

import (
	"fmt"
	"goweb/utils"
)

// User结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// 添加User的方法一
func (user *User) AddUser() error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"

	// 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
		return err
	}
	// 执行
	_, err2 := inStmt.Exec("admin", "123456", "admin@atguigu.com")
	if err2 != nil {
		fmt.Println("执行异常:", err2)
		return err
	}
	return nil
}

// 添加User的方法二
func (user *User) AddUser2() error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"

	// 执行
	_, err := utils.Db.Exec(sqlStr, "admin2", "6766666", "admin2@sina.com")
	if err != nil {
		fmt.Println("执行异常:")
		return err
	}
	return nil
}
