function getNextQuestion() {
    var questionContainer = document.getElementById("question-container");
    var request = new XMLHttpRequest();
    request.onreadystatechange = function() {
        if (this.readyState == 4) {
            if (this.status == 200) {
                var question = JSON.parse(this.responseText);
                var questionText = document.createTextNode(question.question);
                questionContainer.innerHTML = '';
                questionContainer.appendChild(questionText);
            } else if (this.status == 500) {
                questionContainer.innerHTML = "Nincs tobb kerdes!";
            }
        }
    };
    request.onerror = function() {
        questionContainer.innerHTML = "Nincs tobb kerdes!";
    };
    request.open("GET", "/question?id=" + getNextQuestion.questionIndex, true);
    request.send();
    getNextQuestion.questionIndex++;
}

getNextQuestion.questionIndex = 1;
