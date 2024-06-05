package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"lootlocker-quiz/internal/api"
	"lootlocker-quiz/internal/db"
	"lootlocker-quiz/internal/tools"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("======= LootLocker Quiz API server =======")

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db.InitDB()

	// Check if the command is to generate quiz content
	if len(os.Args) > 1 && os.Args[1] == "generate" {
		if len(os.Args) < 3 {
			log.Fatalf("You must provide the subject for the quiz. For example, 'go run cmd/server/main.go generate Pop Music'.")
		}
		subject := strings.Join(os.Args[2:], " ")
		log.Printf("Generating quiz with subject: %s ...", subject)
		runContentGeneration(subject)
		return
	}

	httpPort := os.Getenv("HTTP_PORT")
	// httpsPort := os.Getenv("HTTPS_PORT")
	router := api.SetupRouter()

	// Start HTTP server in a Go routine so that it runs concurrently
	// go func() {
	log.Println("Starting HTTP server on port:", httpPort)
	if err := http.ListenAndServe(":"+httpPort, router); err != nil {
		log.Fatalf("Failed to start HTTP server on port %s: %v", httpPort, err)
	}
	// }()

	// Start HTTPS server
	// log.Println("Starting HTTPS server on port:", httpsPort)
	// if err := http.ListenAndServeTLS(":"+httpsPort, "./web/ssl/fullchain.pem", "./web/ssl/privkey.pem", router); err != nil {
	// 	log.Fatalf("Failed to start HTTPS server on port %s: %v", httpsPort, err)
	// }
}

// redirectTLS redirects HTTP requests to HTTPS
// func redirectTLS(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
// }

// runContentGeneration generates quiz content on defined topic and saves it to the database
func runContentGeneration(subject string) {
	questions, err := tools.GenerateQuizQuestions(subject)
	if err != nil {
		log.Fatalf("Error generating questions: %v", err)
	}
	err = tools.SaveQuizToDatabase(db.DB, subject, questions)
	if err != nil {
		log.Fatalf("Error saving questions to database: %v", err)
	}
	fmt.Printf("Successfully generated a quiz on subject: %s", subject)
}
