package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/gen2brain/beeep"
)

func startPomo() {
    // TODO: read these ints from config file
    work_duration := 10
    break_duration := 10

    num_pomos := 3

    fmt.Println(work_duration)
    for i := 1; i <= num_pomos; i++ {
        fmt.Println("Starting session...")
        startSession(work_duration, "Session", i, num_pomos)
        notifyUser()

        fmt.Println("Starting break")
        startSession(break_duration, "Break", i, num_pomos)
        notifyUser()
    }
}

func startSession(duration int, state string, num int, denom int) {
    //duration *= 60
    startingTime := duration

    for duration >= 0 {
        ClearScreen()
        colouredOutput(duration, startingTime, state, num, denom)
        time.Sleep(1 * time.Second)
        duration--
    }
    ClearScreen()
    fmt.Printf("%v over!", state)
}

func colouredOutput(duration int, startingTime int, state string, num int, denom int) {
    sessionText := color.New(color.FgWhite).Add(color.Underline)
    durationText := color.New(color.FgHiBlack)

    sessionText.Printf("%s Timer: %v/%v\n", state, num, denom)
     
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
