import { displayQuizzes } from './displayQuizzes.js';

// Function to fetch quiz data from the backend API
export async function fetchQuizzes() {
    console.log('Fetching quizzes...');
    fetch('/api/user/quizzes')
        .then(response => response.json())
        .then(data => {
            displayQuizzes(data);
        })
        .catch(error => {
            console.error('Error fetching quizzes:', error);
            alert('Failed to fetch quizzes');
        });
}
