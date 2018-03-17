package main

import (
	"log"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsPara struct {
	Title string
	News  map[string]NewsMap
}

func main() {
	http.HandleFunc("/news/", newsHandler)
	http.ListenAndServe(":8080", nil)
}

func newsHandler(w http.ResponseWriter, r *http.Request) {

	var s Sitemapindex
	var n News

	//获取 xml 信息
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	xml.Unmarshal(bytes, &s) //拿到具体的新闻地址 Sitemapindex.locations
	if len(s.Locations)==0 {
		log.Fatal("get the news xml locations error")
	}

	//调试用
/* 	for url := range s.Locations {
		log.Println(url)
		fmt.Println("s", url)
	} */

	//解析 xml
	newsMap := make(map[string]NewsMap) //存储单个每一条新闻
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n) //拿到具体新闻 News

		//这里应该检查一下 n
		if len(n.Keywords)==0 {
			log.Fatal("get the news xml content error")
		}

		//存储到 NewsMap中
		for idx := range n.Keywords {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}

		//调试用
/* 		for kk, vv := range newsMap {
			log.Printf("title: %s\n", kk)
			log.Printf("News: keywords {%s} ", vv.Keyword)
		} */

	}

	p := NewsPara{Title: "Wow", News: newsMap}
	t, _ := template.ParseFiles("complex.html")
	//t.Execute(w, p)
	fmt.Println(t.Execute(w, p)) //在控制台打印消息 (解析错误会有报错)
}
