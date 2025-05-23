// api/serverapi.go
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"geoquiz/quiz"
	"geoquiz/userstats"
)

type Answer struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func StartServer() {
	http.HandleFunc("/questions", handleQuestions)
	http.HandleFunc("/submit", handleSubmit)
	http.HandleFunc("/stats", handleStats)
	
	fmt.Println("API server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz.GeographyQuiz)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var ans Answer
	err := json.NewDecoder(r.Body).Decode(&ans)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	userstats.SaveGame(ans.Name, ans.Score)
	
	response := map[string]interface{}{
		"score":      ans.Score,
		"percentile": userstats.GetPercentile(ans.Score),
		"total":      len(quiz.GeographyQuiz),
	}
	
	json.NewEncoder(w).Encode(response)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userstats.GetAllStats())
}