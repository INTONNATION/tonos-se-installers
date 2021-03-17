package cmd

import (
    "github.com/spf13/cobra"
    "net/http"
)

func init() {
  rootCmd.AddCommand(graphqlCmd)
}

var graphqlCmd = &cobra.Command{
    Use:   "graphql",
    Short: "Start TON q server service",
    Run: func(cmd *cobra.Command, args []string) {
       graphqlRun()
    },

}

func graphqlRun() {
        resp, err := http.Head("http://localhost:10000/tonsectl/start")
        if err != nil {
    	// handle err
        }
        defer resp.Body.Close()
    }
