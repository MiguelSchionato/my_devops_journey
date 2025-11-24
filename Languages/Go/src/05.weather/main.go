package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	defaultCity string `json:"default_city"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error, couldn´t load the environment file:", err)
	}
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Failed to onpen cpnfig file")
	}
	res, err := http.Get("http://api.openweathermap.org")
	if err != nil {
		panic(err)
	}
	fmt.Print(res)
}
