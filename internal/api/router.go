package api

import (
	"database/sql"
	"encoding/json"
	"lootlocker-quiz/internal/db"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(ContentTypeMiddleware)
	// Admin - Quizzes
	adminRouter.HandleFunc("/quizzes", listQuizzes).Methods("GET", "OPTIONS")
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
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(ContentTypeMiddleware)

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

// ################ 		QUIZZES 		################

func createQuiz(w http.ResponseWriter, r *http.Request) {
	var q struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Invalid request body - JSON expected", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("INSERT INTO quiz (name, description) VALUES (?, ?)", q.Name, q.Description)
	if err != nil {
		http.Error(w, "Server error - database error", http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	quizID := int(id) // for api message, int64 to int
	if err != nil {
		http.Error(w, "Server error - database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"id": quizID, "status": "success", "message": "quiz id " + strconv.Itoa(quizID) + " created"})
}

func listQuizzes(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, description, created_at FROM quiz")
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	quizzes := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var name, description string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &description, &createdAt); err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		quizzes = append(quizzes, map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
			"created_at":  createdAt.Format(time.RFC3339),
		})
	}

	json.NewEncoder(w).Encode(quizzes)
}

func getQuiz(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID := params["id"]

	var quiz struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
	}

	row := db.DB.QueryRow("SELECT id, name, description, created_at FROM quiz WHERE id = ?", quizID)
	err := row.Scan(&quiz.ID, &quiz.Name, &quiz.Description, &quiz.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Quiz not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(quiz)
}

func updateQuiz(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID := params["id"]

	var q struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE quiz SET name = ?, description = ? WHERE id = ?", q.Name, q.Description, quizID)
	if err != nil {
		http.Error(w, "Server error - unable to update quiz, did you provide the correct id?", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": quizID, "status": "success", "message": "quiz id " + quizID + " updated"})
}

func deleteQuiz(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID := params["id"]

	// Check if the quiz exists before attempting to delete it
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM quiz WHERE id = ?", quizID).Scan(&count)
	if err != nil {
		http.Error(w, "Server error - unable to check quiz existence", http.StatusInternalServerError)
		return
	}

	if count == 0 {
		// Quiz with the provided ID does not exist
		http.Error(w, "Quiz not found", http.StatusNotFound)
		return
	}

	// Delete the quiz
	_, err = db.DB.Exec("DELETE FROM quiz WHERE id = ?", quizID)
	if err != nil {
		http.Error(w, "Server error - unable to delete quiz", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": quizID, "status": "success", "message": "Quiz deleted successfully"})
}

// ################ 		QUESTIONS 		################

// CreateQuestion creates a new question and associates it with the specified quiz
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID := params["quizId"]

	var q struct {
		QuestionText  string `json:"question_text"`
		Option1       string `json:"option1"`
		Option2       string `json:"option2"`
		Option3       string `json:"option3"`
		Option4       string `json:"option4"`
		CorrectOption int    `json:"correct_option"`
	}

	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_option) VALUES (?, ?, ?, ?, ?, ?, ?)", quizID, q.QuestionText, q.Option1, q.Option2, q.Option3, q.Option4, q.CorrectOption)
	if err != nil {
		http.Error(w, "Server error - couldn't insert question", http.StatusInternalServerError)
		return
	}
	questionID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"question_id": questionID})
}

// ListQuestions lists all questions associated with the specified quiz
func ListQuestions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID := params["quizId"]

	rows, err := db.DB.Query("SELECT id, question_text FROM questions WHERE quiz_id = ?", quizID)
	if err != nil {
		http.Error(w, "Server error - unable to find questions for quiz ID "+quizID, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var questions []map[string]interface{}
	for rows.Next() {
		var id int
		var questionText string
		err := rows.Scan(&id, &questionText)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		questions = append(questions, map[string]interface{}{"id": id, "question_text": questionText})
	}

	json.NewEncoder(w).Encode(questions)
}

// GetQuestionByID retrieves the details of the specified question
func GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	questionID := params["questionId"]

	// Query the database to get the question with the specified ID
	var q struct {
		QuestionText  string `json:"question_text"`
		Option1       string `json:"option1"`
		Option2       string `json:"option2"`
		Option3       string `json:"option3"`
		Option4       string `json:"option4"`
		CorrectOption int    `json:"correct_option"`
	}

	err := db.DB.QueryRow("SELECT question_text, option1, option2, option3, option4, correct_option FROM questions WHERE id = ?", questionID).Scan(&q.QuestionText, &q.Option1, &q.Option2, &q.Option3, &q.Option4, &q.CorrectOption)
	if err != nil {
		http.Error(w, "Failed to get question data - Did you provide the correct ID?", http.StatusInternalServerError)
		return
	}

	// Convert the question data to JSON and write it to the response
	json.NewEncoder(w).Encode(q)
}

// UpdateQuestion updates the details of the specified question
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	questionID := params["questionId"]

	var q struct {
		QuestionText string `json:"question_text"`
		Option1      string `json:"option1"`
		Option2      string `json:"option2"`
		Option3      string `json:"option3"`
		Option4      string `json:"option4"`
	}
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE questions SET question_text = ?, option1 = ?, option2 = ?, option3 = ?, option4 = ? WHERE id = ?", q.QuestionText, q.Option1, q.Option2, q.Option3, q.Option4, questionID)
	if err != nil {
		http.Error(w, "Server error - unable to update question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Question updated successfully"})
}

// DeleteQuestion deletes the specified question
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	questionID := params["questionId"]

	_, err := db.DB.Exec("DELETE FROM questions WHERE id = ?", questionID)
	if err != nil {
		http.Error(w, "Server error - unable to delete question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Question deleted successfully"})
}
