package main

func CheckGrade(score int) string {
	if score > 100 || score < 0 {
		return "Score, out of range"
	}

	switch {
	case (score >= 70 && score <= 100):
		return "A"
	case (score >= 60 && score <= 69):
		return "B"
	case (score >= 50 && score <= 59):
		return "C"
	case (score >= 45 && score <= 49):
		return "D"
	case (score >= 40 && score <= 44):
		return "E"
	default:
		return "F"
	}
}
