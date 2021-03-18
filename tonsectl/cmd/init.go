package cmd

import (
    "github.com/INTONNATION/tonos-se-installers/tonsectl/app/tonseapi"
    "github.com/spf13/cobra"
    "fmt"
    "os"
    "strings"
    "os/exec"
)

func init() {
  rootCmd.AddCommand(initCmd)
  rootCmd.AddCommand(apiCmd)
}

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Start TONSE API service",
    Run: func(cmd *cobra.Command, args []string) {
         api()
    },
}

var apiCmd = &cobra.Command{
    Use:   "api",
    Short: "Start TONSE API service",
    Run: func(cmd *cobra.Command, args []string) {
         tonseapi.RunApi()
    },
}


func api() {
       if strings.ToLower(os.Args[1]) == "init" {
           cmd := exec.Command(os.Args[0], "api")
           cmd.Start()
           fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
           os.Exit(0)
}
}
