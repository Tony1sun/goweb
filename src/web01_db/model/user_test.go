package model

import (
	"fmt"
	"testing"
)

// 测试执行操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始:")
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法:")
	// 通过t.Run()来执行子测试函数
	// t.Run("测试添加用户", testAddUser)
	// t.Run("测试查询一条记录", testGetUserById)
	t.Run("测试查询所有记录", testGetUsers)

}

// 如果函数名不是Test开头，那么该函数默认不执行，我们可以将它设置成为一个子函数
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户: ")
	// user := &User{}

	// user.AddUser()
	// user.AddUser2()
}

// 查询一条记录
func testGetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录:")
	user := &User{
		ID: 1,
	}
	// 调用获取User的方法
	u, _ := user.GetUserById()
	fmt.Println("得到的User信息是:", u)
}

// 查询所有记录
func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录:")
	user := &User{}
	us, _ := user.GetUsers()
	//遍历切片
	for k, v := range us {
		fmt.Printf("第%d个用户是%v:\n", k+1, v)
	}
}