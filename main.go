package main

import "fmt"

func main() {
	// Entry point for the grade calculator

	fmt.Println("Student Grade Calculator")

	subjectCount := getSubjectCount()
	scores := getScores(subjectCount)

	total := calculateTotal(scores)
	average := float64(total) / float64(subjectCount)
	grade := getGrade(average)
	status, remark := getRemark(average)

	fmt.Println("----- Result -----")
	fmt.Println("Total:", total)
	fmt.Println("Average:", average)
	fmt.Println("Grade:", grade)
	fmt.Println("Status:", status)
	fmt.Println("Remark:", remark)
}

// Gets number of subjects from user
func getSubjectCount() int {
	var count int
	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&count)
	return count
}

// Collects scores for all subjects
func getScores(count int) []int {
	scores := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter score for subject %d: ", i+1)
		fmt.Scanln(&scores[i])
	}

	return scores
}

// Calculates total score
func calculateTotal(scores []int) int {
	total := 0
	for _, score := range scores {
		total += score
	}
	return total
}

// Determines grade based on average
func getGrade(average float64) string {
	if average >= 70 {
		return "A"
	} else if average >= 60 {
		return "B"
	} else if average >= 50 {
		return "C"
	} else if average >= 45 {
		return "D"
	}
	return "F"
}

// Adds pass/fail status and remark
func getRemark(average float64) (string, string) {
	if average >= 50 {
		return "PASS", "Good performance"
	}
	return "FAIL", "Needs improvement"
}