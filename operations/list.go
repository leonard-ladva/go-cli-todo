package operations

import (
	"encoding/json"
	"fmt"
	"os"
)

func List(filePath string) []string {
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

	return todos
}
