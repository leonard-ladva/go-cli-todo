package operations

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

func Remove(index int, filePath string) {
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

	todos = slices.Concat(todos[0:index], todos[index+1:])

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
}
