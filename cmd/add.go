package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use: "add todo",
	Short: "Add a task",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := viper.GetString("todo file") 
		todo := args[0]

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

		todos = append(todos, todo)

		newData, err := json.MarshalIndent(todos, "", " ")
		if err != nil {
			fmt.Printf("Error converting to json %v\n", err)
			os.Exit(2)
		}
		err = os.WriteFile(filePath, newData, 0644)
		if err != nil {
			fmt.Printf("Error converting to json %v\n", err)
			os.Exit(2)
		}
		// file, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		// if err != nil {
		// 	fmt.Printf("Error opening file %v\n", err)
		// 	os.Exit(2)
		// }
		// defer file.Close()

		// if _, err = file.WriteString(todo); err != nil {
		// 	fmt.Printf("Error writing to file %v\n", err)
		// 	os.Exit(2)
		// }

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}