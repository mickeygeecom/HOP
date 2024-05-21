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
        const truncatedTitle = quiz.name.length > 50 ? quiz.name.substring(0, 50) + '...' : quiz.name
        title.textContent = truncatedTitle
        title.style.fontWeight = 'bold'

        const description = document.createElement('p')
        const truncatedDescription = quiz.description.length > 70 ? quiz.description.substring(0, 70) + '...' : quiz.description
        description.textContent = truncatedDescription


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
