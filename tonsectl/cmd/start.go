package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "io/ioutil"
    "net/http"
    "os"
    "os/exec"
    "strconv"
)

func init() {
  rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
    Use:   "start",
    Short: "Start service",
    Run: func(cmd *cobra.Command, args []string) {
       preInit()
       start()
    },

}

func preInit() {
    if _, err := os.Stat(PIDFile); os.IsNotExist(err) {
        fmt.Printf("Running init process\n")
        cmd := exec.Command(os.Args[0], "init")
        cmd.Start()
        cmd.Wait()
        }
    }

func start() {
    data, err := ioutil.ReadFile(PIDFile)
    ProcessID, err := strconv.Atoi(string(data))
    if err != nil {
       fmt.Printf("Unable to get process ID [%v] with error %v \n", ProcessID, err)
       os.Exit(1)
    }
    resp, err := http.Head("http://localhost:10000/tonse/start")
    if err != nil {
        fmt.Printf("Start failed\n")
    } else {
        defer resp.Body.Close()
    }
}
