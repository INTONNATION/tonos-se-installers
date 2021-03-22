package cmd

import (
    "github.com/spf13/cobra"
     "net/http"
)

func init() {
  rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
    Use:   "stop",
    Short: "Stop service",
    Run: func(cmd *cobra.Command, args []string) {
       stop()
    },

}

func stop() {
        resp, err := http.Head("http://localhost:10000/tonse/stop")
        if err != nil {
    	// handle err
        }
        defer resp.Body.Close()
    }
