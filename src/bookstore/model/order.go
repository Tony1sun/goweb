package model

import "time"

type Order struct {
	OrderID     string    // 订单号
	CreateTime  time.Time // 生成订单的时间
	TotalCount  int64     // 订单总数量
	TotalAmount float64   // 订单中图书总金额
	State       int64     // 订单的状态 0 未发货 1 已发货 2 交易完成
	UserID      int64     // 订单所属用户
}