package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "处理请求！")

}

func main() {
	myHandler := MyHandler{}
	//http.Handle("/myHandler", &myHandler)
	// 创建Server结构
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}

	//http.ListenAndServe(":8080", nil)
	server.ListenAndServe()
}
