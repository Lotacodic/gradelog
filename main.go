package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Please enter an argument")
	}

	courseDetails := os.Args[1:]
	totalCreditUnit := 0
	totalQualityPoints := 0

	for _, str := range courseDetails {
		parts := strings.Fields(str)

		if len(parts) < 3 {
			fmt.Printf("Error: Invalid format for %q. Expected 3 parts.\n", str)
			return
		}
		creditUnit, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error: Invalid unit %q\n", parts[1])
			return
		}

		score, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Printf("Error: Invalid score%q\n", parts[2])
			return
		}

		if score < 0 || score > 100 {
			fmt.Printf("Error: Score %d in %q is out of range (0-100)!\n", score, parts[0])
			return
		}

		checkScore := CheckGrade(score)
		point := CheckGradePoint(checkScore)
		qualityPoint := point * creditUnit

		fmt.Printf("%d -> %s -> %d -> %d\n", score, checkScore, point, qualityPoint)

		totalCreditUnit += creditUnit
		totalQualityPoints += qualityPoint
	}

	myCgpa := float64(totalQualityPoints) / float64(totalCreditUnit)
	myCgpaRange := CgpaRange(myCgpa)
	fmt.Printf("myCgpa: %.2f -> %s\n", myCgpa, myCgpaRange)
}
