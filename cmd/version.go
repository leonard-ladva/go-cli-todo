package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("todoer %s\n", version)
		fmt.Println("Build date:", buildDate)
	},
}

var buildDate string

func init() {
	rootCmd.AddCommand(versionCmd)
}