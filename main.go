package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	token := os.Getenv("AYB_BOT_AUTHORIZATION_TOKEN")
	if token == "" {
		log.Fatal("env AYB_BOT_AUTHORIZATION_TOKEN not defined")
	}
	fmt.Println(token)
}
