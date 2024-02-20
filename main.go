package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//retrieve url link
		link := e.Attr("href")
		fmt.Println(link)
	})

	c.Visit("https://google.com")
}
