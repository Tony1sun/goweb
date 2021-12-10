package model

// 订单
type Order struct {
	OrderID     string  // 订单号
	CreateTime  string  // 生成订单的时间
	TotalCount  int64   // 订单总数量
	TotalAmount float64 // 订单中图书总金额
	State       int64   // 订单的状态 0 未发货 1 已发货 2 交易完成
	UserID      int64   // 订单所属用户
}

// 未发货
func (order *Order) NoSend() bool {
	return order.State == 0
}

// 已发货
func (order *Order) SendComplate() bool {
	return order.State == 1
}

// 已收货
func (order *Order) Complete() bool {
	return order.State == 2
}
