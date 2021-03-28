package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "tonsectl",
  Short: "TON OS SE installer",
  Long:  "Cross platform TON OS SE installer",
  Version: "0.0.1",
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}