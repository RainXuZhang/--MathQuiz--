//Created by Rain

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Operation int

const (
	Addition Operation = iota
	Subtraction
	Multiplication
	Division
)

type QuizStats struct {
	TotalQuestions int
	CorrectAnswers int
	CurrentStreak  int
	MaxStreak      int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the Math Quiz!")

	numQuestions := getNumberOfQuestions()
	minRange, maxRange := getNumberRanges()
	operation := selectOperation()

	var stats QuizStats
	stats.TotalQuestions = numQuestions

	startTime := time.Now()
	runQuiz(numQuestions, minRange, maxRange, operation, &stats)
	displayResults(&stats, startTime)
}

func selectOperation() Operation {
	fmt.Println("\nSelect Operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	var choice int
	fmt.Print("Enter your choice (1-4): ")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil || choice < 1 || choice > 4 {
		fmt.Println("Invalid choice. Defaulting to Addition.")
		return Addition
	}
	return Operation(choice - 1)
}

func runQuiz(numQuestions, minRange, maxRange int, op Operation, stats *QuizStats) {
	for i := 0; i < numQuestions; i++ {
		num1, num2, answer := generateQuestion(op, minRange, maxRange)
		if askQuestion(num1, num2, answer, op) {
			stats.CurrentStreak++
			if stats.CurrentStreak > stats.MaxStreak {
				stats.MaxStreak = stats.CurrentStreak
			}
			stats.CorrectAnswers++
		} else {
			stats.CurrentStreak = 0
		}
	}
}

func generateQuestion(op Operation, minRange, maxRange int) (int, int, int) {
	num1 := minRange + rand.Intn(maxRange-minRange+1)
	num2 := minRange + rand.Intn(maxRange-minRange+1)

	switch op {
	case Addition:
		return num1, num2, num1 + num2
	case Subtraction:
		if num1 < num2 {
			num1, num2 = num2, num1 // Ensure positive result
		}
		return num1, num2, num1 - num2
	case Multiplication:
		return num1, num2, num1 * num2
	case Division:
		// Generate division problems with whole number answers
		answer := minRange + rand.Intn(maxRange-minRange+1)
		num2 = minRange + rand.Intn(maxRange-minRange+1)
		num1 = answer * num2
		return num1, num2, answer
	default:
		return num1, num2, num1 + num2
	}
}

func askQuestion(num1, num2, answer int, op Operation) bool {
	var symbol string
	switch op {
	case Addition:
		symbol = "+"
	case Subtraction:
		symbol = "-"
	case Multiplication:
		symbol = "*"
	case Division:
		symbol = "/"
	}

	fmt.Printf("What is %d %s %d? ", num1, symbol, num2)

	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return false
	}

	if input == answer {
		fmt.Println("Correct!")
		return true
	}

	fmt.Printf("Incorrect. The correct answer is %d\n", answer)
	return false
}

func getNumberOfQuestions() int {
	fmt.Print("Enter the number of questions (default 5, max 50): ")

	var numQuestions int
	_, err := fmt.Scanf("%d", &numQuestions)
	if err != nil || numQuestions <= 0 {
		fmt.Println("Invalid input. Using default of 5 questions.")
		return 5
	}
	if numQuestions > 50 {
		fmt.Println("Maximum of 50 questions allowed. Using 50.")
		return 50
	}
	return numQuestions
}

func getNumberRanges() (int, int) {
	var minRange, maxRange int
	fmt.Print("Enter the minimum range (default 1): ")
	_, err := fmt.Scanf("%d", &minRange)
	if err != nil || minRange <= 0 {
		minRange = 1
	}
	fmt.Print("Enter the maximum range (default 10): ")
	_, err = fmt.Scanf("%d", &maxRange)
	if err != nil || maxRange < minRange {
		maxRange = 10
	}
	return minRange, maxRange
}

func displayResults(stats *QuizStats, startTime time.Time) {
	elapsedTime := time.Since(startTime)
	fmt.Println("\n--- Quiz Results ---")
	fmt.Printf("Total Questions: %d\n", stats.TotalQuestions)
	fmt.Printf("Correct Answers: %d\n", stats.CorrectAnswers)

	accuracy := float64(stats.CorrectAnswers) * 100 / float64(stats.TotalQuestions)
	fmt.Printf("Accuracy: %.2f%%\n", accuracy)

	fmt.Printf("Max Streak: %d\n", stats.MaxStreak)
	fmt.Printf("Total Time: %.2f seconds\n", elapsedTime.Seconds())
}
