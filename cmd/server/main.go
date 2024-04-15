package main

import (
	"log"
	"net/http"

	"lootlocker-quiz/internal/api"
	"lootlocker-quiz/internal/db"
)

var httpPort = "80" // Use non-privileged ports for development
var httpsPort = "443"

func main() {
	db.InitDB()
	router := api.SetupRouter()

	log.Println("======= LootLocker Quiz API server =======")

	// Start HTTP server in a Go routine so that it runs concurrently
	go func() {
		log.Println("Starting HTTP server on port:", httpPort)
		if err := http.ListenAndServe(":"+httpPort, http.HandlerFunc(redirectTLS)); err != nil {
			log.Fatalf("Failed to start HTTP server on port %s: %v", httpPort, err)
		}
	}()

	// Start HTTPS server
	log.Println("Starting HTTPS server on port:", httpsPort)
	if err := http.ListenAndServeTLS(":"+httpsPort, "./web/ssl/cert.pem", "./web/ssl/key.pem", router); err != nil {
		log.Fatalf("Failed to start HTTPS server on port %s: %v", httpsPort, err)
	}
}

// redirectTLS redirects HTTP requests to HTTPS
func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
}
