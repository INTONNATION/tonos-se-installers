package cmd

import (
    "github.com/spf13/cobra"
    "net/http"
)

func init() {
  rootCmd.AddCommand(nodeCmd)
}

var nodeCmd = &cobra.Command{
    Use:   "node",
    Short: "Start TON node service",
    Run: func(cmd *cobra.Command, args []string) {
       node()
    },
}

func node() {
    resp, err := http.Head("http://localhost:10000/tonsectl/start")
    if err != nil {
	// handle err
    }
    defer resp.Body.Close()
}

