package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 把购物项添加到购物车
func AddCart(c *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values (?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, c.CartID, c.GetTotalCount(), c.GetTotalAmount(), c.UserID)
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

//  根据用户id从数据库中查询对应的购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	sqlStr := "select id, total_count, total_amount, user_id from carts where user_id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, userID)
	// 创建一个购物车
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	// 获取当前购物车中所有的购物项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	// 将所有的购物项设置到购物车中
	cart.CartItems = cartItems
	return cart, nil
}

// 更新购物车中的图书总数量和总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count = ? , total_amount = ? where id = ?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}
