package userstats

import(
	"fmt"
)

type UserStats struct{
	Name string
	Points int
}

var QuizStats []UserStats
	func SaveGame(name string, points int){
	QuizStats = append (QuizStats, UserStats{
		Name: name,
		Points: points,
	})
}

func ShowStats() {
	fmt.Println("All Players Highscore:")
	fmt.Println("---------------------")

	for i, stat := range QuizStats{
		fmt.Printf("%d. %s - %d correct answers\n",i+1, stat.Name, stat.Points)
	}
	fmt.Println()
}


