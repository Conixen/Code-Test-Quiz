package userstats

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type UserStats struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

var QuizStats []UserStats

const statsFile = "quiz_stats.json"

func SaveGame(name string, points int) {
	QuizStats = append(QuizStats, UserStats{
		Name:   name,
		Points: points,
	})
	SaveToFile()
}

func GetAllStats() []UserStats {	// show score in order
	sortedStats := make([]UserStats, len(QuizStats))
	copy(sortedStats, QuizStats)
	
	sort.Slice(sortedStats, func(i, j int) bool {
		return sortedStats[i].Points > sortedStats[j].Points
	})
	
	return sortedStats
}

func ShowStats() {
	fmt.Println("All Players Highscore:")
	fmt.Println("---------------------")
	fmt.Println()
	
	sortedStats := GetAllStats()
	
	for i, stat := range sortedStats {
		fmt.Printf("%d. %s - %d correct answers\n", i+1, stat.Name, stat.Points)
	}
	fmt.Println()
}

func GetPercentile(score int) int {
	
	betterThan := 0
	for _, s := range QuizStats {
		if score > s.Points {
			betterThan++
		}
	}
	
	percentile := (betterThan * 100) / (len(QuizStats) - 1) 
	return percentile
}

func SaveToFile() {		// save quiz run
	file, err := os.Create(statsFile)
	if err != nil {
		fmt.Println("Could not save stats:", err)
		return
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	err = encoder.Encode(QuizStats)
	if err != nil {
		fmt.Println("Failed to write stats to file:", err)
	}
}

func init() {
	file, err := os.Open(statsFile)
	if err != nil {
		return 
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&QuizStats)
	if err != nil {
		fmt.Println("Failed to load stats:", err)
	}
}