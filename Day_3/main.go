package main

import (
	student "Day_3/Student" // import your Student package
	"fmt"
)

func main() {
	// slice of students from student package
	students := []student.Student{
		{Name: "Venky", Roll: 1, Marks: map[string]int{"Math": 70, "Science": 60, "English": 90}},
		{Name: "Surya", Roll: 2, Marks: map[string]int{"Math": 60, "Science": 75, "English": 55}},
		{Name: "Tarun", Roll: 3, Marks: map[string]int{"Math": 40, "Science": 55, "English": 48}},
	}

	report := make(map[int]student.Student)

	for _, s := range students {
		s.Calculate()      // call method from package
		report[s.Roll] = s // store in map
	}

	fmt.Printf("%-5s %-10s %-7s %-6s\n", "Roll", "Name", "Avg", "Grade")
	fmt.Println("----------------------------------")
	for _, s := range report {
		fmt.Printf("%-5d %-10s %-7f %-6s\n", s.Roll, s.Name, s.Avg, s.Grade)
	}
}
