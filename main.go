package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gradelog \"[Course] [Credits] [Score]\" ...")
		fmt.Println("Example: gradelog \"MTH101 4 75\" \"CHM101 3 62\"")
		os.Exit(1)
	}

	courseDetails := os.Args[1:]
	totalCreditUnit := 0
	totalQualityPoints := 0

	for _, str := range courseDetails {
		parts := strings.Fields(str)

		if len(parts) < 3 {
			fmt.Printf("Error: Invalid format for %q. Expected 3 parts (Name, Units, Score).\n", str)
			return
		}
		creditUnit, err := strconv.Atoi(parts[1])
		if err != nil || creditUnit < 0 {
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

		fmt.Printf("%s: Score %d -> Grade %s -> Points %d -> Quality Points %d\n", parts[0], score, checkScore, point, qualityPoint)

		totalCreditUnit += creditUnit
		totalQualityPoints += qualityPoint
	}

	if totalCreditUnit == 0 {
		fmt.Println("\nError: Total credit units cannot be zero. Cannnot compute CGPA.")
		return
	}

	myCgpa := float64(totalQualityPoints) / float64(totalCreditUnit)
	myCgpaRange := CgpaRange(myCgpa)

	fmt.Println(strings.Repeat("_", 40))
	fmt.Printf("Total Credits: %d | Total Quality Points: %d\n", totalCreditUnit, totalQualityPoints)
	fmt.Printf("Final CGPA: %.2f -> %s\n", myCgpa, myCgpaRange)
}
