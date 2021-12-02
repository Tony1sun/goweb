package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 把购物项添加到购物车
func AddCart(c *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values (?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, c.CartID, c.TotalCount, c.TotalAmount, c.UserID)
	if err != nil {
		return err
	}
	// 获取购物车中的所有购物项
	cartItems := c.CartItems
	// 遍历得到每一个购物项
	for _, cartItem := range cartItems {
		// 将购物项插入到数据库中
		AddCartItem(cartItem)
	}
	return nil
}
