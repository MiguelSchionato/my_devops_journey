package commands

import (
	"errors"
	"fmt"
	"todo/logic"
	"todo/templates"
)

func RemoveTask(taskName, listName, File string) error {
	allLists, err := logic.Unmarshing(File)
	if err != nil {
		return err
	}

	listIndex := logic.FindList(listName, allLists)
	if listIndex == -1 {
		fmt.Print("Task not found")
		return errors.New("Task not found")
	}
	taskIndex := logic.FindTask(listIndex, taskName, allLists)

	err = deletingTasks(taskIndex, listIndex, allLists, File)
	if err != nil {
		return err
	}

	return nil
}

func deletingTasks(taskIndex, listIndex int, allLists []templates.Lists, File string) error {

	if taskIndex != -1 { // Removing logic
		previousTasks := allLists[listIndex].Tasks[:taskIndex]
		nextTasks := allLists[listIndex].Tasks[taskIndex+1:]
		allLists[listIndex].Tasks = append(previousTasks, nextTasks...)

		err := logic.Marshing(allLists, File)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Task not found")
	}
	return nil
}
