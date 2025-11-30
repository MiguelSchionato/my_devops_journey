/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todoCobra/configs"
	"todoCobra/logic"
	// "todoCobra/templates"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
	`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm called")
		allLists, err := logic.UnmarshingCurrentFile()
		if err != nil {
			return
		}

		currentConfig, err := configs.CurrentConfig()
		if err != nil {
			return
		}

		currentList := currentConfig.CurrentList
		listIndex := logic.FindList(currentList, allLists)

		if listIndex == -1 {
			fmt.Print("Task not found")
			return
		}

		err = logic.MarshingToCurrentFile(allLists)
		if err != nil {
			return
		}

		if args[0] == "all" {
			var maxID int
			for _, i := range allLists[listIndex].Tasks {
				if i.ID > maxID {
					maxID = i.ID
				}
			}
		}

		taskIndex := logic.FindTask(listIndex, args[0], allLists)

		if taskIndex != -1 { // Removing logic
			previousTasks := allLists[listIndex].Tasks[:taskIndex]
			nextTasks := allLists[listIndex].Tasks[taskIndex+1:]
			allLists[listIndex].Tasks = append(previousTasks, nextTasks...)

			err := logic.MarshingToCurrentFile(allLists)
			if err != nil {
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
