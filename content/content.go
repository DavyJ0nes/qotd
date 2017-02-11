package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davyj0nes/qotd/cache"
)

var cacheFile = os.Getenv("QOTD_CACHE_FILE")

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

type APIError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

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

	if res.StatusCode != 200 {
		log.Println("Error: ", res.Status)
		http.Redirect(w, req, "/fourohfour", http.StatusSeeOther)
		return errors.New(fmt.Sprintf("API returned: %s", res.Status))
	}

	if err := cache.Write(cacheFile, body); err != nil {
		return err
	}
	return nil
}

func CacheCheck(w http.ResponseWriter, req *http.Request) QuoteAPI {
	cachedData, err := cache.Read(cacheFile)
	if err != nil {
		log.Fatal(err)
	}

	d := QuoteAPI{}
	if err := json.Unmarshal(cachedData, &d); err != nil {
		log.Fatal(err)
	}

	// if cache is invalid then reset it.
	today := time.Now().Format("2006-01-02")
	cacheDate := d.Contents.Quotes[0].Date
	log.Println("cacheDate", cacheDate)
	if today != cacheDate {
		if err := QuoteContent(w, req, "http://quotes.rest/qod.json?category=management"); err != nil {
			log.Fatal(err)
		}
	}
	return d
}
