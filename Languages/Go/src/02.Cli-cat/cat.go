package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please give a file path")
		return 
	}
	fileName := os.Args[1]

	fileText, err := os.ReadFile(fileName)

	if err != nil{
		log.Fatal(err)
	}

	os.Stdout.Write(fileText)
}
