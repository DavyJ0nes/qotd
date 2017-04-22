package content

import (
	"net/http"
	"testing"
	"time"
)

func TestQuoteContent(t *testing.T) {
	var w http.ResponseWriter
	var r *http.Request
	testApi := "http://quotes.rest/qod.json?category=management"

	err := QuoteContent(w, r, testApi)
	if err != nil {
		t.Errorf("Error in QuoteContent(): %v", err)
	}
}

func TestCacheCheck(t *testing.T) {
	var w http.ResponseWriter
	var r *http.Request
	data, err := CacheCheck(w, r)
	if err != nil {
		t.Errorf("Error in CacheCheck(): %v", err)
	}

	today := time.Now().Format("2006-01-02")
	cacheDate := data.Contents.Quotes[0].Date
	if today != cacheDate {
		t.Errorf("Expected: %v | Got: %v", today, cacheDate)
	}
}

func TestCacheDateCheck(t *testing.T) {
	var w http.ResponseWriter
	var r *http.Request
	data, err := CacheCheck(w, r)
	if err != nil {
		t.Errorf("Error in CacheCheck(): %v", err)
	}
	testDate := "2006-01-02"
	cacheDate := data.Contents.Quotes[0].Date
	if err := cacheDateCheck(w, r, testDate, cacheDate); err != nil {

		t.Errorf("cacheDateCheck() Error: %v", err)
	}
}
