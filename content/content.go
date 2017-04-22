package content

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/davyj0nes/qotd/cache"
)

var cacheFile = os.Getenv("QOTD_CACHE_FILE")

// QuoteAPI is used to parse the QUOTE of the DAY API data
type QuoteAPI struct {
	Contents struct {
		Quotes []struct {
			Quote      string `json:"quote"`
			Author     string `json:"author"`
			Date       string `json:"date"`
			Title      string `json:"title"`
			Background string `json:"background"`
		} `json:"quotes"`
	} `json:"contents"`
}

// APIError is used to parse error response
// not used at the moment
type APIError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// QuoteContent calls the QOTD API and writes the response to cache file
func QuoteContent(w http.ResponseWriter, req *http.Request, apiUrl string) error {
	res, err := http.Get(apiUrl)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := cache.Write(cacheFile, body); err != nil {
		return err
	}
	return nil
}

// CacheCheck checks if cache file exists as well as the staleness of the cache data
func CacheCheck(w http.ResponseWriter, req *http.Request) (QuoteAPI, error) {
	d := new(QuoteAPI)
	cachedData, err := cache.Read(cacheFile)
	if err != nil {
		return d, err
	}

	if err := json.Unmarshal(cachedData, &d); err != nil {
		return d, err
	}

	if len(d.Contents.Quotes) < 1 {
		if err := QuoteContent(w, req, "http://quotes.rest/qod.json?category=management"); err != nil {
			return d, err
		}
		return d, err
	}

	// if cache is invalid then reset it.
	today := time.Now().Format("2006-01-02")
	cacheDate := d.Contents.Quotes[0].Date
	err = cacheDateCheck(w, req, today, cacheDate)
	if err != nil {
		return d, err
	}
	return d, nil
}

// cacheDateCheck abstracts the cache date checking
func cacheDateCheck(w http.ResponseWriter, req *http.Request, today, cacheDate string) error {
	if today != cacheDate {
		if err := QuoteContent(w, req, "http://quotes.rest/qod.json?category=management"); err != nil {
			return err
		}
	}
	return nil
}
