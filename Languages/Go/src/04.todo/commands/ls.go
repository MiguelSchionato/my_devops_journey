package commands

import (
	"errors"
	"fmt"
	"todo/logic"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Bold   = "\033[1m"
	Strike = "\033[9m"
)

func Ls(listName, File string) error {
	allLists, err := logic.Unmarshing(File)
	if err != nil {
		return err
	}

	listIndex := logic.FindList(listName, allLists)
	if listIndex == -1 {
		fmt.Println("List not found")
		return errors.New("List not found")
	}

	for _, task := range allLists[listIndex].Tasks {
		if task.State == 1 {
			fmt.Printf("- %s%s%s%s\n", Green, Strike, task.Name, Reset)

		} else {
			fmt.Printf("- %s\n", task.Name)

		}
	}
	return nil
}
