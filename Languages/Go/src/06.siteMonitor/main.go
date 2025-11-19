package main

import (
	"fmt"
	"net/http"
)

func main() {
	site := "https://tech.miguelschionato.com"
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Site not found")
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site on")
		return
	} else {
		fmt.Println("Site off")
		return
	}

}
