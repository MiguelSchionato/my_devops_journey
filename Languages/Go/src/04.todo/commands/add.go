package commands

import (
	"fmt"
	"time"
	"todo/logic"
	"todo/templates"
)

func Add(taskName string, File string) {
	var allLists []templates.Lists
	newTask := templates.Task{
		Name: taskName,
		Date: time.Now(),
	}

	allLists, err := logic.Unmarshing(File)
	if err != nil {
		return
	}

	listIndex := logic.FindList("default", allLists)

	if listIndex != -1 {
		allLists[listIndex].Tasks = append(allLists[listIndex].Tasks, newTask)
	} else {
		newDefaultList := templates.Lists{
			Name:  "default",
			Descr: "Default list",
			Tasks: []templates.Task{newTask},
		}

		allLists = append(allLists, newDefaultList)
		fmt.Printf("List not found, appending to the default list") // debug
	}

	err = logic.Marshing(allLists, File)
	if err != nil {
		return
	}

}
