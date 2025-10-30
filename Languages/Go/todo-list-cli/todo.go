package main

import (
	"os"
	"errors"
	"fmt"
	"todo/commands"
)

func main()  {
	err := checkArgs()
	if err != nil{
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
		return errors.New("Not enought arguments")
	}
	return nil
}


