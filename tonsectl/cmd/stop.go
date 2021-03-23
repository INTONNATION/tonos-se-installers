package cmd

import (
    "github.com/spf13/cobra"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "syscall"
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
    if _, err := os.Stat(PIDFile); err == nil {
        data, err := ioutil.ReadFile(PIDFile)
        if err != nil {
            log.Fatal("Not running")
            os.Exit(1)
        }
        ProcessID, err := strconv.Atoi(string(data))

        if err != nil {
            log.Fatal("Unable to read and parse process id found in ", PIDFile)
            os.Exit(1)
        }
        if err != nil {
            log.Fatal("Unable to find process ID [%v] with error %v \n", ProcessID, err)
            os.Exit(1)
        }
        // remove PID file
        os.Remove(PIDFile)

        log.Printf("Killing process ID [%v] now.\n", ProcessID)
        // kill process and exit immediately
        //err = process.Kill()
        err =syscall.Kill(-ProcessID, syscall.SIGKILL)
        if err != nil {
            log.Fatal("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
            os.Exit(1)
        } else {
            log.Printf("Killed process ID [%v]\n", ProcessID)
            os.Exit(0)
        }
    }
}

