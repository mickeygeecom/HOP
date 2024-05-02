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
	rows, err := db.DB.Query("SELECT id, name, description, join_code, created_at FROM quiz")
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	quizzes := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var name, description, join_code string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &description, &join_code, &createdAt); err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		quizzes = append(quizzes, map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
			"join_code":   join_code,
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
		JoinCode    string    `json:"join_code"`
		CreatedAt   time.Time `json:"created_at"`
	}

	row := db.DB.QueryRow("SELECT id, name, description, join_code, created_at FROM quiz WHERE id = ?", quizID)
	err := row.Scan(&quiz.ID, &quiz.Name, &quiz.Description, &quiz.JoinCode, &quiz.CreatedAt)
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
	// Parse quiz ID from the request URL
	params := mux.Vars(r)
	quizID := params["quizId"]

	// Parse the request body to extract quiz answers
	var answers []struct {
		QuestionID int `json:"question_id"`
		Answer     int `json:"answer"`
	}
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch correct answers for the quiz from the database
	var correctAnswers []struct {
		QuestionID    int `json:"question_id"`
		CorrectOption int `json:"correct_option"`
	}
	rows, err := db.DB.Query("SELECT id, correct_option FROM questions WHERE quiz_id = ?", quizID)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var questionID, correctOption int
		if err := rows.Scan(&questionID, &correctOption); err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		correctAnswers = append(correctAnswers, struct {
			QuestionID    int `json:"question_id"`
			CorrectOption int `json:"correct_option"`
		}{QuestionID: questionID, CorrectOption: correctOption})
	}

	// Compare user's answers with correct answers and calculate the score
	score := 0
	for _, userAnswer := range answers {
		for _, correctAnswer := range correctAnswers {
			if userAnswer.QuestionID == correctAnswer.QuestionID && userAnswer.Answer == correctAnswer.CorrectOption {
				// Increment the score for each correct answer
				score++
			}
		}
	}

	// Prepare response JSON indicating correctness of answers and the score
	response := struct {
		Correct bool `json:"correct"`
		Score   int  `json:"score"`
	}{Correct: score == len(correctAnswers), Score: score}

	// Encode response JSON and send it in the HTTP response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
