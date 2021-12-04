package model

// 购物项
type CartItem struct {
	CartItemId int64   // 购物项id
	Book       *Book   // 购物项中的图书信息
	Count      int64   // 购物项中的图书数量
	Amount     float64 // 购物项中的图书金额小计，通过计算得到
	CartID     string  // 当前购物项属于哪一个购物车
}

// 获取购物项中图书的金额小计，数量*图书价格
func (c *CartItem) GetAmount() float64 {
	// 获取当前购物项中图书的价格
	price := c.Book.Price
	return float64(c.Count) * price
}
