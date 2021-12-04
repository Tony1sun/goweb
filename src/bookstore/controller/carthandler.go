package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"net/http"
)

// 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	// 获取要添加的图书id
	bookID := r.FormValue("bookId")
	// 根据图书id获取图书信息
	book, _ := dao.GetBookById(bookID)
	// 判断是否登陆
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 判断购物车中是否有当前用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		// 当前用户已经有购物车

	} else {
		// 当前用户还没有购物车，创建一个购物车并添加到数据库
		// 1.创建购物车
		// 生成uuid
		cartID := utils.CreateUUID()
		cart := &model.Cart{
			CartID: cartID,
			UserID: int64(userID),
		}
		// 2.创建购物车中的购物项
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartID: cartID,
		}
		// 将购物项添加到切片中
		cartItems = append(cartItems, cartItem)
		// 3.将切片设置到cart中
		cart.CartItems = cartItems
		// 将购物车cart保存到数据库
		dao.AddCart(cart)
	}
}