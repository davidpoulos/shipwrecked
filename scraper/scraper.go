package scraper

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/davidpoulos/shipwrecked/shipwreck"
	"github.com/gocolly/colly/v2"
)

// Website
const (
	ShipWreckCrawlSite = "https://www.shipwreckworld.com/maps/a-e-vickery-schooner-st-lawrence-river-shipwreck"
)

func ScrapeWebsite() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	//	colly.Debugger(&debug.LogDebugger{}),
	)

	// Find and visit all links
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))

		r := ExtractShipWreckInfo(e)
		fmt.Printf("%v\n", r)

		b, _ := json.Marshal(r)
		fmt.Printf("%s", string(b))

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Limit(&colly.LimitRule{
		RandomDelay: 5 * time.Second,
	})

	c.Visit(ShipWreckCrawlSite)
}

func ExtractShipWreckInfo(e *colly.HTMLElement) *shipwreck.Shipwreck {
	// Scrape the fields from the HTML
	wreckName := e.ChildText("#ContentPlaceHolder1_TitleSecondLabel")
	latitude := e.ChildText("#map-description > div > div:nth-child(3) > strong:nth-child(2)")
	longitude := e.ChildText("#map-description > div > div:nth-child(3) > strong:nth-child(3)")
	yearBuilt := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(1) > h4 > span")
	yearSank := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(2) > h4 > span")
	difficultyLevel := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(4) > h4 > span")
	depth := e.ChildText("#ContentPlaceHolder1_articleWrapperDiv > div.row.map-times-wrapper > div:nth-child(3) > h4")

	return &shipwreck.Shipwreck{
		Name:            CleanName(wreckName),
		Latitude:        CleanLatitude(latitude),
		Longitude:       CleanLongitude(longitude),
		YearSank:        CleanYearSank(yearSank),
		YearBuilt:       CleanYearBuilt(yearBuilt),
		DifficultyLevel: CleanDifficultyLevel(difficultyLevel),
		Depth:           CleanDepth(depth),
	}
}

func CleanDepth(depth string) string {
	depthContent := []rune(depth)
	return strings.TrimSpace(string(depthContent[5:]))
}

func CleanName(name string) string {
	return name
}

func CleanYearBuilt(year string) int {
	y, _ := strconv.Atoi(year)
	return y
}

func CleanYearSank(year string) int {
	y, _ := strconv.Atoi(year)
	return y
}

func CleanLongitude(long string) string {
	longRunes := []rune(long)
	long = string(longRunes[10:])
	return long
}

func CleanLatitude(lat string) string {
	latRunes := []rune(lat)
	lat = string(latRunes[9:])
	return lat
}

func CleanDifficultyLevel(dLevel string) string {
	return dLevel
}
