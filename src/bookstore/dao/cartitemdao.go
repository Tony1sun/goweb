package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 向购物项表中插入购物项
func AddCartItem(c *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values (?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, c.Count, c.GetAmount(), c.Book.ID, c.CartID)
	if err != nil {
		return err
	}
	return nil
}

// 根据book_id和cart_id获取对应的购物项
func GetCartItemsByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	// 创建cartItem
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	// 根据bookID获取图书信息
	book, _ := GetBookById(bookID)
	// 将book设置到购物项中
	cartItem.Book = book
	return cartItem, nil
}

// 更新购物项中的信息更新图书数量和金额小计
func UpdateBookCount(cartItem *model.CartItem) error {
	sqlStr := "update cart_items set count = ?, amount = ? where book_id = ? and cart_id = ?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

// 根据cart_id获取购物车中所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	// 执行
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		// 设置一个变量接收bookId
		var bookID string
		// 创建cartItem
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		// 根据bookID获取图书信息
		book, _ := GetBookById(bookID)
		// 将book设置到购物项中
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

// 根据购物车id删除所有购物项
func DeleteCartItemsByCartID(cartID string) error {
	sql := "delete from cart_items where cart_id = ?"
	// 执行
	_, err := utils.Db.Exec(sql, cartID)
	if err != nil {
		return err
	}
	return nil
}
