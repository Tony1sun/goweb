package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 插入订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	sql := "insert into order_items(count, amount, title, author, price, img_path, order_id) values(?,?,?,?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}
