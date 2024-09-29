package api

import (
	"fmt"
	"io"
	"net/http"
)
func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func GetManga(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<html>
		<head>
			<title>Manga Release Dates</title>
		</head>
		<body>
			<h1 style="font-size: 2em;">Here is the list of release dates of upcoming Manga</h1>
		</body>
	</html>`)
}