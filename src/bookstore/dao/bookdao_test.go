package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的方法")
	// t.Run("书籍列表:", testgetBook)
	// t.Run("添加图书:", testAddBook)
	// t.Run("删除图书:", testDeleteBook)
	// t.Run("获取图书:", testGetBook)
	// t.Run("更新图书后:", testUpdateBook)
	t.Run("分页查询图书:", testGetPageBook)
}

func testgetBook(t *testing.T) {
	books, _ := GetBooks()
	// 遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%d本图书的信息是: %v\n", k+1, v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	// 调用添加图书函数
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	DeleteBook("49")
}

func testGetBook(t *testing.T) {
	book, _ := GetBookById("55")
	fmt.Println("获取的图书信息是:", book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      55,
		Title:   "三国演义哈",
		Author:  "罗贯中",
		Price:   66.88,
		Sales:   10,
		Stock:   1000,
		ImgPath: "/static/img/default.jpg",
	}
	UpdateBook(book)
}

func testGetPageBook(t *testing.T) {
	page, _ := GetPageBooks("10")
	fmt.Println("当前页是:", page.PageNo)
	fmt.Println("总页数是:", page.TotalPageNo)
	fmt.Println("总记录数是:", page.TotalRecord)
	// fmt.Println("当前页中的图书有:", books.PageNo)
	for _, v := range page.Books {
		fmt.Println("图书信息是", v)
	}
}
