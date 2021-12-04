package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

// 获取所有图书
// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	// 获取图书
// 	books, _ := dao.GetBooks()
// 	// 解析
// 	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
// 	// 执行
// 	t.Execute(w, books)
// }

// 添加图书
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	// 获取图书信息
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")
// 	// 把价格、销量和库存进行转换
// 	fPrice, _ := strconv.ParseFloat(price, 64)
// 	iSales, _ := strconv.ParseInt(sales, 10, 0)
// 	iStock, _ := strconv.ParseInt(stock, 10, 0)
// 	// 创建Book
// 	book := &model.Book{
// 		Title:   title,
// 		Author:  author,
// 		Price:   fPrice,
// 		Sales:   int(iSales),
// 		Stock:   int(iStock),
// 		ImgPath: "/static/img/default.jpg",
// 	}
// 	// 调用添加图书函数
// 	dao.AddBook(book)
// 	// 调用GetBooks函数查询一次数据库
// 	GetBooks(w, r)
// }

// 首页
// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	// 获取页码
// 	pageNo := r.FormValue("pageNo")
// 	if pageNo == "" {
// 		pageNo = "1"
// 	}
// 	// 调用分页函数
// 	page, _ := dao.GetPageBooks(pageNo)
// 	// 解析模板
// 	t := template.Must(template.ParseFiles("views/index.html"))
// 	// 执行
// 	t.Execute(w, page)
// }

// 获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	// 获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		// 调用分页函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		// 调用 获取带分页和价格的图书信息 函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		// 将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	// 调用IsLogin函数判断是否已经登陆
	flag, session := dao.IsLogin(r)
	if flag {
		// 已经登录,设置page中的IsLogin字段和Username的字段值
		page.IsLogin = true
		page.Username = session.UserName
	}
	// 解析
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行
	t.Execute(w, page)
}

// 获取分页后的图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	// 调用分页函数
	page, _ := dao.GetPageBooks(pageNo)
	// 解析
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	// 执行
	t.Execute(w, page)
}

// 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// 获取要删除的图书id
	bookID := r.FormValue("bookId")
	// 调用删除图书函数
	dao.DeleteBook(bookID)
	// 调用GetBooks函数查询一次数据库
	GetPageBooks(w, r)
}

// 去添加或更新图书页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	// 获取要删除的图书id
	bookID := r.FormValue("bookId")
	// 调用获取图书函数
	book, _ := dao.GetBookById(bookID)
	if book.ID > 0 {
		// 更新图书
		// 解析模版
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		// 执行
		t.Execute(w, book)
	} else {
		// 添加图书
		// 解析模版
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		// 执行
		t.Execute(w, "")
	}

	// 调用GetBooks函数查询一次数据库
	GetPageBooks(w, r)
}

// 添加或更新图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	// 获取图书信息
	ibook := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	// 把价格、销量和库存进行转换
	ibookID, _ := strconv.ParseInt(ibook, 10, 0)
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	// 创建Book
	book := &model.Book{
		ID:      int(ibookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		// 更新
		// 调用更新图书函数
		dao.UpdateBook(book)
	} else {
		// 添加
		dao.AddBook(book)
	}
	// 调用GetBooks函数查询一次数据库
	GetPageBooks(w, r)
}
