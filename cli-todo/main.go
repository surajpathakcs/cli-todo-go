// This Go program is a simple command-line todo list application that allows users to add, list, mark
// tasks as done, and delete tasks.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// The Task type represents a task with a title and a status indicating whether it is done or not.
// @property {string} Title - The `Title` property in the `Task` struct represents the title or name of
// the task. It is a string type field that is tagged with `json:"title"` for JSON marshaling and
// unmarshaling purposes.
// @property {bool} Done - The `Done` property in the `Task` struct is a boolean type that represents
// whether the task has been completed or not. If `Done` is `true`, it means the task has been
// completed; if `Done` is `false`, it means the task is still pending or not completed
type Task struct {
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}

const dataFile = "data.json"

// The function `loadTasks` reads tasks from a data file and returns them as a slice of Task structs.
func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(dataFile)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
	return tasks, nil
}

// The `saveTasks` function serializes a slice of Task structs into JSON format and writes it to a
// file.
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

// The main function in this Go program implements a simple to-do list application with commands for
// adding, listing, marking tasks as done, and deleting tasks.
func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: todo [add|list|done|delete] [task]")
		return
	}

	cmd := args[1]
	tasks, _ := loadTasks()

	switch cmd {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a task title.")
			return
		}
		task := Task{Title: args[2], Done: false}
		tasks = append(tasks, task)
		saveTasks(tasks)
		fmt.Println("Task added.")
	case "list":
		for i, task := range tasks {
			status := "❌"
			if task.Done {
				status = "✅"
			}
			fmt.Printf("%d. %s %s\n", i+1, status, task.Title)
		}
	case "done":
		if len(args) < 3 {
			fmt.Println("Provide task number to mark done.")
			return
		}
		index, _ := strconv.Atoi(args[2])
		if index <= len(tasks) {
			tasks[index-1].Done = true
			saveTasks(tasks)
			fmt.Println("Marked as done.")
		}
	case "delete":
		if len(args) < 3 {
			fmt.Println("Provide task number to delete.")
			return
		}
		index, _ := strconv.Atoi(args[2])
		if index <= len(tasks) {
			tasks = append(tasks[:index-1], tasks[index:]...)
			saveTasks(tasks)
			fmt.Println("Task deleted.")
		}
	default:
		fmt.Println("Unknown command.")
	}
}
