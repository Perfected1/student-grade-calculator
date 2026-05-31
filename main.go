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
	status, remark := getRemark(grade)

	fmt.Println("----- Result -----")
	fmt.Println("Total:", total)
	fmt.Println("Average:", average)
	fmt.Println("Grade:", grade)
	fmt.Println("Status:", status)
	fmt.Println("Remark:", remark)
}

// Gets number of subjects from user with validation
func getSubjectCount() int {
	var count int

	for {
		fmt.Print("Enter number of subjects: ")
		fmt.Scanln(&count)

		if count > 0 {
			return count
		}

		fmt.Println("Please enter a valid number greater than 0")
	}
}

// Collects validated scores for all subjects
func getScores(count int) []int {
	scores := make([]int, count)

	for i := 0; i < count; i++ {
		for {
			var score int

			fmt.Printf("Enter score for subject %d (0 - 100): ", i+1)
			fmt.Scanln(&score)

			if score >= 0 && score <= 100 {
				scores[i] = score
				break
			}

			fmt.Println("Invalid score. Enter a value between 0 and 100.")
		}
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
func getRemark(grade string) (string, string) {
	switch grade {
	case "A":
		return "PASS", "Excellent performance"
	case "B":
		return "PASS", "Very good performance"
	case "C":
		return "PASS", "Good performance"
	case "D":
		return "PASS", "Fair performance, needs improvement"
	default:
		return "FAIL", "Poor performance, repeat recommended"
	}
}