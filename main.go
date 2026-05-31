package main

import "fmt"

func main() {
	// Entry point for the grade calculator

	fmt.Println("Student Grade Calculator")

	var subjectCount int

	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&subjectCount)

	fmt.Println("You entered:", subjectCount)

	// Store all subject scores here
	scores := make([]int, subjectCount)

	// Collect scores one by one
	for i := 0; i < subjectCount; i++ {
		fmt.Printf("Enter score for subject %d: ", i+1)
		fmt.Scanln(&scores[i])
	}

	// Just printing what we collected for now
	fmt.Println("Scores entered:", scores)
}