package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/davyj0nes/qotd/content"
)

var (
	tmpl      *template.Template
	cacheFile = os.Getenv("QOTD_CACHE_FILE")
)

// init Parses all html templates for the app
func init() {
	templatePath := fmt.Sprintf("%s/*.html", os.Getenv("QOTD_TEMPLATE_PATH"))
	tmpl = template.Must(template.ParseGlob(templatePath))
}

// requestLogger logs request information in standard way
func requestLogger(req *http.Request) {
	log.Printf(">> %s | %s || %s => %s || %s", req.Method, req.URL.Path, req.RemoteAddr, req.Host, req.Header.Get("User-Agent"))
}

// Router is main mux wrangler. Keeps main() clean
func Router() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("templates/static"))))
	log.Println("Server Starting")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// index is the main route for the application
func index(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		if err := content.QuoteContent(w, req, os.Getenv("QOTD_URL")); err != nil {
			log.Fatal(err)
		}
	}
	d, err := content.CacheCheck(w, req)
	if err != nil {
		log.Fatal(err)
		return
	}

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
