package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of tonos_installer",
  Long:  `All software has versions. This is tonos_installer`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("v0.28.5")
  },
}