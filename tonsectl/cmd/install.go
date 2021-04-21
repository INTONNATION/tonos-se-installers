package cmd

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
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
var nodejs_version string
var tonosse_version string
var arango_version string
var qserver string

func install() {
        var data []uint8
	var cmd *exec.Cmd
	fmt.Printf("nodejs_version:" + nodejs_version+"\n")
	fmt.Printf("tonosse_version:" + tonosse_version+"\n")
	fmt.Printf("arango_version:" + arango_version+"\n")
	fmt.Printf("qserver:" + qserver+"\n")
        os.Mkdir(tonossePath,0755)
        if runtime.GOOS == "darwin" {
            data, _ = f.ReadFile("init-scripts/init.mac.sh")
	    os.WriteFile(tonossePath+"/install.sh", data, 0755)
	    args := []string{tonossePath+"/install.sh",nodejs_version,tonosse_version,arango_version}
            cmd = exec.Command("/bin/bash", args...)
        }
        if runtime.GOOS == "linux" {
	    data, _ = f.ReadFile("init-scripts/init.linux.sh")
	    os.WriteFile(tonossePath+"/install.sh", data, 0755)
	    args := []string{nodejs_version,tonosse_version,arango_version}
	    cmd = exec.Command(tonossePath+"/install.sh", args...)
        }
        if runtime.GOOS == "windows" {
	    data, _ = f.ReadFile("init-scripts/init.windows.bat")
	    os.WriteFile(tonossePath+"/install.bat", data, 0755)
	    args := []string{nodejs_version,tonosse_version,arango_version,qserver}
	    cmd = exec.Command(tonossePath+"/install.bat", args...)
        }
	//fmt.Printf("Running in background init script")
	//f, err := os.OpenFile("./APIlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Printf("error opening file: %v", err)
	//}
	//defer f.Close()
	//cmd.Stdout = f
	//cmd.Stderr = f
	fmt.Printf("Start installation\n")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
}
