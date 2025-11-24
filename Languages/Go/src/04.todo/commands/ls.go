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

func ListTaskName(listName, File string) error {
	allLists, err := logic.Unmarshing(File)
	if err != nil {
		return err
	}

	listIndex := logic.FindList(listName, allLists)
	if listIndex == -1 {
		fmt.Println("List not found")
		return errors.New("list not found")
	}

	for _, task := range allLists[listIndex].Tasks {
		if task.State == 1 {
			fmt.Printf("• %s%s%s%s\n", Green, Strike, task.Name, Reset)

		} else {
			fmt.Printf("•%d %s\n", task.ID, task.Name)

		}
	}
	return nil
}

func ListTask(listID, listName, File string) error {
	allLists, err := logic.Unmarshing(File)
	if err != nil {
		return err
	}

	listIndex := logic.FindList(listName, allLists)
	if listIndex == -1 {
		fmt.Println("task not found")
		return errors.New("task not found")
	}

	taskIndex := logic.FindTask(listIndex, listID, allLists)
	list := allLists[listIndex].Tasks

	fmt.Printf("• %s\n%s\n%s\n", list[taskIndex].Name, list[taskIndex].Descr, list[taskIndex].Due)

	return nil
}
