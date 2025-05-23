// clientapi.go
package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Id       string   `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
}

type Stats struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func RunAPIQuiz() {		// Fetch quiz questions from the API
	reader := bufio.NewReader(os.Stdin)
	
	time.Sleep(1 * time.Second)
	
	fmt.Print("Enter your name: ")
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)
	
	resp, err := http.Get("http://localhost:8080/questions")
	if err != nil {
		log.Fatal("Could not fetch questions:", err)
	}
	defer resp.Body.Close()
	
	var questions []Question
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		log.Fatal("Could not decode questions:", err)
	}
	
	score := 0
	
	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for idx, option := range q.Options {
			fmt.Printf(" %d) %s\n", idx+1, option)
		}
		
		for {
			fmt.Print("Your answer (enter number 1-4): ")
			answerInput, _ := reader.ReadString('\n')
			answerInput = strings.TrimSpace(answerInput)
			answerNum, err := strconv.Atoi(answerInput)
			
			if err != nil || answerNum < 1 || answerNum > len(q.Options) {
				fmt.Println("Invalid input, please enter a number between 1-4.")
				continue
			}
			
			if answerNum-1 == q.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Printf("Wrong! Correct answer: %s\n", q.Options[q.Answer])
			}
			time.Sleep(2 * time.Second)
			break
		}
	}
	payload := map[string]interface{}{		// Sends final score to API
		"name":  name,
		"score": score,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Failed to marshal result:", err)
	}
	resp, err = http.Post("http://localhost:8080/submit", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Failed to send score:", err)
	}
	defer resp.Body.Close()
	
	var result struct {
		Score      int `json:"score"`
		Percentile int `json:"percentile"`
		Total      int `json:"total"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal("Failed to decode result:", err)
	}
	
	fmt.Printf("\nQuiz finished!\n")
	fmt.Printf("Your score: %d/%d\n", result.Score, result.Total)
	fmt.Printf("You scored better than %d%% of all players!\n", result.Percentile)
	time.Sleep(2 * time.Second)
}

func ShowHighscores() {	    // Print all players score in order
	time.Sleep(1 * time.Second)
	
	resp, err := http.Get("http://localhost:8080/stats")
	if err != nil {
		fmt.Println("Could not fetch stats:", err)
		return
	}
	defer resp.Body.Close()
	
	var stats []Stats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		fmt.Println("Could not decode stats:", err)
		return
	}
	
	fmt.Println("\nAll Players Highscore:")
	fmt.Println("------------------------")
	
	if len(stats) == 0 {
		fmt.Println("No games played yet!")
		return
	}
	
	for i, stat := range stats {
		fmt.Printf("%d. %s - %d correct answers\n", i+1, stat.Name, stat.Points)
	}
	fmt.Println()
}
