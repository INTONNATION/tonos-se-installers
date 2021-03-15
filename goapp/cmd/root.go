package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "tonos_installer",
  Short: "tonos_installer",
  Long: "tonos_installer log description",
  Version: "0.0.1",
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}