package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use: "remove index",
	Short: "Remove a task",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := viper.GetString("todo file") 
		index, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Printf("Error converting index %v\n", err)
			os.Exit(2)
		}

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

		todos = slices.Concat(todos[0:index],todos[index+1:])
		
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
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}