package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataFile = "tasks.json"

// Load tasks from the JSON file
func loadTasks() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{}
		} else {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
	} else {
		err = json.Unmarshal(file, &tasks)
		if err != nil {
			fmt.Println("Error parsing file:", err)
			os.Exit(1)
		}
	}
}

// Save tasks to the JSON file
func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}
	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}
