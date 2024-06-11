package cmd

import (
    "fmt"

    "github.com/tokisuno/gocliselect"
	"github.com/spf13/viper"
)

type Config struct {
    Topics []string `yaml:"topics"`
}

func getTopic() (int, int, string) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("$HOME/.config/phat")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %w", err))
    }

    work_duration := viper.Get("session_length").(int)
    break_duration := viper.Get("break_length").(int)

    menu := gocliselect.NewMenu("What're you studying? (declare options in config.yaml)")

    menu.AddItem("Programming", "programming")
    menu.AddItem("Japanese", "japanese")

    option := menu.Display()
    
    return work_duration, break_duration, option
}

