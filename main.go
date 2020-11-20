package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/davidpoulos/shipwrecked/shipwreck"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// ...
	}

	es, _ := elasticsearch.NewClient(cfg)

	log.Println(es.Info())

	swdb := shipwreck.NewShipwreckDB(es)

	f, err := os.Open("./resources/data/shipwrecks.txt")

	if err != nil {
		panic(err)
	}

	seedShipwrecks(f, swdb)

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

// seedShipwrecks
func seedShipwrecks(r io.ReadCloser, swdb *shipwreck.ShipwreckDB) {
	defer r.Close()

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		temp := shipwreck.NewShipwreck()
		err := json.Unmarshal(scanner.Bytes(), temp)

		if err != nil {
			fmt.Println(err)
		}

		swdb.Insert(*temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
