import { showQuizDetails } from './quizDetails.js';

// Function to display quiz items in the HTML document
export async function displayQuizzes(quizzes) {
    const quizList = document.getElementById('quizList');

    // Clear existing quiz items
    quizList.innerHTML = '';

    // Loop through each quiz and create HTML elements to represent them
    quizzes.forEach(quiz => {
        const quizItem = document.createElement('div');
        quizItem.classList.add('quiz-item');

        const title = document.createElement('h2');
        title.textContent = quiz.name; // Display quiz name instead of title
        title.style.fontWeight = 'bold'; // Set the font weight of the title

        const description = document.createElement('p');
        description.textContent = quiz.description;

        // Automatically start the quiz when the user clicks on the quiz name
        quizItem.addEventListener('click', async () => {
            // Check if the quiz has a join code
            if (quiz.join_code) {
                //const newUrl = `/join/${quiz.join_code}`;
                //window.history.pushState({ path: newUrl }, '', newUrl);
                await showQuizDetails(quiz.id);
            }
        });

        // Append elements to quizItem
        quizItem.appendChild(title);
        quizItem.appendChild(description);

        // Append quizItem to quizList
        quizList.appendChild(quizItem);
    });
}
