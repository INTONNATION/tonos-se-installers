// +build windows

package cmd

import (
    "fmt"
    "github.com/INTONNATION/tonos-se-installers/tonsectl/app/tonseapi"
    "github.com/spf13/cobra"
    "log"
    "os"
    "os/exec"
    "os/user"
    "strconv"
    "strings"
    "syscall"
    "time"
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

var usr, e = user.Current()
var tonossePath = usr.HomeDir + "/tonse/"

var PIDFile = tonossePath+".daemonize.pid"

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
           cmd.SysProcAttr = &syscall.SysProcAttr{
               CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
           }
           cmd.Start()
           fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
           savePID(cmd.Process.Pid)
           time.Sleep(time.Second * 5)
           os.Exit(0)
       }
}
