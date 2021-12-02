package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestCart(t *testing.T) {
	fmt.Println("测试购物车相关函数")
	t.Run("测试添加购物车:", testAddCart)
}

func testAddCart(t *testing.T) {
	// 设置要买的第一本书
	book := &model.Book{
		ID:    3,
		Price: 30,
	}
	// 设置要买的第二本书
	book2 := &model.Book{
		ID:    4,
		Price: 27,
	}
	var cartItems []*model.CartItem
	// 创建两个购物项
	cartitem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartId: "66668888",
	}
	cartItems = append(cartItems, cartitem)
	cartitem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartId: "66668888",
	}
	cartItems = append(cartItems, cartitem2)
	// 创建购物车
	cart := &model.Cart{
		CartID:    "66668888",
		CartItems: cartItems,
		UserID:    1,
	}
	AddCart(cart)
}
