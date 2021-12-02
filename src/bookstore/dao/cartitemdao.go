package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 向购物项表中插入购物项
func AddCartItem(c *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values (?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, c.Count, c.GetAmount(), c.Book.ID, c.CartId)
	if err != nil {
		return err
	}
	return nil
}
