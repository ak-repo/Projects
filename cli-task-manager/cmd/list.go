package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCMD = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long:  "Display all tasks with their status (pending or done)",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		if len(tasks) == 0 {

			fmt.Println("No tasks found")
			return
		}

		for i, task := range tasks {
			status := "pending"
			if task.Done {
				status = "Done"
			}
			fmt.Printf("%d. %s [%s] \n", i+1, task.Description, status)
		}

	},
}
