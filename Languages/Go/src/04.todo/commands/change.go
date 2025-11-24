package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"todo/configs"
)

func ChangeList(fileName string) error {
	config, err := configs.UnmarshingConfigs()
	if err != nil {
		return err
	}

	config.CurrentList = fileName
	fmt.Printf("List changed successfully: %v", config.CurrentList)

	err = configs.MarshingConfigs(config)
	if err != nil {
		return err
	}

	return nil
}

func ChangeListFile(fileName string) error {
	config, err := configs.UnmarshingConfigs()
	if err != nil {
		return err
	}

	if filepath.Ext(fileName) != ".json" {
		fileName = fileName + ".json"
	}

	path, err := configs.GetConfigPath(fileName)
	if err != nil {
		return err
	}

	_, err = os.Stat(path) // checks if file exists
	if err != nil {
		fmt.Println("teste") // debug
		return fmt.Errorf("'%s' not found, must be a json file: ", fileName)
	}

	config.CurrentListPath = path
	fmt.Printf("List file changed successfully: %v", config.CurrentListPath)

	err = configs.MarshingConfigs(config)
	if err != nil {
		return err
	}

	return nil
}
