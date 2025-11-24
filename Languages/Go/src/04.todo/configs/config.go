package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo/templates"
)

const File = "config.json"

var defaultConfig templates.Config

func CurrentConfig() (templates.Config, error) {
	config, err := UnmarshingConfigs()
	if err != nil {
		if os.IsNotExist(err) {
			defaultConfigFilePath, err := GetConfigPath("default.json")
			if err != nil {
				return templates.Config{}, err
			}
			config = templates.Config{CurrentList: "default", CurrentListPath: defaultConfigFilePath}
			err = MarshingConfigs(config)
			if err != nil {
				return templates.Config{}, err
			}
			return config, nil
		}
		return templates.Config{}, err
	}
	return config, nil
}

func UnmarshingConfigs() (templates.Config, error) {
	var config templates.Config
	path, err := GetConfigPath(File)
	if err != nil {
		return config, err
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return config, err // os.IsNotExist will be handled by CurrentConfig
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, fmt.Errorf("error on reading config file: %v", err)
	}
	return config, nil
}

func MarshingConfigs(newDefaultList templates.Config) error {
	data, err := json.Marshal(newDefaultList)
	if err != nil {
		fmt.Printf("Fail writting file: %v\n", err)
		return err
	}

	path, err := GetConfigPath(File)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0640)
	if err != nil {
		fmt.Printf("Fail writting file: %v\n", err)
		return err
	}

	return nil
}

func GetConfigPath(fileName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home folder not found %v", err)
	}

	configPath := filepath.Join(home, ".config", "todo")

	err = os.MkdirAll(configPath, 0755)
	if err != nil {
		return "", fmt.Errorf("error creating parent folder %v", err)
	}

	fullPath := filepath.Join(configPath, fileName)
	return fullPath, nil

}
