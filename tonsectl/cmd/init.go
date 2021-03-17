package cmd

import (
    "github.com/INTONNATION/tonos-se-installers/tonsectl/app/tonseapi"
    "github.com/spf13/cobra"
    "fmt"
    "os"

    "os/exec"
)

func init() {
  rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Start TONSE API service",
    Run: func(cmd *cobra.Command, args []string) {
       main()
    },
}




func main() {
    tonseapi.RunApi()
    cmd := exec.Command(os.Args[1], "main")
    cmd.Start()
    fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
}

