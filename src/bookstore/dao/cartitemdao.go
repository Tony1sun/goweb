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

// 根据book_id获取对应的购物项
func GetCartItemsByBookID(bookID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	// 创建cartItem
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

// 根据cart_id获取购物车中所有的购物项
func GetCartItemsByCartID(CartID string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where cart_id = ?"
	// 执行
	rows, err := utils.Db.Query(sqlStr, CartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		// 创建cartItem
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}
