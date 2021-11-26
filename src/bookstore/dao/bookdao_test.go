package dao

import (
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的方法")
	t.Run("书籍列表:", testgetBook)
}

func testgetBook(t *testing.T) {
	books, _ := GetBooks()
	// 遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%d本图书的信息是: %v\n", k+1, v)
	}
}
