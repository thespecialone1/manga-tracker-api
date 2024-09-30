package manga

import (
	"fmt"
	"log"
	// "log"
	// "os"

	"github.com/gocolly/colly"
)


//struct
type Manga struct{
	Title		string  
	Pages		string
	Distributor	string
	ReleaseDate	string
	SRP			string	
}

func ScrapingManga()([]Manga, error) {
	var mangas []Manga
	c:= colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	url:= "https://www.animenewsnetwork.com/encyclopedia/releases.php?format=manga"

	c.OnRequest(func (r *colly.Request)  {
		fmt.Println("visiting...", r.URL)
	})

	c.OnHTML("#content-zone > div > table > tbody> tr", func (e *colly.HTMLElement)  {
		title:= e.ChildText("td:nth-child(1)")
		pages:= e.ChildText("td:nth-child(2)")
		distributor:= e.ChildText("td:nth-child(3)")
		releasedate:= e.ChildText("td:nth-child(4)")
		srp:= e.ChildText("td:nth-child(5)")

		manga:= Manga{
			Title: 	     title,
			Pages: 		 pages,
			Distributor: distributor,
			ReleaseDate: releasedate,
			SRP: 		 srp,	
		}

		mangas= append(mangas, manga)
	
	})

	c.OnError(func (_ *colly.Response, err error)  {
		log.Println("Error...", err)
	})

	err:= c.Visit(url)
	if err!=nil {
		log.Fatal("Error while visiting the url...", err)
	}
	return mangas, nil
}

	
