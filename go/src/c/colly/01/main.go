package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var visited = map[string]bool{}

func main() {
	// Instantiate default collector.
	c := colly.NewCollector(
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("link found: %q -> %s\n", e.Text, link)

		// Only those links are visited which are in AllowedDomains.
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ...".
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting:", r.URL.String())
	})

	if err := c.Visit("https://hackerspaces.org/"); err != nil {
		fmt.Println(err)
	}
}
