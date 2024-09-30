package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: taskcli [list|add|complete|delete] [arguments]")
		return
	}

	loadTasks()

	switch os.Args[1] {
	case "list":
		listTasks()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskcli add [task title]")
			return
		}
		addTask(os.Args[2])
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskcli complete [task ID]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		completeTask(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: taskcli delete [task ID]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		deleteTask(id)
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
