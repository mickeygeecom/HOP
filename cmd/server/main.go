package main

import (
	"log"
	"net/http"
	"os"

	"lootlocker-quiz/internal/api"
	"lootlocker-quiz/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("======= LootLocker Quiz API server =======")

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	httpPort := os.Getenv("HTTP_PORT")
	// httpsPort := os.Getenv("HTTPS_PORT")

	db.InitDB()
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
