package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	fmt.Println("测试订单相关函数")
	t.Run("测试添加订单和订单项", testAddorder)
}

func testAddorder(t *testing.T) {
	// 生成订单号
	orderID := "13811118888"
	// 创建订单
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  time.Now(),
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      1,
	}
	// 创建订单项
	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  300,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   300,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID,
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID,
	}
	// 保存订单
	AddOrder(order)
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
}
