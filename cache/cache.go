package cache

import (
	"io/ioutil"
	"os"
)

func Write(file string, data []byte) error {
	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte(""), err
	}
	return data, nil

}

func Reset(file string) error {
	if err := os.Remove(file); err != nil {
		return err
	}
	return nil
}
