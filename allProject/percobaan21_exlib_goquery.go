package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	ExampleCrawl()
}

func ExampleCrawl() {
	// Request the HTML page.
	res, err := http.Get("https://www.merdeka.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".mdk-left-l .mdk-link-item").Each(func(i int, s *goquery.Selection) {
		html_tag := s.Find("a").First()
		href, _ := html_tag.Attr("href")
		text := html_tag.Text()
		fmt.Println(i, text, " ==link==> ", href)
	})
}