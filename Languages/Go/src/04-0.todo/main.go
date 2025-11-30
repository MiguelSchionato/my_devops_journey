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

	config, err := configs.CurrentConfig()
	if err != nil {
		fmt.Printf("Unable to read list file: %v", err)
		return
	}
	File := config.CurrentListPath
	List := config.CurrentList

	switch command {
	case "add":
		err := logic.CheckArgs(2)
		if err != nil {
			return
		}
		commands.Add(os.Args[2], List, File)

	case "ls", "list":
		if len(os.Args) > 2 {
			err := commands.ListTask(os.Args[2], List, File)
			if err != nil {
				return
			}
		} else {
			err := commands.ListTaskName(List, File)
			if err != nil {
				return
			}
		}
	case "rm", "remove":
		if len(os.Args) == 3 {
			err := commands.RemoveTask(os.Args[2], List, File)
			if err != nil {
				return
			}
		} else if len(os.Args) == 4 {
			err := commands.RemoveTask(os.Args[2], os.Args[3], File)
			if err != nil {
				return
			}
		} else {
			fmt.Println("Usage: todo rm <task> [list]")
		}
	case "done":
		if len(os.Args) == 3 {
			err := commands.Done(os.Args[2], List, File)
			if err != nil {
				return
			}
		} else if len(os.Args) == 4 {
			err := commands.Done(os.Args[2], os.Args[3], File)
			if err != nil {
				return
			}
		} else {
			fmt.Println("Usage: todo done <task> [list]")
		}
	case "append":
		if len(os.Args) == 4 {
			err := commands.AppendList(os.Args[2], os.Args[3])
			if err != nil {
				return
			}
		} else {
			fmt.Println("Usage: todo append <listToAppend> <DataBase>")
		}

	case "change":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo change <file/list> <newFile/list>")
			return
		}

		switch os.Args[2] {
		case "file":
			err := commands.ChangeListFile(os.Args[3])
			if err != nil {
				return
			}

		case "list":
			err := commands.ChangeList(os.Args[3])
			if err != nil {
				return
			}
			return
		}
	}
}
