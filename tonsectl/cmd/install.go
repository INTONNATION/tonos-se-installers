package cmd

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func init() {
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},

}
//go:embed init-scripts
var f embed.FS

func install() {
        var data []uint8
        if runtime.GOOS == "darwin" {
            data, _ = f.ReadFile("init-scripts/init.mac.sh")
        }
        if runtime.GOOS == "linux" {
	    data, _ = f.ReadFile("init-scripts/init.linux.sh")
        }
        if runtime.GOOS == "windows" {
	    data, _ = f.ReadFile("init-scripts/init.windows.bat")
        }
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = strings.NewReader(string(data))
	//fmt.Printf("Running in background init script")
	//f, err := os.OpenFile("./APIlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Printf("error opening file: %v", err)
	//}
	//defer f.Close()
	//cmd.Stdout = f
	//cmd.Stderr = f
	fmt.Printf("Start installation")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
}
