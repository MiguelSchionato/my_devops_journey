/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todoCobra/logic"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done called")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const (
	StatusPending = 0
	StatusDone    = 1
)

func Done(taskName, listName string) error {
	allLists, err := logic.UnmarshingCurrentFile()
	if err != nil {
		fmt.Printf("erro %v", err)
		return err
	}

	listIndex := logic.FindList(listName, allLists)
	if listIndex == -1 {
		return err
	}

	targetList := &allLists[listIndex]

	taskIndex := logic.FindTask(listIndex, taskName, allLists)
	if taskIndex == -1 {
		return err
	}
	targetTask := &targetList.Tasks[taskIndex]

	targetTask.State = StatusDone
	err = logic.MarshingToCurrentFile(allLists)
	if err != nil {
		return err
	}
	return nil
}
