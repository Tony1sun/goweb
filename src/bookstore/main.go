package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行
	t.Execute(w, "")
}


func main() {
	// 设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main", IndexHandler)

	http.ListenAndServe(":8080", nil)
}