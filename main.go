package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Please enter an argument")
	}

	score := os.Args[1]

	myScore, err := strconv.Atoi(score)
	if err != nil {
		fmt.Println("Invalid score entered")
	}

	myGrade := CheckGrade(myScore)
	fmt.Printf("%d -> %s\n", myScore, myGrade)
}
