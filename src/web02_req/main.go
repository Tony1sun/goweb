package main

import (
	"fmt"
	"net/http"
)

// 处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "发送的请求的请求地址是:", r.URL.Path)
	fmt.Fprintln(w, "发送的请求的请求地址后的查询字符串是:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有:", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息有:", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值有:", r.Header.Get("Accept-Encoding"))
	// 获取请求体内容的长度
	// len := r.ContentLength
	// // 创建byte切片
	// body := make([]byte, len)
	// // 将请求体中的内容读到body中
	// r.Body.Read(body)
	// // 在浏览器中显示请求体中的内容
	// fmt.Fprintln(w, "请求体中的内容有:", string(body))

	// 解析表单,在调用r.Form之前必须执行该操作
	r.ParseForm()
	// 获取请求参数
	fmt.Fprintln(w, "请求参数有:", r.Form)
	fmt.Fprintln(w, "POST请求的form表单中的请求参数有:", r.PostForm)
}

func main() {
	http.HandleFunc("/hello", handler)

	http.ListenAndServe(":8080", nil)
}
