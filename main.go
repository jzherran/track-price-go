package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {

	// List of items with urls
	data := make(map[string]string)

	data["nevera:alkosto"] = "https://www.alkosto.com/nevera-samsung-394-litros-negro-rt38k5992bs"

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#product-price-88274 > span:nth-child(1) > span:nth-child(2)", func(e *colly.HTMLElement) {
		// Print link
		fmt.Printf("Price found: %q\n", e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for _, url := range data {
		fmt.Println("Iterate")
		c.Visit(url)
	}
}
