package shipwreck

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

// Shipwreck ...
type Shipwreck struct {
	Name            string `json:"name"`
	Coordinates     string `json:"coordinates"`
	YearBuilt       int    `json:"yearBuilt"`
	YearSank        int    `json:"yearSank"`
	Depth           string `json:"depth"`
	DifficultyLevel string `json:"difficultyLevel"`
	Latitude        string `json:"latitude"`
	Longitude       string `json:"longitude"`
}

// New ...
func NewShipwreck() *Shipwreck {
	return &Shipwreck{}
}

// ShipwreckDB ...
type ShipwreckDB struct {
	db *elasticsearch.Client
}

// NewShipwreckDB ...
func NewShipwreckDB(e *elasticsearch.Client) *ShipwreckDB {
	return &ShipwreckDB{
		db: e,
	}
}

// Insert ...
func (sw *ShipwreckDB) Insert(s Shipwreck) {

	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	res, err := sw.db.Index(
		"shipwreck",
		strings.NewReader(string(b)),
	)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res)

}

// Delete ...
func (sw *ShipwreckDB) Delete() {

}
