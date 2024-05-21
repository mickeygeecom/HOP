// Initialize an array to store the answers
let answers = [];

// Function to display the score
export async function displayScore(quizId, quizQuestions, answers) {
    let score = 0;

    // Iterate through answers and increment score for each correct answer
    answers.forEach(answer => {
        // Find the corresponding question
        const question = quizQuestions.find(q => q.id === answer.question_id);
        if (question && answer.answer === question.correct_option) {
            score++;
        }
    });

    // Clear the quizForm element
    const quizForm = document.getElementById('quizForm');
    quizForm.innerHTML = '';

    // Display the score on the page
    const scoreElement = document.createElement('p');
    scoreElement.innerHTML = `Score percentage: <u>${Math.floor(score / answers.length * 100)}%</u><br>You answered <u>${score}</u> questions right, out of <u>${answers.length}</u> total.`;
    quizForm.appendChild(scoreElement);
}
