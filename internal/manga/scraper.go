package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Set user-agent to mimic a browser request
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	// Visit the URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Extract data from the table
	c.OnHTML("#content-zone > div > table > tbody > tr", func(e *colly.HTMLElement) {
		// Each <tr> represents a row in the table
		title := e.ChildText("td:nth-child(1)")  // First column for release date
		timePages := e.ChildText("td:nth-child(2)")        // Second column for title
		publisher := e.ChildText("td:nth-child(3)")       // Third column for format
		releaseDate := e.ChildText("td:nth-child(4)")    // Fourth column for publisher

		// Print the extracted data
		fmt.Printf("%s, %s, %s, %s\n", title, timePages, publisher, releaseDate)
	})

	// Handle any errors that occur during the scraping
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Visit the page
	err := c.Visit("https://www.animenewsnetwork.com/encyclopedia/releases.php?format=manga")
	if err != nil {
		log.Fatal("Error visiting the page:", err)
	}
}
