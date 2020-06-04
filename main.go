package main

import (
	"fmt"
	"log"

	"github.com/davidpoulos/shipwrecked/scraper"
	"github.com/davidpoulos/shipwrecked/shipwreck"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

func main() {
	scraper.ScrapeWebsite()
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
