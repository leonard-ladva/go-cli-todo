package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/leonard-ladva/todoer/operations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use:   "remove index",
	Short: "Remove a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := viper.GetString("todo file")
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error converting index %v\n", err)
			os.Exit(2)
		}
		operations.Remove(index, filePath)

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
