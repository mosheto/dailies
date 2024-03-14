/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := TaskList{}
		tasks.Load("tasks.json")

		if len(args) == 0 {
			tasks.List()
			return
		}

		reportType := args[0]

		if reportType == "standup" {

			filteredTasks := TaskList{}
			for _, task := range tasks {
				if task.Status == Done && isTodayOrYesterday(task.CompletedAt) {
					filteredTasks.AddTask(task)
				} else if task.Status == InProgress {

					filteredUpdates := []TaskUpdate{}
					for _, update := range task.Updates {
						if isTodayOrYesterday(update.WrittenAt) {
							filteredUpdates = append(filteredUpdates, update)
						}
					}

					if len(filteredUpdates) > 0 {
						task.Updates = filteredUpdates
						filteredTasks.AddTask(task)
					}
				}
			}

			filteredTasks.List()
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
