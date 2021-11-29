package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"strconv"
)

// 获取所有图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books"
	// 执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		// 给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		// 将book添加到books中
		books = append(books, book)
	}
	return books, nil
}

// 添加图书
func AddBook(b *model.Book) error {
	sqlStr := "INSERT INTO books(title, author, price, sales, stock, img_path) values(?,?,?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// 根据图书id删除图书
func DeleteBook(bookID string) error {
	sqlStr := "delete from books where id = ?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

// 根据图书id获取图书
func GetBookById(bookID string) (*model.Book, error) {
	sqlStr := "select id, title, author, price, sales, stock, img_path from books where id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, bookID)
	// 创建Book
	book := &model.Book{}
	// 为book中的字段赋值
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// 根据图书id更新图书
func UpdateBook(b *model.Book) error {
	sqlStr := "update books set title=?, author=?, price=?,sales=?, stock=? where id = ?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

// 获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	// 把页码转换为int64类型
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	// 总记录数
	sqlStr := "select count(*) from books"
	// 接收总记录数的变量
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	// 设置每页只显示4条记录
	var pageSize int64 = 8
	// 设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	// 获取当前页中的图书
	sqlStr2 := "select id, title, author, price, sales, stock, img_path from books limit ?, ?"
	// 执行
	rows, _ := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	// 创建page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
