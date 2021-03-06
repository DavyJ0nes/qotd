package cache

import (
	"os"
	"testing"
)

// TestCache checks that data can be written and read from cache file
func TestCache(t *testing.T) {
	data := []byte("This is some data")
	// checking write to cache
	if err := Write("cache-test.txt", data); err != nil {
		t.Errorf("Write() returned error: %s", err)
	}

	fromCache, err := Read("cache-test.txt")
	if err != nil {
		t.Errorf("Read() returned error: %s", err)
	}

	if string(fromCache) != string(data) {
		t.Errorf("Expected: %s | Got: %s", data, fromCache)
	}
}

// TestReset checks that removing the cache file works as expected
// Also has benefit of cleaning up after testing
func TestReset(t *testing.T) {
	if err := Reset("cache-test.txt"); err != nil {
		t.Errorf("Reset() returned error: %s", err)
	}

	_, err := Read("cache-test.txt")
	if os.IsExist(err) {
		t.Errorf("Reset() did not delete cache file")
	}
}
