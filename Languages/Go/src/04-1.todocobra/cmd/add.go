/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	"todoCobra/configs"
	"todoCobra/logic"
	"todoCobra/templates"

	"github.com/spf13/cobra"
)

// addCmd represents the add command

var taskName, taskDescr, taskDueDateStr string
var taskDueDate = time.Now().AddDate(0, 0, 7)
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command is used to add either a new task or a new list of tasks",
	Long: `

	Usage:
	Use this command with either the flag --task, the short version -t or without any flag to add a new task to the default tasklist.
	You can also use the flag --list, or the short version -l, to add a new list of tasks on the same file.

	`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Task added successfully")

		if len(args) > 0 && len(taskName) == 0 {
			taskName = args[0]
		}

		var allLists []templates.Lists
		allLists, err := logic.UnmarshingCurrentFile()
		if err != nil {
			return
		}

		ConfigFile, err := configs.CurrentConfig()
		if err != nil {
			return
		}

		CurrentList := ConfigFile.CurrentList
		listIndex := logic.FindList(CurrentList, allLists)

		var maxID int
		for _, i := range allLists[listIndex].Tasks {
			if i.ID > maxID {
				maxID = i.ID
			}
		}

		if len(taskDueDateStr) != 0 {
			taskDueDate, err = time.Parse("2006-01-02", taskDueDateStr)
			if err != nil {
				return
			}
		}

		newTask := templates.Task{
			Name:  taskName,
			Date:  time.Now(),
			Descr: taskDescr,
			Due:   taskDueDate,
			ID:    maxID + 1,
		}

		if listIndex != -1 {
			allLists[listIndex].Tasks = append(allLists[listIndex].Tasks, newTask)
		} else {
			newDefaultList := templates.Lists{
				Name:  CurrentList,
				Descr: "Default list",
				Tasks: []templates.Task{newTask},
			}

			allLists = append(allLists, newDefaultList)
			fmt.Printf("List not found, appending to the new list") // debug
		}

		err = logic.MarshingToCurrentFile(allLists)
		if err != nil {
			return
		}
	},
}

func init() {

	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&taskName, "task", "t", "", "Task name")
	addCmd.Flags().StringVarP(&taskDescr, "descr", "d", "", "Task description")
	// addCmd.Flags().AddAlias("descr", "description")
	addCmd.Flags().StringVar(&taskDueDateStr, "due", "", "Task due date")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
