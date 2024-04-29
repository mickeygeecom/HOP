package api

import (
	"database/sql"
	"encoding/json"
	"lootlocker-quiz/internal/db"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func userListQuizzes(w http.ResponseWriter, r *http.Request) {
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

func userGetQuiz(w http.ResponseWriter, r *http.Request) {
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

// SubmitAnswers handles the submission of quiz answers
func userSubmitAnswers(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to extract quiz answers
	var answers []struct {
		QuestionID int `json:"question_id"`
		Answer     int `json:"answer"`
	}
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Process and save the submitted answers (implementation details depend on your requirements)

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Quiz answers submitted successfully"))
}
