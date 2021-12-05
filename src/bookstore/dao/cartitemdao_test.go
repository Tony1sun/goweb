package dao

import (
	"fmt"
	"testing"
)

func TestCartItems(t *testing.T) {
	fmt.Println("测试购物车相关函数")
	// t.Run("根据book_id对应购物项:", testGetCartItemByBookID)
	// t.Run("根据cart_id对应购物项:", testGetCartItemByCartID)
	t.Run("更新图书id和购物车id以及图书数量更新购物项中图书的数量:", testUpdateBookCount)
}

// 根据book_id获取对应购物项
func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemsByBookIDAndCartID("4", "66668888")
	fmt.Println("图书id=4的购物项信息是:", cartItem)
}

// 根据cart_id获取购物车中所有的购物项
func testGetCartItemByCartID(t *testing.T) {
	cartItems, _ := GetCartItemsByCartID("66668888")
	for k, v := range cartItems {
		fmt.Printf("第%v个购物项是:%v\n", k+1, v)
	}
}

// 更新图书id和购物车id以及图书数量更新购物项中图书的数量
func testUpdateBookCount(t *testing.T) {
	UpdateBookCount(20, 3, "66668888")
}
