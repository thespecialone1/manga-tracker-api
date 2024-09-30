package api

import (
	"encoding/json"
	"fmt"
	"io"
	"manga-tracker-api/internal/manga"
	"net/http"
)
func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func GetManga(w http.ResponseWriter, r *http.Request) {
				//call ScrapingManga and get the scraped data
	mangaList, err:= manga.ScrapingManga()
	if err!= nil {
		http.Error(w, "Failed to fetch manga data", http.StatusInternalServerError)
		return
	}
				//Json :- set reponse type as json
	w.Header().Set("Content-Type", "application/json")
				//send the data as json
	json.NewEncoder(w).Encode(mangaList)
}