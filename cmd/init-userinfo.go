package cmd

import (
    "log"

    "github.com/spf13/cobra"
    "github.com/adrg/xdg"
)

func init() {
    rootCmd.AddCommand(configCmd) 
    rootCmd.AddCommand(dataCmd) 
}


var configCmd = &cobra.Command {
    Use: "whereconf",
    Short: "Finds config",
    Long: "Prints where on your system your config.yaml file is stored.",
    Run: func(cmd *cobra.Command, args []string) {
        saveConfig() 
    },
}


var dataCmd = &cobra.Command {
    Use: "wheredata",
    Short: "Finds data.csv",
    Long: "Prints where on your system your data.csv file is stored.",
    Run: func(cmd *cobra.Command, args []string) {
        saveData() 
    },
}

func saveConfig() {
    configFilePath, err := xdg.ConfigFile("phat/config.yaml")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Config file was found at: ", configFilePath)
}

func saveData() {
    saveDataPath, err := xdg.ConfigFile("phat/data.csv")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Saved user data at: ", saveDataPath)
}


