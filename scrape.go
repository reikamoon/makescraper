package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type post struct {
	Title string
	Entry string
	Info  string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML(".tag-Dimitri", func(e *colly.HTMLElement) {
		post := post{e.ChildText(".post-title a"), e.ChildText(".entry p"), e.ChildText(".postinfo a")}

		fmt.Println("Title: ", post.Title)
		fmt.Println("Entry: ", post.Entry)
		fmt.Println("Info: ", post.Info)
		fmt.Println()
	})

	c.OnHTML(".next_post_links", func(e *colly.HTMLElement) {
		// Navigate to page
		e.Request.Visit(e.Attr("href"))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://serenesforest.net/tag/three-houses/")
}

func writeJson(data []byte, flag int) {
	f, err := os.OpenFile("output.json", flag, 0644)
	checkErr(err)
	defer f.Close()

	n, err := f.Write(data)
	checkErr(err)
	fmt.Printf("Wrote %d bytes to %s\n", n, f.Name())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
