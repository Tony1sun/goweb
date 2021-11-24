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

// 根据用户id查询
func (user *User) GetUserById() (*User, error) {
	sqlStr := "select id, username, password, email from users where id = ?"

	//执行
	row := utils.Db.QueryRow(sqlStr, user.ID)
	// 声明
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

// 获取所有记录
func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select id, username, password, email from users"

	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	// 创建User切片
	var users []*User
	for rows.Next() {
		// 声明
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, nil
}
