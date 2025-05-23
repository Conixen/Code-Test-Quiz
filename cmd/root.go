package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/spf13/cobra"
	"geoquiz/api"
	"geoquiz/client"
)

var rootCmd = &cobra.Command{
	Use:   "geoquiz",
	Short: "A fun geography quiz game with API backend",
	Long: `Geography Quiz is a CLI application that lets you test your geography knowledge.
The quiz uses a REST API backend to serve questions and track scores.

Available commands:
- play: Start the geography quiz
- highscore: View player rankings  
- server: Start the API server manually
- menu: Interactive menu (like the original quiz.go)`,
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start the geography quiz directly",
	Long:  "Start playing the geography quiz directly. The API server will start automatically.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🌍 Welcome to the Geography Quiz!")
		fmt.Println("Starting quiz server...")
		
		// Starta API server i bakgrunden
		go api.StartServer()
		
		// Vänta så servern hinner starta
		time.Sleep(2 * time.Second)
		
		// Kör quiz via API
		client.RunAPIQuiz()
	},
}

var highscoreCmd = &cobra.Command{
	Use:   "highscore",
	Short: "View player rankings and statistics",
	Long:  "Display the highscore table showing all players and their scores.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📊 Loading highscores...")
		
		// Starta API server för att hämta stats
		go api.StartServer()
		
		// Vänta så servern hinner starta  
		time.Sleep(2 * time.Second)
		
		client.ShowHighscores()
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the API server manually",
	Long:  "Start the REST API server manually on port 8080. Useful for development.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting API server...")
		api.StartServer() // Detta blockerar
	},
}

// DENNA ÄR DIN URSPRUNGLIGA MENY FRÅN quiz.go!
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Interactive menu (original quiz.go style)",
	Long:  "Shows the original interactive menu with numbered options.",
	Run: func(cmd *cobra.Command, args []string) {
		showOriginalMenu()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Lägg till subkommandona
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(highscoreCmd) 
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(menuCmd)  // DIN URSPRUNGLIGA MENY!
	
	// Root command kör original meny som standard
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		showOriginalMenu()
	}
}

// DIN URSPRUNGLIGA MENY FRÅN quiz.go - EXAKT SAMMA!
func showOriginalMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Hello & Welcome to my Geography Quiz! :)")
	fmt.Println("---------------------------------------")
	
	// Starta API server EN GÅNG i bakgrunden
	serverStarted := false                    
	
	for {
		fmt.Println()
		fmt.Println("1. Play Quiz")
		fmt.Println("2. Check Score")
		fmt.Println("3. Quit")
		fmt.Println("\nPick a number between 1-3:")
		
		scanner.Scan()
		userInput := strings.TrimSpace(scanner.Text())
		
		switch userInput {
		case "1":
			if !serverStarted {               
				fmt.Println("Starting quiz server...")
				go api.StartServer()
				time.Sleep(2 * time.Second)
				serverStarted = true          
			}
			
			client.RunAPIQuiz()
			
		case "2":
			if !serverStarted {              
				fmt.Println("Starting server to fetch stats...")
				go api.StartServer()
				time.Sleep(2 * time.Second)
				serverStarted = true         
			}
			
			client.ShowHighscores()
			
		case "3":
			fmt.Println("Fare thy well thanks for playing!")
			return
			
		default:
			fmt.Println("Please select a valid option (1, 2 or 3).")
		}
	}
}
