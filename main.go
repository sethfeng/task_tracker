package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending", "in progress", "done"
}

var tasks []Task

const tasksFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-tracker <command> [arguments]")
		fmt.Println("Available commands: add, list, update, delete, mark")
		return
	}

	loadTasks() // Load tasks from file on start-up

	switch os.Args[1] {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "update":
		handleUpdate()
	case "delete":
		handleDelete()
	case "mark":
		handleMark()
	default:
		fmt.Println("Unknown command")
	}
}

// loadTasks Function to load tasks from file
func loadTasks() {
	file, err := os.ReadFile(tasksFile)
	if os.IsNotExist(err) {
		return // File doesn't exist, no tasks to load
	}
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Println("Error decoding tasks:", err)
	}
}

// saveTasks Function to save tasks to file
func saveTasks() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
		return
	}
	if err := os.WriteFile(tasksFile, data, 0644); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

// handleAdd Function to handle adding a new task
func handleAdd() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: task-tracker add <title> <description>")
		return
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	tasks = append(tasks, Task{
		ID:          newID,
		Title:       os.Args[2],
		Description: os.Args[3],
		Status:      "pending",
	})

	saveTasks() // Save tasks to file after adding a new task
	fmt.Println("Task added successfully! ID: ", newID)
}

// handleUpdate Function to handle updating a task
func handleUpdate() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: task-tracker update <id> <title> <description>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = os.Args[3]
			tasks[i].Description = os.Args[4]
			saveTasks() // Save tasks to file after updating a task
			fmt.Println("Task updated successfully! ID: ", id)
			return
		}
	}

	fmt.Println("Task not found, ID: ", id)
}

// handleDelete Function to handle deleting a task
func handleDelete() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task-tracker delete <id>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks() // Save tasks to file after deleting a task
			fmt.Println("Task deleted successfully! ID: ", id)
			return
		}
	}

	fmt.Println("Task not found, ID: ", id)
}

// handleMark Function to handle marking a task as status
func handleMark() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: task-tracker mark <id> <status>(pending, in progress, done)")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	status := os.Args[3]
	if status != "pending" && status != "in progress" && status != "done" {
		fmt.Println("Invalid status")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			saveTasks() // Save tasks to file after marking a task as status
			fmt.Println("Task marked as status successfully! ID: ", id)
		}
	}

	fmt.Println("Task not found, ID: ", id)
}

// handleList Function to handle listing all tasks
func handleList() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	filter := ""
	if len(os.Args) > 2 {
		filter = os.Args[2]
	}

	switch filter {
	case "done":
		printTasks(filterByStatus("done"))
	case "in-progress":
		printTasks(filterByStatus("in progress"))
	case "not-done":
		printTasks(filterByStatusNot("done"))
	default:
		printTasks(tasks)
	}
}

// filterByStatus Function to filter tasks by status
func filterByStatus(status string) []Task {
	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

// filterByStatusNot Function to filter tasks by status not equal to status
func filterByStatusNot(status string) []Task {
	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status != status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

// printTasks Function to print tasks in a table format
func printTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Printf("%-5s %-20s %-40s %-10s\n", "ID", "Title", "Description", "Status")
	fmt.Println("--------------------------------------------------------------------------------")
	for _, task := range tasks {
		fmt.Printf("%-5d %-20s %-40s %-10s\n", task.ID, task.Title, task.Description, task.Status)
		fmt.Println("--------------------------------------------------------------------------------")
	}
}
