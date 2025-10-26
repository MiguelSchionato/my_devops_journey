package main

import (
	"os"
	"log"
	"fmt"
)

func main(){
	var directory string

	if len(os.Args) < 2 {
		directory = "."
	} else {
		directory = os.Args[1]
	}

	filesList, errors := os.ReadDir(directory)
	if errors != nil {
		log.Fatal(errors)
	}

	for _, files := range filesList{
	fmt.Println(files)
	}
}
