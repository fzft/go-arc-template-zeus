package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(serverCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error executing command: %v", err)
	}
}
