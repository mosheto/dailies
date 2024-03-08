/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add title",
	Short: "A brief description of your command",
	Long:  `Add a task to the list of tasks to be done.`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		taskTitle := args[0]

		tasks := TaskList{}
		tasks.Load("tasks.json")
		tasks.Add(taskTitle)
		tasks.Store("tasks.json")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
