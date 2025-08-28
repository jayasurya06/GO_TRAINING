package main

import (
	grading "Day_2/Grading"
	"fmt"
)

func main() {
	// Example marks for 10 students
	marks := [10]int{95, 82, 67, 45, 30, 76, 54, 92, 49, 88}

	totalPass, totalFail := 0, 0

	for i, m := range marks {
		grade, status := grading.GetGrade(m, true)

		// Handle re-exam if failed
		if status == "fail" {
			var reMark int
			fmt.Printf("Enter re-exam marks for student %d (original %d): ", i+1, m)
			fmt.Scan(&reMark)

			finalMarks := grading.Max(m, reMark)
			grade, status = grading.GetGrade(finalMarks, false)

			fmt.Printf("Student: %d | Final Marks: %d | Grade: %s | Status: %s\n",
				i+1, finalMarks, grade, status)
		} else {
			fmt.Printf("Student: %d | Marks: %d | Grade: %s | Status: %s\n",
				i+1, m, grade, status)
		}

		// Update summary counters
		if status == "fail" {
			totalFail++
		} else {
			totalPass++
		}
	}

	// Final summary
	fmt.Println("\nSummary:")
	fmt.Println("Total Passed:", totalPass)
	fmt.Println("Total Failed:", totalFail)
}
