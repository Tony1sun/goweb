package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestSession(t *testing.T) {
	fmt.Println("测试Session中的方法")
	// t.Run("添加Session:", testAddsession)
	// t.Run("删除Session:", testDeleteSession)
	t.Run("获取Session:", testGetSession)
}

func testAddsession(t *testing.T) {
	sessions := &model.Session{
		SessionID: "13838381438",
		UserName:  "admin",
		UserID:    1,
	}
	AddSession(sessions)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("13838381438")
}

func testGetSession(t *testing.T) {
	sess, _ := GetSession("fa70d500-2bfe-42d3-4ef0-08d34cfa6356")
	fmt.Println("Session的信息是：", sess)
}
