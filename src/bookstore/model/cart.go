package model

// 购物车结构体
type Cart struct {
	CartID      string      // 购物车id
	CartItems   []*CartItem // 购物车中的购物项
	TotalCount  int64       // 购物车中图书总数量,通过计算得到
	TotalAmount float64     // 购物车中图书总金额，通过计算得到
	UserID      int64       // 当前购物车所属的用户
}

// 获取购物车图书总数量
func (c *Cart) GetTotalCount() int64 {
	// 遍历购物车中的购物项切片
	var totalCount int64
	for _, v := range c.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

// 获取购物车中图书总金额
func (c *Cart) GetTotalAmount() float64 {
	// 遍历购物车中的购物项切片
	var totalAmount float64
	for _, v := range c.CartItems {
		totalAmount = totalAmount + v.Amount
	}
	return totalAmount
}
