package cmd

import (
    "github.com/spf13/cobra"
     "net/http"
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
        resp, err := http.Head("http://localhost:10000/tonse/start")
        if err != nil {
    	// handle err
        }
        defer resp.Body.Close()
    }
