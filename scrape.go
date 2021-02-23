package main

import (
	"fmt"

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
	c.OnHTML(".tag-dimitri", func(e *colly.HTMLElement) {
		post := post{e.ChildText(".post-title a"), e.ChildText(".entry p"), e.ChildText(".postinfo a")}

		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, post)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://serenesforest.net/")
}
