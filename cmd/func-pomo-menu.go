package cmd

import (
	"fmt"
    "os"

	"github.com/spf13/viper"
	"github.com/tokisuno/gocliselect"
	"gopkg.in/yaml.v3"
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

    // scuffed rn, to be fixed
    f, err := os.ReadFile("/home/poto/.config/phat/config.yaml")
    if err != nil {
        panic("couldn't find config file")
    }

    var config Config

    err = yaml.Unmarshal(f, &config)
    if err != nil {
        panic("couldn't decode")
    }


    work_duration := viper.Get("session_length").(int)
    break_duration := viper.Get("break_length").(int)

    menu := gocliselect.NewMenu("What're you studying? (declare options in config.yaml)")

    for _, topic := range config.Topics {
        menu.AddItem(topic, topic)
    }

    option := menu.Display()
    
    return work_duration, break_duration, option
}

