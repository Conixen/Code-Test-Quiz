package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PlayGeoQuiz(questions []Questions) int{
	scan := bufio.NewScanner(os.Stdin) 	// Reader of terminals input

	playerScore := 0

	fmt.Println("Hello & Welcome to my Geography Quiz! :)")
	fmt.Println("---------------------------------------")

	for {
		fmt.Println(`		
		1. Play Quiz
		2. Check Score
		3. Quit`)		// Raw string literal, no \n
		fmt.Println("Pick a number between 1-3:")

		scan.Scan()
		userInput := strings.TrimSpace(scan.Text())
		
		switch userInput {
		case "1":
			fmt.Printf("Starting quiz...\n")
			playerScore = runQuiz(questions) 
		case "2":
			fmt.Printf("Your score: %d correct answers\n", playerScore)
		case "3":
			fmt.Println(" Goodbye!")
			return playerScore 
		default:
			fmt.Println("Please select a valid option (1, 2 or 3).")
		}
	}
}
func runQuiz(questions []Questions) int {
	scanner := bufio.NewScanner(os.Stdin)
	score := 0

	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for idx, alt := range q.Alternatvs {
			fmt.Printf("  %d. %s\n", idx+1, alt)
		}

		fmt.Print("Your answer (1–4): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		answerIndex := int(input[0] - '1')

		if answerIndex == q.RightAnswear {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Incorrect! Correct answer: %s\n", q.Alternatvs[q.RightAnswear])
		}
	}
	return score
}
