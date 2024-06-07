package cmd 

import (
	"fmt"
	"os"
	_ "time"

	"github.com/spf13/cobra"
)

func helloWorld() string {
    return "Hello world" 
}

var (
    cfgFile string
	userLicense string
)

var rootCmd = &cobra.Command {
    Use: "phat",
    Short: "Multipurpose productivity tool",
    Long: "(P)omodoro (Habit-tracker) and (T)ime tracker tool",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Fprintf(cmd.OutOrStdout(), "%s", helloWorld())
    },
}
func Execute() {
    if err := rootCmd.Execute(); err != nil {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
    }
}
