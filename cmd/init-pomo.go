package cmd 

import (
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(pomoCmd) 
}

var pomoCmd = &cobra.Command {
    Use: "pomo",
    Short: "Pomodoro functionality",
    Long: "Pomodoro functionality + time tracker",
    Run: func(cmd *cobra.Command, args []string) {
        startPomo() 
    },
}
