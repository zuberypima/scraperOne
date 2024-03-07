package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:imageurl`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)

	var items []item
	c.OnHTML("div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item{
			Name:     h.ChildText("h2.product-title"),
			Price:    h.ChildText("div.sale-price"),
			ImageUrl: h.ChildAttr("img", "src"),
		}

		fmt.Println(item)

		items = append(items, item)

	})

	c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Net Page")
		fmt.Println(r.URL.String())
	})

	c.Visit("https://j2store.net/demo/index.php/shop")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile("product.json", content, 0644)

}
