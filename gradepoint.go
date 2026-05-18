package main

func CheckGradePoint(str string) int {
	switch {
	case str == "A":
		return 5
	case str == "B":
		return 4
	case str == "C":
		return 3
	case str == "D":
		return 2
	case str == "E":
		return 1
	default:
		return 0
	}
}
