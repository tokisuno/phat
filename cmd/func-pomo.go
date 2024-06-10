package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/gen2brain/beeep"
	"github.com/spf13/viper"
)


func startPomo() {
    // Viper configuration setup
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("$HOME/.config/phat")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %w", err))
    }

    work_duration, break_duration, topic := getTopic()

    num_pomos := viper.Get("session_amount").(int)

    fmt.Println(work_duration)
    for i := 1; i <= num_pomos; i++ {
        fmt.Println("Starting session...")
        sessionTimer(work_duration, "Session", i, num_pomos, topic)
        notifyUser()

        fmt.Println("Starting break")
        sessionTimer(break_duration, "Break", i, num_pomos, topic)
        notifyUser()
    }
}

func sessionTimer(duration int, state string, num int, denom int, topic string) {
    //duration *= 60
    startingTime := duration

    for duration >= 0 {
        ClearScreen()
        colouredOutput(duration, startingTime, state, num, denom, topic)
        time.Sleep(1 * time.Second)
        duration--
    }
    ClearScreen()
    fmt.Printf("%v over!", state)
}

func colouredOutput(duration int, startingTime int, state string, num int, denom int, topic string) {
    sessionText := color.New(color.FgWhite).Add(color.Underline)
    durationText := color.New(color.FgHiBlack)

    sessionText.Printf("%s Timer for %s: %v/%v\n", state, topic, num, denom)
     
    if duration <= startingTime/2 {
        durationText = color.New(color.FgHiWhite)
    } 

    if duration < 30 {
        durationText = color.New(color.FgHiRed)
    }

    if duration <= 5 {
        durationText = color.New(color.FgHiYellow)
    }

    durationText.Printf("✨ %.2d:%.2d ✨\n\r", duration/60, duration%60)
}

func notifyUser() {
    err := beeep.Notify("Time's up!", ":)", "assets/tomatoman.png")
    if err != nil {
        panic(err)
    }
}
