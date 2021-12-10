package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"net/http"
	"time"
)

// 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	// 获取session
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 获取购物车
	cart, _ := dao.GetCartByUserID(userID)
	// 生成订单号
	orderID := utils.CreateUUID()
	// 生成订单的时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	// 创建order
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	// 将订单保存到数据库
	dao.AddOrder(order)
	// 将订单项保存到数据库
	// 获取购物车中的购物项
	cartItems := cart.CartItems
	// 遍历每一个购物项
	for _, v := range cartItems {
		// 创建OrderItem
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderID,
		}
		// 将订单项保存到数据库
		dao.AddOrderItem(orderItem)
		// 更新当前购物项中图书的库存和销量
		book := v.Book
		book.Stock = book.Stock - int(v.Count)
		book.Sales = book.Sales + int(v.Count)
		// 更新图书信息
		dao.UpdateBook(book)
	}
	// 清空购物车
	dao.DeleteCartByCartID(cart.CartID)
	// 将订单号设置到session中
	session.Order = order
	// 解析
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	// 执行
	t.Execute(w, session)
}

// 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	// 获取所有订单
	orders, _ := dao.GetOrders()
	// 解析
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	// 执行
	t.Execute(w, orders)
}

// 获取订单订单对应的订单项
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	// 获取订单号
	orderID := r.FormValue("orderId")
	// 根据订单号获取所有订单项
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	// 解析模版
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	// 执行
	t.Execute(w, orderItems)
}

// 获取我的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	// 获取session
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 获取用户的所有订单
	orders, _ := dao.GetMyOrders(userID)
	// 将订单设置到session中
	session.Orders = orders
	// 解析
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)
}

// 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	// 获取要发货的订单号
	orderID := r.FormValue("orderId")
	// 更新订单状态
	dao.UpdateOrderState(orderID, 1)
	// 查询一下所有订单
	GetOrders(w, r)
}

// 收货
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	// 获取
	// 获取要收货的订单号
	orderID := r.FormValue("orderId")
	// 更新订单状态
	dao.UpdateOrderState(orderID, 2)
	// 查询我的订单
	GetMyOrders(w, r)
}
