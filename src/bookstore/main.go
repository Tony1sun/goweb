package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	// 设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main", controller.GetPageBooksByPrice)

	// 登录
	http.HandleFunc("/login", controller.Login)
	// 注销
	http.HandleFunc("/logout", controller.Logout)
	// 注册
	http.HandleFunc("/regist", controller.Regist)
	// 通过Ajax请求验证用户名是否可用
	http.HandleFunc("/CheckUserName", controller.CheckUserName)
	// 获取所有图书
	// http.HandleFunc("/getBooks", controller.GetBooks)
	// 获取分页后的图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)

	// 添加或更新图书
	http.HandleFunc("/updateOraddBook", controller.UpdateOrAddBook)
	// 删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	// 去添加或更新图书页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	// 更新图书
	// http.HandleFunc("/updateBook", controller.UpdateBook)
	// 添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	// 获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	// 清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	// 删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	// 更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	http.ListenAndServe(":8080", nil)
}
