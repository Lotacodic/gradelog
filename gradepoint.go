package main

func CheckGradePoint(str string) int {
	switch str {
	case "A":
		return 5
	case "B":
		return 4
	case "C":
		return 3
	case "D":
		return 2
	case "E":
		return 1
	default:
		return 0
	}
}
