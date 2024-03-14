/*
Copyright © 2024 mosheto <msalshakhatreh@gmail.com>
*/
package cmd

import (
	"os"

	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dailies",
	Short: "Keep track of your daily tasks",
	Long:  `Follow your daily tasks with ease and track your progress`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Something went wrong when executing the command: ", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.TodoApp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
