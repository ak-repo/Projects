package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var addCMD = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add new task ",
	Long:  "Add a new task with a description to your task list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		tasks := loadTasks()
		tasks = append(tasks, Task{Description: description, Done: false})
		saveTasks(tasks)
		fmt.Printf("Task added: %s\n", description)
	},
}

func loadTasks() []Task {
	file, err := os.Open("tasks.json")
	if os.IsNotExist(err) {
		return []Task{}
	} else if err != nil {
		panic(err)
	}
	defer file.Close()

	var task []Task
	json.NewDecoder(file).Decode(&task)
	return task
}

func saveTasks(tasks []Task) {
	file, err := os.Create("tasks.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	json.NewEncoder(file).Encode(tasks)
}
