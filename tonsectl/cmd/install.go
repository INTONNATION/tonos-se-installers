package cmd

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
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
	data, _ := f.ReadFile("init-scripts/init.mac.sh")
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = strings.NewReader(string(data))
	fmt.Printf("Running in background init script")
	f, err := os.OpenFile("./APIlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	cmd.Stdout = f
	cmd.Stderr = f
	cmd.Start()
	cmd.Wait()
}
