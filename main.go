package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func loadNotes(filename string) ([]string, error) {
	var (
		err   error
		bytes []byte
		notes []string
	)

	if bytes, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bytes, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func main() {
	fmt.Print("Hello, world!")
}
