package cmd 

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_"github.com/spf13/viper"
)

func writeData() {
    // open the file
    f, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal(err)
    }

    w := csv.NewWriter(f)
    topic := "english"
    minutes := 40
    w.Write([]string{topic, fmt.Sprint(minutes)})
    w.Flush() 
}

// FOR LATER
// func readData() {} 
