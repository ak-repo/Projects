package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCMD = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
	Long:  "Delete a task by specifying its task number.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("Invalid task number.")
			return
		}

		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Println("Task number out of range.")
			return
		}

		tasks = append(tasks[:index-1], tasks[index:]...)
		saveTasks(tasks)
		fmt.Println("task deleted")

	},
}
