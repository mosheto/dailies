/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `List tasks All tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := TaskList{}
		tasks.Load("tasks.json")
		tasks.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
