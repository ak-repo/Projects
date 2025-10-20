package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCMD = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark a task as done",
	Long:  "Mark a task as done by specifying its task number.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("Invalid task number")
			return
		}

		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Println("Task number out of range.")
			return
		}

		tasks[index-1].Done = true
		str := tasks[index-1].Description
		saveTasks(tasks)
		fmt.Printf("task %d makred as done. task: %s \n", index, str)
	},
}
