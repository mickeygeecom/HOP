package api

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

// SetupRouter initializes and returns a new HTTP router for the API.
// It configures routes and attaches the necessary handlers for managing quizzes.
func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Apply CORS to the main router to handle all incoming requests, needs to be on top of the sub-router
	router.Use(enableCORS)

	// Specific sub-router for API paths
	apiRouter := router.PathPrefix("/").Subrouter()
	apiRouter.Use(ContentTypeMiddleware)

	// ## Admin endpoints ##
	adminRouter := router.PathPrefix("/api/admin").Subrouter()
	adminRouter.Use(ContentTypeMiddleware)
	// Admin - Quizzes
	adminRouter.HandleFunc("/quizzes", adminListQuizzes).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/quizzes", createQuiz).Methods("POST")
	adminRouter.HandleFunc("/quizzes/{id}", getQuiz).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/quizzes/{id}", updateQuiz).Methods("PUT", "OPTIONS")
	adminRouter.HandleFunc("/quizzes/{id}", deleteQuiz).Methods("DELETE", "OPTIONS")
	// Admin - Questions
	adminRouter.HandleFunc("/quizzes/{quizId}/questions", CreateQuestion).Methods("POST")
	adminRouter.HandleFunc("/questions/{questionId}", GetQuestionByID).Methods("GET")
	adminRouter.HandleFunc("/quizzes/{quizId}/questions", ListQuestions).Methods("GET")
	adminRouter.HandleFunc("/questions/{questionId}", UpdateQuestion).Methods("PUT")
	adminRouter.HandleFunc("/questions/{questionId}", DeleteQuestion).Methods("DELETE")

	// ## Enduser endpoints
	userRouter := router.PathPrefix("/api/user").Subrouter()
	userRouter.Use(ContentTypeMiddleware)
	userRouter.HandleFunc("/quizzes", userListQuizzes).Methods("GET", "OPTIONS")
	userRouter.HandleFunc("/quizzes/{id}", userGetQuiz).Methods("GET", "OPTIONS")
	userRouter.HandleFunc("/quizzes/{quizId}/submit", userSubmitAnswers).Methods("POST", "OPTIONS")
	userRouter.HandleFunc("/questions/{quizId}/", ListQuestions).Methods("GET")

	// Serve static files
	staticDir := "/web/" // Directory where static files are located
	router.PathPrefix("/").Handler(ServeStatic(filepath.Join(".", staticDir)))

	return router
}

// ContentTypeMiddleware is a middleware that sets the Content-Type header to application/json.
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// ServeStatic is a custom middleware function that sets the appropriate MIME type for CSS files.
func ServeStatic(dir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve CSS files with the correct MIME type
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}

		// Serve the file from the specified directory
		http.ServeFile(w, r, filepath.Join(dir, r.URL.Path))
	})
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Pre-flight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
