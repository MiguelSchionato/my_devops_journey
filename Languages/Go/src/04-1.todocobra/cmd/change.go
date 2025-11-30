/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"todoCobra/configs"

	"github.com/spf13/cobra"
)

var changeList string
var changeFile string

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("change called")

		err = ChangeList(changeList)
		if err != nil {
			return
		}

		err = ChangeListFile(changeFile)
		if err != nil {
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
	changeCmd.Flags().StringVarP(&changeList, "list", "l", "", "Change list on the same file")
	changeCmd.Flags().StringVarP(&changeFile, "file", "f", "", "Change file list")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
