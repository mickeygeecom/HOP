package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// Function to check if a join code already exists in the database
func joinCodeExists(joinCode string) (bool, error) {
	// Prepare the query to check if the join code exists
	query := "SELECT COUNT(*) FROM quiz WHERE join_code = ?"

	// Execute the query and retrieve the count
	var count int
	err := DB.QueryRow(query, joinCode).Scan(&count)
	if err != nil {
		return false, err
	}

	// If the count is greater than 0, the join code exists
	return count > 0, nil
}

// GenerateJoinCode ... Function to generate a random join code with 6 characters, consisting of numbers and uppercase letters
func GenerateJoinCode() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

// UpdateSingleQuizJoinCode ... Function to update the join code for a single quiz
func UpdateSingleQuizJoinCode(quizID int) error {
	// Generate a unique join code for the current quiz
	newJoinCode := GenerateJoinCode()

	// Update the quiz record with its join code
	if _, err := DB.Exec("UPDATE quiz SET join_code = ? WHERE id = ?", newJoinCode, quizID); err != nil {
		return fmt.Errorf("error updating join code for quiz ID %d: %v", quizID, err)
	}

	log.Printf("Generated join code %s for quiz ID %d\n", newJoinCode, quizID)

	return nil
}

// UpdateQuizJoinCodes ... Function to generate and update join codes for quizzes
func UpdateQuizJoinCodes() error {
	// Query all quizzes from the database
	rows, err := DB.Query("SELECT id, join_code FROM quiz")
	if err != nil {
		return fmt.Errorf("error querying quizzes: %v", err)
	}
	defer rows.Close()

	// Loop through each quiz and update its join code if not already present
	for rows.Next() {
		var quizID int
		var joinCode sql.NullString
		if err := rows.Scan(&quizID, &joinCode); err != nil {
			return fmt.Errorf("error scanning quiz data: %v", err)
		}

		// If a join code is already present, skip generating a new one
		if joinCode.Valid && joinCode.String != "" {
			//log.Printf("Join code already exists for quiz ID %d\n", quizID)
			continue
		}

		// Generate a unique join code for the current quiz
		newJoinCode := GenerateJoinCode()

		// Update the quiz record with its join code
		if _, err := DB.Exec("UPDATE quiz SET join_code = ? WHERE id = ?", newJoinCode, quizID); err != nil {
			return fmt.Errorf("error updating join code for quiz ID %d: %v", quizID, err)
		}

		log.Printf("Generated join code %s for quiz ID %d\n", newJoinCode, quizID)
	}

	// Check for any errors during row iteration
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating over rows: %v", err)
	}

	return nil
}

// GetQuizIDFromJoinCode ... Function to retrieve the quiz ID from the join code
func GetQuizIDFromJoinCode(joinCode string) (string, error) {
	// Query the database to find the quiz ID based on the join code
	var quizID string
	err := DB.QueryRow("SELECT id FROM quiz WHERE join_code = ?", joinCode).Scan(&quizID)
	if err != nil {
		// If the join code doesn't exist, return a custom error
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("quiz not found")
		}
		return "", err
	}

	return quizID, nil
}
