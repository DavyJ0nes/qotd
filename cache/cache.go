package cache

import (
	"io/ioutil"
	"os"
)

// Write is a wrapper around WriteFile
// It Writes to cache file
func Write(file string, data []byte) error {
	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Read is a wrapper around ReadFile
// It reads from cache file
func Read(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}
	return data, nil

}

// Reset removes the cache file
func Reset(file string) error {
	if err := os.Remove(file); err != nil {
		return err
	}
	return nil
}
