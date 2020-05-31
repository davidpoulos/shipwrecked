package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
) //import "github.com/gin-gonic/gin"

// Website
const (
	ShipWreckCrawlSite = "https://www.shipwreckworld.com/"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// ...
	}

	es, _ := elasticsearch.NewClient(cfg)
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
