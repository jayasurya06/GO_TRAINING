package grading

func GetGrade(marks int, firstAttempt bool) (string, string) {
	var grade string

	switch {
	case marks >= 90:
		grade = "A"
	case marks >= 70:
		grade = "B"
	case marks >= 50:
		grade = "C"
	default:
		return "Fail", "fail"
	}

	status := "passed"
	if firstAttempt {
		status = "pass"
	}

	return grade, status
}

// Max returns the maximum of two numbers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
