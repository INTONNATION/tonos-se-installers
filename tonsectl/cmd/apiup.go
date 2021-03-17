package cmd

import (
    "github.com/INTONNATION/tonos-se-installers/tonsectl/app/tonseapi"
    "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(apiupCmd)
}

var apiupCmd = &cobra.Command{
    Use:   "apiup",
    Short: "Start TON apiup service",
    Run: func(cmd *cobra.Command, args []string) {
       apiup()
    },
}

func apiup() {
    tonseapi.RunApi()
}

