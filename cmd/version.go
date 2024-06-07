package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
    Use: "version",
    Short: "Prints version",
    Long: "All software has versions. This is phat's",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("phat -> version 0.1") 
    },
}
