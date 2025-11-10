package main

import (
	"todo/commands"
	"errors"
	"fmt"
	"os"
)

func main() {
	err := checkArgs()
	if err != nil {
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		commands.Add()
	default:
		fmt.Println("default option")
	}
}

func checkArgs() error {
	if len(os.Args) <= 2 {
		fmt.Println("Use: <Command> <Task>")
		return errors.New("Not enough arguments")
	}
	return nil
}