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
	sess, _ := GetSession("bab7a01d-c686-447b-52ac-2c00997a812e")
	fmt.Println("Session的信息是：", sess)
}
