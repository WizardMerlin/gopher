package main

import (
	"sync"
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

var wg sync.WaitGroup

func main() {
	http.HandleFunc("/news/", newsHandler)
	http.ListenAndServe(":8080", nil)
}

func newsHandler(w http.ResponseWriter, r *http.Request) {

	var s Sitemapindex


	//获取 xml 信息
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	if err != nil {
		log.Fatal(err)
	}

	bytes, err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	xml.Unmarshal(bytes, &s) //拿到具体的新闻地址 Sitemapindex.locations
	if len(s.Locations)==0 {
		log.Fatal("get the news xml locations error")
	}

	queue := make(chan News, 50) //缓冲队列有50，然后开启100个 goroutines
	log.Println(len(s.Locations)) //22
	for _, Location := range s.Locations {
		wg.Add(1)
		go worker(queue, Location) //往 News 队列里面写拿到的 xml 具体信息
	}

	wg.Wait()
	close(queue)


	// 从 chan 中拿到 News 写入 map
	newsMap := make(map[string]NewsMap) //存储单个每一条新闻
	for elem := range queue {
		for idx := range elem.Keywords {
			newsMap[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

	p := NewsPara{Title: "Wow", News: newsMap}
	t, _ := template.ParseFiles("complex.html")
	//t.Execute(w, p)
	fmt.Println(t.Execute(w, p)) //在控制台打印消息 (解析错误会有报错)
}

func worker(newsChan chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location) //1
	bytes, _ := ioutil.ReadAll(resp.Body) //2
	resp.Body.Close()
	xml.Unmarshal(bytes, &n) //拿到具体新闻 News

	newsChan <- n
}