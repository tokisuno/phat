package cmd

import (
    "os"
    "os/exec"
)

func ClearScreen() {
    localCmd := exec.Command("clear")
    localCmd.Stdout = os.Stdout
    localCmd.Run()
}
