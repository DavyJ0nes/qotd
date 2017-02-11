package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/davyj0nes/quoteOfTheDay/content"
)

var (
	tmpl      *template.Template
	e         = content.APIError{}
	cacheFile = "cache.txt"
)

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		if err := content.QuoteContent(w, req, "http://quotes.rest/qod.json?category=management"); err != nil {
			log.Fatal(err)
		}
	}
	d := content.CacheCheck(w, req)

	data := struct {
		Quote  string
		Author string
		BG     string
		Date   string
	}{
		Quote:  d.Contents.Quotes[0].Quote,
		Author: d.Contents.Quotes[0].Author,
		BG:     d.Contents.Quotes[0].Background,
		Date:   d.Contents.Quotes[0].Date,
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "index.html", data)

}
