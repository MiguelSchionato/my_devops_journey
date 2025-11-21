package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"todo/templates"
)

const File = "config.json"

var defaultConfig templates.Config

func CurrentList() (string, error) {
	config, err := UnmarshingConfigs()
	if os.IsNotExist(err) || err != nil { // best err case, os.IsNotExist.
		config := templates.Config{CurrentList: "default.json"} // using a default list
		defaultList := config.CurrentList
		return defaultList, nil
	}

	return config.CurrentList, nil
}

func UnmarshingConfigs() (templates.Config, error) {
	bytes, err := os.ReadFile(File)
	if os.IsNotExist(err) {
		config := templates.Config{CurrentList: "default.json"} // using a default list
		return config, nil
	} else if err != nil {
		return defaultConfig, err
	}

	err = json.Unmarshal(bytes, &defaultConfig)
	if err != nil {
		err = errors.New("error on reading config file, using default list")
		return defaultConfig, err
	}

	return defaultConfig, nil
}

func MarshingConfigs(newDefaultList templates.Config) error {
	data, err := json.Marshal(newDefaultList)
	if err != nil {
		fmt.Printf("Fail writting file: %v\n", err)
		return err
	}

	err = os.WriteFile(File, data, 0640)
	if err != nil {
		fmt.Printf("Fail writting file: %v\n", err)
		return err
	}

	return nil
}
