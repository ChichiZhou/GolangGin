package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	// 解析模板
	t,_ := template.ParseFiles("./hello.tmpl");

	name := "HeZho"
	// 渲染模板
	err := t.Execute(w, name);
	if err != nil {
		fmt.Println("Render template failed, err:%v", err);
		return
	}

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil{
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}

