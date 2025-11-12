package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var newListCmd = &cobra.Command{
	Use:   "new-list Filename",
	Short: "Make a new todo list file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		file, err := os.Create(fileName + ".json")
		if err != nil {
			fmt.Printf("Error creating file %v\n", err)
			os.Exit(2)
		}
		defer file.Close()

		file.WriteString("[]")

		fmt.Println("Successfully createad file", file.Name())
	},
}

func init() {
	rootCmd.AddCommand(newListCmd)
}
