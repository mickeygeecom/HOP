// Function to fetch quiz data from the backend API
function fetchQuizzes() {
    fetch('/api/user/quizzes')
        .then(response => response.json())
        .then(data => {
            // Call function to display quiz items
            displayQuizzes(data);
        })
        .catch(error => {
            console.error('Error fetching quizzes:', error);
            alert('Failed to fetch quizzes');
        });
}

// Function to display quiz items in the HTML document
function displayQuizzes(quizzes) {
    const quizList = document.getElementById('quizList');

    // Clear existing quiz items
    quizList.innerHTML = '';

    // Loop through each quiz and create HTML elements to represent them
    quizzes.forEach(quiz => {
        const quizItem = document.createElement('div');
        quizItem.classList.add('quiz-item');

        const title = document.createElement('h2');
        title.textContent = quiz.title;

        const description = document.createElement('p');
        description.textContent = quiz.description;

        // You can customize the appearance of each quiz item here
        // For example, you can add additional elements like buttons or images

        // Append title and description to quizItem
        quizItem.appendChild(title);
        quizItem.appendChild(description);

        // Append quizItem to quizList
        quizList.appendChild(quizItem);
    });
}

// Call fetchQuizzes function when the page loads
window.onload = fetchQuizzes;