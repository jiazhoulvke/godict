package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("请输入要查询的单词")
		return
	}
	url := "http://dict.youdao.com/search?q=" + strings.Join(flag.Args(), "%20")
	dom, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	wordbook := dom.Find(".wordbook-js .baav .pronounce")
	if wordbook.Length() > 0 {
		pronounces := ""
		wordbook.Each(func(_ int, node *goquery.Selection) {
			pronounce := ""
			node.Contents().Each(func(_ int, snode *goquery.Selection) {
				t := strings.TrimSpace(snode.Text())
				if t != "" {
					pronounce += t + " "
				}
			})
			pronounces += pronounce + " "
		})
		fmt.Println(pronounces)
	}
	dom.Find("#phrsListTab .trans-container ul li").Each(func(_ int, node *goquery.Selection) {
		fmt.Println(strings.TrimSpace(node.Text()))
	})
}
