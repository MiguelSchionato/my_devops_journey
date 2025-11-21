package commands

import (
	"fmt"
	"os"
	"todo/configs"
	"todo/templates"
)

func ChangeList(fileName string) error {
	var config templates.Config
	config, err := configs.UnmarshingConfigs()
	if err != nil {
		return err
	}

	_, err = os.Stat(fileName) // checks if file exists
	if err != nil {
		fmt.Println("teste")
		return fmt.Errorf("'%s' not found, must be a json file: ", fileName)
	}

	config.CurrentList = fileName
	fmt.Printf("List changed successfully: %v", config.CurrentList)

	err = configs.MarshingConfigs(config)
	if err != nil {
		return err
	}

	return nil
}
