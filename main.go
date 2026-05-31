package main

import "fmt"

func main() {
	// Entry point for the grade calculator

	fmt.Println("Student Grade Calculator")

	var subjectCount int

	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&subjectCount)

	fmt.Println("You entered:", subjectCount)
}