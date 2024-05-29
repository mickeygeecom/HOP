import { showQuizDetails } from './quizDetails.js';

// Function to handle joining the quiz
export async function joinQuiz() {
    // Get the value entered by the user in the input field
    const joinCode = document.getElementById('joinCodeInput').value;

    // Make a request to the backend API to resolve the join code to a quiz ID
    try {
        const response = await fetch(`/code/${joinCode}`);
        if (!response.ok) {
            throw new Error('Failed to join quiz');
        }
        const data = await response.json();
        const quizId = data.quiz_id;

        // Call the showQuizDetails function with the retrieved quiz ID
        await showQuizDetails(quizId);

        alert('Quiz found and joined', 'success')
    } catch (error) {
        console.error('Error joining quiz:', error.message);
        alert('Failed to join quiz. Did you type the right Join Code?', 'error');
    }
}

// Add event listener for the button click
document.getElementById('joinQuizButton').addEventListener('click', joinQuiz);

// Add event listener for the ENTER key press
document.getElementById('joinCodeInput').addEventListener('keypress', function(event) {
    if (event.key === 'Enter') {
        event.preventDefault();
        joinQuiz();
    }
});
