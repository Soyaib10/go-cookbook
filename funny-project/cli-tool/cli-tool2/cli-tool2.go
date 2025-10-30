package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var priority int

	rootCmd := &cobra.Command{
		Use:   "tasks",
		Short: "Task management CLI.",
	}

	addCmd := &cobra.Command{
		Use:   "add [task]",
		Short: "Add a new task.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Added task: %s (Priority: %d)\n", args[0], priority)
		},
	}

	addCmd.Flags().IntVarP(&priority, "priority", "p", 1, "Task priority (1-5)")
	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
