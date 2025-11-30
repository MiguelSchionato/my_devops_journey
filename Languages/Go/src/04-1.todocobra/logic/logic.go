package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"todoCobra/configs"
	"todoCobra/templates"
)

var allLists []templates.Lists

func UnmarshingCurrentFile() ([]templates.Lists, error) {
	File, err := configs.CurrentConfig()
	if err != nil {
		return allLists, err
	}
	data, err := os.ReadFile(File.CurrentListPath) //reading logic
	if err != nil && !os.IsNotExist(err) {         // worst err case
		fmt.Printf("Couldn't read file: %v\n", err)
		return nil, err
	}
	if os.IsNotExist(err) || len(data) == 0 { // best err case, os.IsNotExist.
		allLists := []templates.Lists{} // using a default empty list
		return allLists, nil
	}
	err = json.Unmarshal(data, &allLists) //Unmarshing json
	if err != nil {
		fmt.Printf("Couldn't read file: %v\n", err)
		return nil, err
	}

	// fmt.Println("Listas carregadas (ou inicializadas) com sucesso.") // debug
	return allLists, err
}

func MarshingToCurrentFile(Lists []templates.Lists) error {
	File, err := configs.CurrentConfig()
	if err != nil {
		return err
	}

	data, err := json.Marshal(Lists)
	if err != nil {
		fmt.Printf("Fail writting file: %v\n", err)
		return err
	}

	err = os.WriteFile(File.CurrentListPath, data, 0640)
	if err != nil {
		fmt.Printf("Fail writting file: %v\n", err)
		return err
	}
	return nil
}

func FindList(listName string, allLists []templates.Lists) int {
	// Linear search: O(n).
	// O(1) on the best case (first item)

	listIndex := -1

	for i, list := range allLists { // searching for the right list
		if list.Name == listName {
			listIndex = i
			break
		}
	}
	return listIndex
}

func FindTask(listIndex int, taskID string, allLists []templates.Lists) int {
	taskIndex := -1
	taskIdInt, err := strconv.Atoi(taskID)
	if err != nil {
		return -1
	}

	for i, task := range allLists[listIndex].Tasks {
		if task.ID == taskIdInt {
			taskIndex = i
			break
		}
	}
	return taskIndex
}

func CheckArgs(id int) error {
	if len(os.Args) <= id {
		fmt.Println("Use: <Command> <Task>")
		return errors.New("not enough arguments")
	}
	return nil
}
