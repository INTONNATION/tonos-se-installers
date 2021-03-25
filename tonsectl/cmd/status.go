package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of the service",
	Run: func(cmd *cobra.Command, args []string) {
		status()
	},

}

func status() {
	if _, err := os.Stat(PIDFile); err == nil {
		data, err := ioutil.ReadFile(PIDFile)
		if err != nil {
			log.Fatal("Not running \n")
			os.Exit(1)
		}
		ProcessID, err := strconv.Atoi(string(data))
		fmt.Printf("Process ID: [%v]\n", ProcessID)
		os.Exit(1)
	}
	fmt.Print("Unable to find process ID \n")
}