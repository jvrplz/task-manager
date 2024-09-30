package main

import "fmt"

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

var tasks []Task

// List all tasks
func listTasks() {
	fmt.Println("Task List:")
	for _, task := range tasks {
		status := "❌"
		if task.Complete {
			status = "✅"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, status, task.Title)
	}
}

// Add a new task
func addTask(title string) {
	id := len(tasks) + 1
	task := Task{ID: id, Title: title, Complete: false}
	tasks = append(tasks, task)
	saveTasks()
	fmt.Println("Task added:", title)
}

// Mark a task as completed
func completeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Complete = true
			saveTasks()
			fmt.Println("Task completed:", tasks[i].Title)
			return
		}
	}
	fmt.Println("Task not found.")
}

// Delete a task
func deleteTask(id int) {
	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found.")
		return
	}

	// Remove the task from the slice
	tasks = append(tasks[:index], tasks[index+1:]...)
	saveTasks()
	fmt.Println("Task deleted successfully.")
}
