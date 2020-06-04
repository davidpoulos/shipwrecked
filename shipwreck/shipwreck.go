package shipwreck

import "github.com/elastic/go-elasticsearch/v7"

// Shipwreck ...
type Shipwreck struct {
	Name            string
	Coordinates     string
	YearBuilt       int
	YearSank        int
	Depth           string
	DifficultyLevel string
	Latitude        string
	Longitude       string
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
