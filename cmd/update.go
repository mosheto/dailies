/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// updateCmd represents the take command
var updateCmd = &cobra.Command{
	Use:   "update id",
	Short: "Update a task",
	Long:  `Update a task by adding an update and optionally changing the status`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		tasks := TaskList{}
		tasks.Load("tasks.json")

		id, _ := strconv.Atoi(args[0])
		status, _ := cmd.Flags().GetString("status")
		description, _ := cmd.Flags().GetString("description")

		if description != "" {
			tasks.Update(id, description)
		}

		if status != "" {
			task := tasks.Get(id)
			task.Status = StatusType(status)

			if task.Status == Done {
				task.CompletedAt = time.Now()
			}
		}

		tasks.Store("tasks.json")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("status", "s", "", "Status of the task")
	updateCmd.Flags().StringP("description", "d", "", "Description of the update")
}
