package cmd

import (
    "github.com/INTONNATION/tonos-se-installers/tonsectl/app/tonseapi"
    "github.com/spf13/cobra"
    "fmt"
    "os"
    "strings"
    "os/exec"
    "log"
    "strconv"
    "syscall"
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

var PIDFile = "./.daemonize.pid"

 func savePID(pid int) {

         file, err := os.Create(PIDFile)
         if err != nil {
                 log.Printf("Unable to create pid file : %v\n", err)
                 os.Exit(1)
         }

         defer file.Close()

         _, err = file.WriteString(strconv.Itoa(pid))

         if err != nil {
                 log.Printf("Unable to create pid file : %v\n", err)
                 os.Exit(1)
         }

         file.Sync() // flush to disk

 }



func api() {
       if strings.ToLower(os.Args[1]) == "init" {
           cmd := exec.Command(os.Args[0], "api")
           cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
           cmd.Start()
           fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
           savePID(cmd.Process.Pid)
           os.Exit(0)
}
}
