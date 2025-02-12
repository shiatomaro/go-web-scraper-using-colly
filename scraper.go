package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("webscraper.io"),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Request URL: %v failed with response: %v\nError: %v\n", r.Request.URL, r, err)
	})

	var items []item

	c.OnHTML("div.product-wrapper", func(h *colly.HTMLElement) {
		item := item{
			Name:        h.ChildText("a.title"),
			Description: h.ChildText("p.description"),
			Price:       h.ChildText("h4.price"),
			ImgUrl:      h.ChildAttr("img", "src"),
		}
		items = append(items, item)
	})

	c.Visit("https://webscraper.io/test-sites/e-commerce/allinone")
	fmt.Println(items)

	content, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("sample-products.json", content, 0644)

}
