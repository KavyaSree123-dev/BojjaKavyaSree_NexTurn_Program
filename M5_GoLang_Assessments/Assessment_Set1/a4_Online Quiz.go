package main

import (
	"fmt"
	"strconv"
	"time"
)

// Struct to represent a question
type Question struct {
	Question string
	Options  [3]string
	Ans      int
}

func main() {
	// Question Bank
	questions := []Question{
		{
			Question: "What has keys but can't open locks?",
			Options:  [3]string{"1. A piano", "2. A map", "3. A clock"},
			Ans:      1,
		},
		{
			Question: "I speak without a mouth and hear without ears. What am I?",
			Options:  [3]string{"1. An echo", "2. A shadow", "3. A wind"},
			Ans:      1,
		},
		{
			Question: "What has to be broken before you can use it?",
			Options:  [3]string{"1. An egg", "2. A seal", "3. A nut"},
			Ans:      1,
		},
	}

	var score int
	var userInput string
	timer := 15 * time.Second // Time limit for each question

	// Take the quiz
	for i, q := range questions {
		// Display question and options
		fmt.Println("Question", i+1, ":", q.Question)
		for _, option := range q.Options {
			fmt.Println(option)
		}

		// Timer starts
		timerExpired := make(chan bool)
		go func() {
			time.Sleep(timer)
			timerExpired <- true
		}()

		// User input with error handling
		inputDone := make(chan bool)
		go func() {
			fmt.Print("Enter your answer (1-3) or 'exit' to quit: ")
			fmt.Scanln(&userInput)

			// Handle exit command
			if userInput == "exit" {
				fmt.Println("You exited the quiz.")
				close(inputDone)
				return
			}

			userAnswer, err := strconv.Atoi(userInput)
			if err != nil || userAnswer < 1 || userAnswer > 3 {
				fmt.Println("Invalid input, please enter a number between 1 and 3.")
			} else {
				// Valid input, process the answer
				if userAnswer == q.Ans {
					score++
				}
				close(inputDone)
			}
		}()

		// Check if timer expired or user input is done
		select {
		case <-timerExpired:
			fmt.Println("Time's up for this question!")
			close(inputDone)
		case <-inputDone:
		}
	}

	// Score Calculation
	fmt.Println("\nYour final score is:", score, "/", len(questions))

	// Performance classification
	switch {
	case score == len(questions):
		fmt.Println("Excellent")
	case score >= len(questions)/2:
		fmt.Println("Good")
	default:
		fmt.Println("Needs Improvement")
	}
}
