package cmd

import (
	"fmt"

	"github.com/leonard-ladva/todoer/operations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := viper.GetString("file")

		var todos = operations.List(filePath)

		fmt.Println(filePath)
		for index, todo := range todos {
			fmt.Printf("%d: %s\n", index, todo)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
