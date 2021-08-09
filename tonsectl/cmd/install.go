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
		if len(args) == 0 {
			port="80"
			db_port="8529"
			install(port,db_port)
		}
        install(args[0],args[1])
		},

}
//go:embed init-scripts
var f embed.FS
var nodejs_version string
var tonosse_version string
var arango_version string
var qserver string
var port string
var db_port string


func install(port string, db_port string) {
	var data []uint8
	var cmd *exec.Cmd

	fmt.Printf("nodejs_version:" + nodejs_version+"\n")
	fmt.Printf("tonosse_version:" + tonosse_version+"\n")
	fmt.Printf("arango_version:" + arango_version+"\n")
	fmt.Printf("qserver:" + qserver+"\n")
	fmt.Printf("port:" + port+"\n")
        fmt.Printf("dbport:" + db_port+"\n")
        os.Mkdir(tonossePath,0755)
        if runtime.GOOS == "darwin" {
            data, _ = f.ReadFile("init-scripts/init.mac.sh")
	    os.WriteFile(tonossePath+"/install.sh", data, 0755)
	    args := []string{tonossePath+"/install.sh",nodejs_version,tonosse_version,arango_version,port,db_port}
            cmd = exec.Command("/bin/bash", args...)
        }
        if runtime.GOOS == "linux" {
	    data, _ = f.ReadFile("init-scripts/init.linux.sh")
	    os.WriteFile(tonossePath+"/install.sh", data, 0755)
	    args := []string{nodejs_version,tonosse_version,arango_version,port,db_port}
	    cmd = exec.Command(tonossePath+"/install.sh", args...)
        }
        if runtime.GOOS == "windows" {
	    data, _ = f.ReadFile("init-scripts/init.windows.bat")
	    os.WriteFile(tonossePath+"/install.bat", data, 0755)
	    args := []string{"/c", tonossePath+"/install.bat", nodejs_version,tonosse_version,arango_version,qserver,port,db_port}
	    cmd = exec.Command("cmd", args...)
        }
	fmt.Printf("Start installation\n")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
}
