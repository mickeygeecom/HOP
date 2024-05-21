import { fetchQuizzes } from './fetchQuizzes.js';

// Call fetchQuizzes function when the page loads
window.onload = fetchQuizzes;


// Username stuff
document.addEventListener("DOMContentLoaded", function() {
    // Get the form and input element
    var userForm = document.getElementById("userForm");
    var usernameInput = document.getElementById("usernameInput");
    var userGreeting = document.getElementById("userGreeting");

    // Check if a session exists
    if (sessionStorage.getItem("username")) {
        userGreeting.innerHTML = "Welcome back, <b>" + sessionStorage.getItem("username") + "</b>😍";
        alert("Welcome back " + sessionStorage.getItem("username") + "! 😍");
        userForm.style.display = "none";
    }

    // Add submit event listener to the form
    userForm.addEventListener("submit", function(event) {
        event.preventDefault(); // Prevent default form submission

        // Get the username entered by the user
        var username = usernameInput.value.trim();

        // Check if username is provided
        if (username !== "") {
            sessionStorage.setItem("username", username);
            alert("Welcome " + username + "! 😍 Your results will now be saved along with your name");
            userGreeting.innerHTML = "Welcome back, <b>" + username + "</b>😍";
            userForm.style.display = "none";
        }
    });
});


