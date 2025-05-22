package userstats

import(
	"fmt"
	"os"
	"encoding/json"
)

type UserStats struct{
	Name string
	Points int
}

var QuizStats []UserStats
const statsFile = "quiz_stats.json"

	func SaveGame(name string, points int){
	QuizStats = append (QuizStats, UserStats{
		Name: name,
		Points: points,
	})
	SaveToFile()
}

func ShowStats() {
	fmt.Println("All Players Highscore:")
	fmt.Println("---------------------")

	for i, stat := range QuizStats{
		fmt.Printf("%d. %s - %d correct answers\n",i+1, stat.Name, stat.Points)
	}
	fmt.Println()
}

func GetPercentile(score int) int{
	if len (QuizStats) == 0 {
		return 0
	}
	betterThan := 0
	for _, s := range QuizStats {
		if score > s.Points {
			betterThan++
		}
	}
	percentile := (betterThan * 100) / len(QuizStats)
	return percentile
}

func SaveToFile() {		//Method/function to save a json file for userstats
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


func init() {   	// Method/function to read json file when program starts
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
