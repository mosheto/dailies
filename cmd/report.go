/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := TaskList{}
		tasks.Load("tasks.json")

		for _, task := range tasks {
			if task.Status == Done && isTodayOrYesterday(task.CompletedAt) {
				fmt.Println(task.Title)
			}
		}
	},
}

func isTodayOrYesterday(t time.Time) bool {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)

	if t.Year() == now.Year() && t.YearDay() == now.YearDay() {
		return true
	} else if t.Year() == yesterday.Year() && t.YearDay() == yesterday.YearDay() {
		return true
	} else {
		return false
	}
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
