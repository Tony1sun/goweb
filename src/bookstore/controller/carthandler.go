package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

// 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	// 判断是否登陆
	flag, session := dao.IsLogin(r)
	if flag {
		// 获取要添加的图书id
		bookID := r.FormValue("bookId")
		// 根据图书id获取图书信息
		book, _ := dao.GetBookById(bookID)

		// 获取用户id
		userID := session.UserID
		// 判断当前用户是否有购物车
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			// 当前用户已经有购物车,需要判断购物车中是否有当前这本书
			cartItem, _ := dao.GetCartItemsByBookIDAndCartID(bookID, cart.CartID)
			if cartItem != nil {
				// 购物车中的购物项已经有该图书，只需要将该图书所对应的购物项中的数量加1即可
				// 1.获取购物车切片中的所有购物项
				cts := cart.CartItems
				// 2.遍历得到的每一个购物项
				for _, v := range cts {
					// 3.找到当前的购物项
					if v.Book.ID == cartItem.Book.ID {
						// 将购物项中的图书数量加1
						v.Count = v.Count + 1
						// 更新数据库中该购物项的图书数量
						dao.UpdateBookCount(v)
					}
				}
			} else {
				// 购物车中的购物项还没有该图书，此时需要创建一个购物项并添加到数据库中
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				// 将购物项添加到cart切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				// 将新创建的购物项添加到数据库中
				dao.AddCartItem(cartItem)
			}
			// 不管之前购物车中是否有当前图书对应的购物项，都需要更新购物车中的图书总数量和总金额
			dao.UpdateCart(cart)
		} else {
			// 当前用户还没有购物车，创建一个购物车并添加到数据库
			// 1.创建购物车
			// 生成uuid
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: int64(userID),
			}
			// 2.创建购物车中的购物项
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}
			// 将购物项添加到切片中
			cartItems = append(cartItems, cartItem)
			// 3.将切片设置到cart中
			cart.CartItems = cartItems
			// 将购物车cart保存到数据库
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车！"))
	} else {
		// 没有登录
		w.Write([]byte("请先登录！"))
	}
}

// 根据用户id获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	islogin, session := dao.IsLogin(r)
	if islogin {
		// 获取用户id
		userID := session.UserID
		// 根据用户id从数据库获取对应的购物车
		cart, _ := dao.GetCartByUserID(userID)
		// 把cart加到session的cart中
		session.Cart = cart
		// 解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		// 执行
		t.Execute(w, session)
	} else {
		// 该用户还没有购物车
		// 解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		// 执行
		t.Execute(w, session)
	}
}

// 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	// 获取要删除的购物车id
	cartID := r.FormValue("cartId")
	// 清空购物车
	dao.DeleteCartByCartID(cartID)
	// 再次查询购物车信息
	GetCartInfo(w, r)
}

// 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	// 获取要删除的购物项id
	cartItemID := r.FormValue("cartItemId")
	// 将购物项id转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	// 获取session
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	// 获取购物车中所有的购物项
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemId == iCartItemID {
			// 找到要删除的购物
			// 将当前购物项从切片中移除
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			// 将删除购物项之后的切片再次赋给购物车中的切片
			cart.CartItems = cartItems
			// 将当前购物项从数据库中移除
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	// 更新购物车中图书的总数量和金额
	dao.UpdateCart(cart)
	// 获取购物车信息
	GetCartInfo(w, r)
}

// 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	// 获取要删除的购物项id
	cartItemID := r.FormValue("cartItemId")
	// 将购物项id转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	// 获取用户输入的图书数量
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	// 获取session
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	// 获取购物车中所有的购物项
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.CartItemId == iCartItemID {
			// 找到要更新的购物项
			// 将当前购物项中的图书数量设置为用户输入的值
			v.Count = iBookCount
			// 更新数据库中该购物项的图书数量和金额小计
			dao.UpdateBookCount(v)
		}
	}
	// 更新购物车中图书的总数量和金额
	dao.UpdateCart(cart)
	// 获取购物车信息
	cart, _ = dao.GetCartByUserID(userID)
	// GetCartInfo(w, r)
	// 获取购物车中图书的总数量
	totalCount := cart.TotalCount
	// 获取购物车中图书的总金额
	totalAmount := cart.TotalAmount
	var amount float64
	// 获取购物车中更新的购物项的金额小计
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemID == v.CartItemId {
			// 找到要更新的购物项，获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	// 创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}

	// 将cart转换为json字符串
	json, _ := json.Marshal(data)
	w.Write(json)
}
