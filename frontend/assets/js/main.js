let correctAnswer = '';

document.getElementById("getRandomQuestion").addEventListener("click", function() {
    getRandomQuestion();
});

document.getElementById("submitAnswer").addEventListener("click", function() {
    checkAnswer();
});

async function getRandomQuestion() {
    try {
        const response = await fetch('/question/random');
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        const questionTextbox = document.getElementById("question");
        questionTextbox.value = data.question;
        correctAnswer = data.answer;

        // Reset the answer box
        const answerTextbox = document.getElementById("answer");
        answerTextbox.value = '';
        answerTextbox.style.backgroundColor = ''; // Reset the background color
    } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
    }
}

function checkAnswer() {
    const answerTextbox = document.getElementById("answer");
    const userAnswer = answerTextbox.value;

    if (userAnswer === correctAnswer) {
        answerTextbox.style.backgroundColor = 'limegreen';
    } else {
        answerTextbox.style.backgroundColor = 'red';
        answerTextbox.value = `${userAnswer} -> A helyes válasz: ${correctAnswer}`;
    }
}

getRandomQuestion(); // Call the function on page load
