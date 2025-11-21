package main

import (
	"fmt"
	"os"
	"todo/commands"
	"todo/configs"
	"todo/logic"
)

func main() {
	err := logic.CheckArgs(1)
	if err != nil {
		return
	}
	command := os.Args[1]
	File, err := configs.CurrentList()
	if err != nil {
		fmt.Printf("Unable to read list file: %v", err)
		return
	}

	switch command {
	case "add":
		err := logic.CheckArgs(2)
		if err != nil {
			return
		}
		commands.Add(os.Args[2], File)

	case "ls", "list":
		if len(os.Args) > 2 {
			err := commands.Ls(os.Args[2], File)
			if err != nil {
				return
			}
		} else {
			err := commands.Ls("default", File)
			if err != nil {
				return
			}
		}
	case "rm", "remove":
		if len(os.Args) >= 2 {
			err := commands.RemoveTask(os.Args[2], "default", File)
			if err != nil {
				return
			}
		} else {
			err := commands.RemoveTask(os.Args[2], os.Args[3], File)
			if err != nil {
				return
			}
		}
	case "done":
		if len(os.Args) >= 2 {
			err := commands.Done(os.Args[2], "default", File)
			if err != nil {
				return
			}
		} else {
			err := commands.Done(os.Args[2], os.Args[3], File)
			if err != nil {
				return
			}
		}
	case "append":
		if len(os.Args) >= 2 {
			err := commands.AppendList(os.Args[2], os.Args[3])
			if err != nil {
				return
			}
		} else {
			fmt.Println("Usage: todo <listToAppend> <DataBase>")
		}

	case "change":
		if len(os.Args) > 2 {
			err := commands.ChangeList(os.Args[2])
			if err != nil {
				return
			}
		} else {
			fmt.Println("Usage: todo <NewDefaultList>")
			return
		}

	default:
		fmt.Println("default option")
	}
}
