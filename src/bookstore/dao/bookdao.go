package dao

import (
	"bookstore/model"
	"bookstore/utils"
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