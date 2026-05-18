package main

func CgpaRange(nb float64) string {
	if nb < 1 || nb > 5 {
		return "Invalid CGPA"
	}

	switch {
	case nb >= 4.5 && nb <= 5.00:
		return "First class honours"
	case nb >= 3.50 && nb <= 4.49:
		return "Second class honours(Upper Division)"
	case nb >= 2.40 && nb <= 3.49:
		return "Second class honours(Lower Division)"
	case nb >= 1.50 && nb <= 2.39:
		return "Third class honours"
	default:
		return "Pass"
	}
}
