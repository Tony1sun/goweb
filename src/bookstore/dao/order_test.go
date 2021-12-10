package dao

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	fmt.Println("测试订单相关函数")
	// t.Run("测试添加订单和订单项", testAddorder)
	// t.Run("测试获取订单信息", testGetOrders)
	// t.Run("测试获取所有订单项目", testGetOrderItems)
	// t.Run("测试获取我的订单", testGetMyOrders)
	t.Run("测试发货和收货", testUpdateOrderState)
}

// func testAddorder(t *testing.T) {
// 	// 生成订单号
// 	orderID := "13811118888"
// 	// 生成订单时间
// 	timeStr := time.Now().Format("2006-01-02 15:04:05")
// 	// 创建订单
// 	order := &model.Order{
// 		OrderID:     orderID,
// 		CreateTime:  timeStr,
// 		TotalCount:  2,
// 		TotalAmount: 400,
// 		State:       0,
// 		UserID:      1,
// 	}
// 	// 创建订单项
// 	orderItem := &model.OrderItem{
// 		Count:   1,
// 		Amount:  300,
// 		Title:   "三国演义",
// 		Author:  "罗贯中",
// 		Price:   300,
// 		ImgPath: "/static/img/default.jpg",
// 		OrderID: orderID,
// 	}
// 	orderItem2 := &model.OrderItem{
// 		Count:   1,
// 		Amount:  100,
// 		Title:   "西游记",
// 		Author:  "吴承恩",
// 		Price:   100,
// 		ImgPath: "/static/img/default.jpg",
// 		OrderID: orderID,
// 	}
// 	// 保存订单
// 	AddOrder(order)
// 	AddOrderItem(orderItem)
// 	AddOrderItem(orderItem2)
// }

// func testGetOrders(t *testing.T) {
// 	orders, _ := GetOrders()
// 	for _, v := range orders {
// 		fmt.Println("订单信息是:", v)
// 	}
// }

// func testGetOrderItems(t *testing.T) {
// 	orderItems, _ := GetOrderItemsByOrderID("c6701c00-232f-450f-5d6e-0d80737961f4")
// 	for _, v := range orderItems {
// 		fmt.Println("订单项信息是:", v)
// 	}
// }

// func testGetMyOrders(t *testing.T) {
// 	orders, _ := GetMyOrders(1)
// 	for _, v := range orders {
// 		fmt.Println("我的订单是", v)
// 	}
// }

func testUpdateOrderState(t *testing.T) {
	UpdateOrderState("2a2d1f1b-9241-44cf-51e4-349ce0aeb36d", 1)
}
