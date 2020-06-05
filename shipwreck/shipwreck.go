package shipwreck

import "github.com/elastic/go-elasticsearch/v7"

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

// ShipwreckES ...
type ShipwreckDB struct {
	db *elasticsearch.Client
}

func NewShipwreckES(e *elasticsearch.Client) *ShipwreckDB {
	return &ShipwreckDB{
		db: e,
	}
}

func (*ShipwreckDB) Insert(s Shipwreck) {
	// Insert Code Here

}

func (*ShipwreckDB) Delete() {

}

// ETC /ETC
