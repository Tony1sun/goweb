package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 插入订单
func AddOrder(order *model.Order) error {
	sql := "insert into orders(id, create_time, total_count, total_amount, state, user_id) values (?, ?, ?, ?, ?, ?)"
	// 执行
	_, err := utils.Db.Exec(sql, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}
