package cmd

import (
    "fmt"
    "os/exec"
    "log"
    "runtime"
    "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
    Use:   "start",
    Short: "Start service",
    Run: func(cmd *cobra.Command, args []string) {
       start()
    },

}

func start() {
    // Run this Python program from Go.
    cmd := exec.Command("python3", "./test.py")
    if runtime.GOOS == "windows" {
	cmd = exec.Command("python3", "./test.py")
	}
    fmt.Println("Running graphql")
    // Wait for the Python program to exit.
    err := cmd.Start()
    fmt.Println("GraphQL is running:", err)
    if err != nil {
	log.Fatalf("GraphQL failed with %s\n", err)
	}

}
