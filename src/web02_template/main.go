package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")

	t.Execute(w, "hello Template")
}


func main() {
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}
