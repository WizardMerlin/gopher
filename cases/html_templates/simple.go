package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new/", newsAggHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hi, there.</h1") //without template
}

type NewsPara struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsPara{Title: "Wow", News: "news"}
	t, _ := template.ParseFiles("simple.html") // {{ .Title}} , {{.News}}
	//t.Execute(w, p)
	fmt.Println(t.Execute(w, p)) //在控制台打印消息 (解析错误会有报错)
}
