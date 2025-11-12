package cmd

import (
	"github.com/leonard-ladva/todoer/operations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add todo",
	Short: "Add a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo := args[0]
		operations.Add(todo)

		if viper.GetBool("list after add") {
			listCmd.Run(listCmd, nil)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
