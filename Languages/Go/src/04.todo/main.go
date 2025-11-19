package main

import (
	"fmt"
	"os"
	"todo/commands"
	"todo/logic"
)

func main() {
	err := logic.CheckArgs(1)
	if err != nil {
		return
	}
	command := os.Args[1]
	var File = "test.json"

	switch command {
	case "add":
		err := logic.CheckArgs(2)
		if err != nil {
			return
		}
		commands.Add(os.Args[2], File)

	case "ls", "list":
		if len(os.Args) < 2 {
			commands.Ls(os.Args[2], File)
		} else {
			commands.Ls("default", File)
		}
	case "rm", "remove":
		if len(os.Args) >= 2 {
			commands.RemoveTask(os.Args[2], "default", File)
		} else {
			commands.RemoveTask(os.Args[2], os.Args[3], File)
		}
	case "done":
		if len(os.Args) >= 2 {
			commands.Done(os.Args[2], "default", File)
		} else {
			commands.Done(os.Args[2], os.Args[3], File)
		}

	default:
		fmt.Println("default option")
	}
}
