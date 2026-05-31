package main

import "fmt"

func main() {
	// Entry point for the grade calculator

	fmt.Println("Student Grade Calculator")

	var subjectCount int

	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&subjectCount)

	scores := make([]int, subjectCount)

	for i := 0; i < subjectCount; i++ {
		fmt.Printf("Enter score for subject %d: ", i+1)
		fmt.Scanln(&scores[i])
	}

	// Calculate total score
	total := 0
	for _, score := range scores {
		total += score
	}

	// Calculate average
	average := float64(total) / float64(subjectCount)

	// Determine grade
	var grade string

	if average >= 70 {
		grade = "A"
	} else if average >= 60 {
		grade = "B"
	} else if average >= 50 {
		grade = "C"
	} else if average >= 45 {
		grade = "D"
	} else {
		grade = "F"
	}

	// Output results
	fmt.Println("----- Result -----")
	fmt.Println("Total:", total)
	fmt.Println("Average:", average)
	fmt.Println("Grade:", grade)
}