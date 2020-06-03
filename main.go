package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/davidpoulos/shipwrecked/shipwreck"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

// Website
const (
	ShipWreckCrawlSite = "https://www.shipwreckworld.com/maps/a-e-vickery-schooner-st-lawrence-river-shipwreck"
)

func main() {
	ScrapeWebsite()
}

func launchServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	i := shipwreck.NewShipwreck()

	fmt.Println(i)

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// ...
	}

	es, _ := elasticsearch.NewClient(cfg)

	log.Println(es.Info())
}

func ScrapeWebsite() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// Find and visit all links
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		wreckName := e.ChildText("#ContentPlaceHolder1_TitleSecondLabel")
		latitude := e.ChildText("#map-description > div > div:nth-child(3) > strong:nth-child(2)")
		longitude := e.ChildText("#map-description > div > div:nth-child(3) > strong:nth-child(3)")
		yearBuilt := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(1) > h4 > span")
		yearSank := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(2) > h4 > span")
		difficultyLevel := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(4) > h4 > span")
		depth := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(3) > h4")

		fmt.Printf("Wreck Name: %s\n", wreckName)
		fmt.Printf("Latitude:   %s\n", latitude)
		fmt.Printf("Longitude:  %s\n", longitude)
		fmt.Printf("Year Built: %s\n", yearBuilt)
		fmt.Printf("Year Sank: %s\n", yearSank)
		fmt.Printf("Difficulty Level: %s\n", difficultyLevel)

		depthContent := []rune(depth)
		depth = strings.TrimSpace(string(depthContent[5:]))
		fmt.Printf("Depth: %s\n", depth)

		//	fmt.Printf("Description:%s\n", description)

		// type Shipwreck struct {
		// 	Name            string
		// 	Coordinates     string
		// 	YearSank        int
		// 	Depth           string
		// 	DifficultyLevel string
		// 	Latitude        string
		// 	Longitude       string
		//}

		// JSON marshall -> Write to file

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(ShipWreckCrawlSite)
}
