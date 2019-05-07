package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Note struct {
	Text string `json:"text"`
}

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

func constructNote(text string) Note {
	return Note{
		Text: text,
	}
}

func main() {
	token := os.Getenv("AYB_BOT_AUTHORIZATION_TOKEN")
	if token == "" {
		log.Fatal("env AYB_BOT_AUTHORIZATION_TOKEN not defined")
	}
	fmt.Println(token)
}
