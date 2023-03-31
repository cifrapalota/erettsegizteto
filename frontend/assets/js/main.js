let correctAnswer = '';

document.getElementById("getRandomQuestion").addEventListener("click", function() {
    getRandomQuestion();
});

document.getElementById("submitAnswer").addEventListener("click", function() {
    checkAnswer();
});

document.getElementById("majomButton").addEventListener("click", addMonkeyImage);

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

        // Set question info
        const questionInfo = document.getElementById("questionInfo");
        if (data.generated) {
            questionInfo.innerText = 'Ez egy korábbi érettségik alapján generált, ellenőrzött feladat.';
        } else {
            const semesterText = data.semester === 1 ? 'tavaszi' : 'őszi';
            questionInfo.innerText = `Ez volt a ${data.number}. feladat a ${data.year}-s ${semesterText} érettségiben.`;
        }

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


function addMonkeyImage() {
    const imageList = [
        "https://cdn-icons-png.flaticon.com/512/4600/4600333.png",
        "https://cdn-icons-png.flaticon.com/512/4600/4600377.png",
        "https://cdn-icons-png.flaticon.com/512/4600/4600357.png",
        "https://cdn-icons-png.flaticon.com/512/4600/4600328.png",
        "https://cdn-icons-png.flaticon.com/512/4600/4600296.png",
        "https://cdn-icons-png.flaticon.com/512/4600/4600270.png",
        "https://cdn-icons-png.flaticon.com/512/4600/4600280.png",
        "https://cdn-icons-png.flaticon.com/512/4600/4600349.png"
    ];

    // Choose a random image from the list
    const imageUrl = imageList[Math.floor(Math.random() * imageList.length)];

    const img = document.createElement("img");
    img.src = imageUrl;
    img.className = "monkey";
    img.style.position = "absolute";

    // Set random size between 50 and 100 pixels
    const randomSize = Math.floor(Math.random() * (100 - 50 + 1)) + 50;
    img.style.width = `${randomSize}px`;
    img.style.height = "auto";

    img.style.top = `${Math.random() * (window.innerHeight - img.height)}px`;
    img.style.left = `${Math.random() * (window.innerWidth - img.width)}px`;

    // Set random angle
    const randomAngle = Math.random() * 360;
    img.style.transform = `rotate(${randomAngle}deg)`;

    document.body.appendChild(img);
}


getRandomQuestion(); // Call the function on page load
