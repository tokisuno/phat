package cmd 

import (
	"fmt"
	"os"
	_ "time"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command {
    Use: "phat",
    Short: "Multipurpose productivity tool",
    Long: "(P)omodoro (Habit-tracker) and (T)ime tracker tool",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello world")
    },
}
func Execute() {
    if err := rootCmd.Execute(); err != nil {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
    }
}
