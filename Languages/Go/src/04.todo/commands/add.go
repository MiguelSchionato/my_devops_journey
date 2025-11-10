package commands

import (
	"fmt"
	"os"
)

func Add() {
	fmt.Println("simple file was made")
	add, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Couldn't create file:", err)
		return
	}

	add.Write([]byte(os.Args[2]))

}
