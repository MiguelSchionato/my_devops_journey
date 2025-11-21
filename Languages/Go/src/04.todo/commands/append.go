package commands

import (
	"todo/logic"
)

func AppendList(listName, File string) error {
	allLists, err := logic.Unmarshing(File)
	if err != nil {
		return err
	}

	data, err := logic.Unmarshing(listName)
	if err != nil {
		return err
	}
	allLists = append(data, allLists...)

	err = logic.Marshing(allLists, File)
	if err != nil {
		return err
	}

	return nil
}
