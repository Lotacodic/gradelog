package main

import "math"

func CgpaRange(nb float64) string {
	if nb < 1.0 || nb > 5.0 || math.IsNaN(nb) {
		return "Invalid CGPA"
	}

	switch {
	case nb >= 4.5:
		return "First class honours"
	case nb >= 3.5:
		return "Second class honours(Upper Division)"
	case nb >= 2.4:
		return "Second class honours(Lower Division)"
	case nb >= 1.5:
		return "Third class honours"
	default:
		return "Pass"
	}
}
