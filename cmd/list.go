package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := viper.GetString("todo file") 

		var todos []string

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading from file %v\n", err)
			os.Exit(2)
		}

		err = json.Unmarshal(data, &todos)
		if err != nil {
			fmt.Printf("Error converting from json %v\n", err)
			os.Exit(2)
		}

		for index, todo := range todos {
			fmt.Printf("%d: %s\n", index, todo)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}