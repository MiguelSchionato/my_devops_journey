/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todoCobra/configs"
	"todoCobra/logic"

	"github.com/spf13/cobra"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Bold   = "\033[1m"
	Strike = "\033[9m"
)

var configFile, err = configs.CurrentConfig()
var allLists, err1 = logic.UnmarshingCurrentFile()
var listName = configFile.CurrentList
var listIndex = logic.FindList(listName, allLists)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			err := ListTaskName()
			if err != nil {
				return
			}
		} else {
			err := ListTask(args[0])
			if err != nil {
				return
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListTaskName() error {

	if err != nil {
		return err
	}

	if err1 != nil {
		return err1
	}

	for _, task := range allLists[listIndex].Tasks {
		if task.State == 1 {
			fmt.Printf("• %s%s%s%s\n", Green, Strike, task.Name, Reset)

		} else {
			fmt.Printf("•%d %s\n", task.ID, task.Name)
		}
	}

	return nil
}

func ListTask(listID string) error {
	if err != nil {
		return err
	}

	if err1 != nil {
		return err1
	}

	if listIndex == -1 {
		fmt.Println("task not found")
		return err
	}

	taskIndex := logic.FindTask(listIndex, listID, allLists)
	list := allLists[listIndex].Tasks

	fmt.Printf("• %s\n%s\n%s\n", list[taskIndex].Name, list[taskIndex].Descr, list[taskIndex].Due.Format("01/02/2006"))

	return nil
}
