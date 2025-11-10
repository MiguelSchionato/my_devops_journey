package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://api.openweathermap.org")
	if err != nil {
		panic(err)
	}
	fmt.Print(res) 

}
