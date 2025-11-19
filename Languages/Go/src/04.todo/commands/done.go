package commands

import (
	"errors"
	"fmt"
	"todo/logic"
)

const (
	StatusPending = 0
	StatusDone    = 1
)

func Done(taskName, listName, File string) error {
	allLists, err := logic.Unmarshing(File)
	if err != nil {
		fmt.Printf("erro %v", err)
		return err
	}

	listIndex := logic.FindList(listName, allLists)
	if listIndex == -1 {
		return errors.New("Todo List not found")
	}

	targetList := &allLists[listIndex]

	taskIndex := logic.FindTask(listIndex, taskName, allLists)
	if taskIndex == -1 {
		return errors.New("Task not found")
	}
	targetTask := &targetList.Tasks[taskIndex]

	targetTask.State = StatusDone
	err = logic.Marshing(allLists, File)
	if err != nil {
		return err
	}
	return nil
}
