package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("./hello.tmpl");
	if err != nil{
		fmt.Println("Parse Template fail, err:%v", err);
		return
	}
	name := "HeZho"
	err = t.Execute(w, name);
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
