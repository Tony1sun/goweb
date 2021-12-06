package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"net/http"
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
	_, session := dao.IsLogin(r)
	// 获取用户id
	userID := session.UserID
	// 根据用户id从数据库获取对应的购物车
	cart, _ := dao.GetCartByUserID(userID)
	// 设置用户名
	cart.UserName = session.UserName
	if cart != nil {
		// 解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		// 执行
		t.Execute(w, cart)
	}
}
