package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "taskmanager",
	Short: "Task Manager CLI Application",
	Long:  "Task Manager is a simple CLI tool to manage your tasks,  supporting add, list, mark done, and delete operations.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(addCMD)
	rootCmd.AddCommand(deleteCMD)
	rootCmd.AddCommand(listCMD)
	rootCmd.AddCommand(doneCMD)

}
