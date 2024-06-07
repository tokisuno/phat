package cmd

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

func startPomo() {
    // TODO: read these ints from config file
    work_duration := 60
    break_duration := 10

    num_pomos := 3

    fmt.Println(work_duration)
    for i := 0; i < num_pomos; i++ {
        fmt.Println("Starting session...")
        startSession(work_duration, "Session", i, num_pomos)
        notifyUser()

        fmt.Println("Starting break")
        startSession(break_duration, "Break", i, num_pomos)
        notifyUser()
    }
}

func startSession(duration int, state string, num int, denom int) {
    duration *= 60
    for duration >= 0 {
        ClearScreen()
        fmt.Printf("%s Timer:%v/%v >> %.2d:%.2d\r", state, num+1, denom, duration/60, duration%60)
        time.Sleep(1 * time.Second)
        duration--
    }
    fmt.Printf("%v over!", state)
}

func notifyUser() {
    err := beeep.Notify("Time's up!", ":)", "assets/tomatoman.png")
    if err != nil {
        panic(err)
    }
}
