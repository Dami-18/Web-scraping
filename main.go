package main

import(
	"fmt"
	"strings"
	"log"
	"github.com/gocolly/colly"
)

func main(){
	c := colly.NewCollector()
	url := "https://www.studypool.com/studyGuides"
	c.OnHTML(".study-guide-card", func(h *colly.HTMLElement){
		title := h.ChildText("h3.book-title")
		author := h.ChildText("p.book-author")

		title = strings.TrimSpace(title)
		author = strings.TrimSpace(author)

		fmt.Printf("%v : %v\n",title,author)
	})
	err := c.Visit(url)
    if err != nil {
    log.Fatal(err)
    }
}