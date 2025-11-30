package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

type Config struct {
	DefaultCity string `json:"default_city"`
}

type WeatherResponse struct {
	Name    string        `json:"name"`
	Weather []WeatherInfo `json:"weather"`
	Main    MainData      `json:"main"`
}

type MainData struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"humidity"`
}

type WeatherInfo struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

var config Config

func main() {
	if len(os.Args) <= 1 {
		byte, err := os.ReadFile("config.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.Unmarshal(byte, &config)
		if err != nil {
			fmt.Println(err)
			return
		}

	} else {
		config = Config{DefaultCity: os.Args[1]}

	}

	APIkey := os.Getenv("API")
	err := godotenv.Load()
	if err != nil {
		return
	}

	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + config.DefaultCity + "&appid=" + APIkey + "&units=metric")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	weatherRes := WeatherResponse{}

	err = json.Unmarshal(data, &weatherRes)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(weatherRes.Name)
	fmt.Println(weatherRes.Weather[0].Description)
	fmt.Println(weatherRes.Main.Temp, "°C")
}
