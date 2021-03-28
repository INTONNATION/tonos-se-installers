package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset directory with local node",
	Run: func(cmd *cobra.Command, args []string) {
                stopReset()
		reset()
                startReset()
	},
}

func reset(){
	resp, err := http.Head("http://localhost:10000/tonse/reset")
	if err != nil {
		fmt.Printf("Reset failed")
	}
	defer resp.Body.Close()
}

func stopReset(){
	resp, err := http.Head("http://localhost:10000/tonse/stop")
	if err != nil {
		fmt.Printf("Stop failed")
	}
	defer resp.Body.Close()
}

func startReset(){
	resp, err := http.Head("http://localhost:10000/tonse/start")
	if err != nil {
		fmt.Printf("Start failed")
	}
	defer resp.Body.Close()
}
