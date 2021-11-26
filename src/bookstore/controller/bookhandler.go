package controller

import (
	"bookstore/dao"
	"html/template"
	"net/http"
)

// 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// 获取图书
	books, _ := dao.GetBooks()
	// 解析
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	// 执行
	t.Execute(w, books)
}
