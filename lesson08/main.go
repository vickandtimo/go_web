package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.New("index.tmpl").Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	//渲染模板
	msg := "cfmoto"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("execute template failed,err:%v\n", err)
		return
	}
}
func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v", err)
		return
	}
}
