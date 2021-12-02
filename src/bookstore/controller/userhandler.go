package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"html/template"
	"net/http"
)

// 登录
func Login(w http.ResponseWriter, r *http.Request) {
	// 判断是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		// 已经登录
		// 去首页
		GetPageBooksByPrice(w, r)
	} else {
		// 获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		// 调用 userdao 中验证用户名和密码的方法
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			// 用户名和密码正确
			// 用户名不正确
			// 生成uuid作为Session的id
			uuid := utils.CreateUUID()
			// 创建session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.Username,
				UserID:    user.ID,
			}
			// 将Session保持到数据库
			dao.AddSession(sess)
			// 创建Cookie， 与Session关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			// 将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			// 用户名和密码不正确
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "登录失败，请检查输入的用户名和密码。")
		}
	}
}

// 注销
func Logout(w http.ResponseWriter, r *http.Request) {
	// 获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		// 获取cookie的value值
		cookieValue := cookie.Value
		// 删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		// 设置cookie失效
		cookie.MaxAge = -1
		// 将修改之后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	// 去首页
	GetPageBooksByPrice(w, r)
}

// 注册
func Regist(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	// 调用 userdao 中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		// 用户名已经存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！请重新输入。")
	} else {
		// 把用户名保存到数据库
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	// 获取用户名
	username := r.PostFormValue("username")
	fmt.Println("传入的用户名是：", username)
	// 验证用户名
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		// 用户名已经存在
		w.Write([]byte("用户名已存在！"))
	} else {
		// 用户名可用
		w.Write([]byte("用户名可用！"))
	}
}
