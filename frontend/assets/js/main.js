function getNextQuestion() {
    var questionContainer = document.getElementById("question-container");
    var id = questionContainer.childElementCount + 1;
  
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        var question = JSON.parse(this.responseText);
        var questionDiv = document.createElement("div");
        questionDiv.innerHTML = "<p>" + question.question + "</p>";
        questionContainer.appendChild(questionDiv);
      }
    };
    xhttp.open("GET", "/question?id=" + id, true);
    xhttp.send();
  }
  