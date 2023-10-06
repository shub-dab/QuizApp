package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fObj, _ := os.Open("quiz.csv")
	fmt.Println(fObj)

	csvR := csv.NewReader(fObj)
	fmt.Println(csvR)
	
	cLines, _ := csvR.ReadAll()

	fmt.Println(cLines)
}