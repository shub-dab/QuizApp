package main

import (
	"flag"
	"fmt"
)

func main() {
    fName := flag.String("f", "quiz.csv", "path of the csv file")

	timer := flag.Int("t", 30, "timer for the quiz")

    fmt.Println(fName, *fName)
    fmt.Println(timer, *timer)
}