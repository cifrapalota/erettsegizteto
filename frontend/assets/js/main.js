class App {
    constructor() {
      this.bindEvents();
      this.getRandomQuestion();
    }
  
    bindEvents() {
      document.getElementById("getRandomQuestion").addEventListener("click", () => {
        this.getRandomQuestion();
      });
  
      document.getElementById("submitAnswer").addEventListener("click", () => {
        this.checkAnswer();
      });
    }
  
    async getRandomQuestion() {
      try {
        const response = await fetch('/question/random');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
  
        const data = await response.json();
  
        const questionDiv = document.getElementById('question');
        const questionHTML = data.question;
        questionDiv.innerHTML = questionHTML;
  
        const answerContainer = this.createAnswerContainer(data);
        const questionAnswer = document.getElementById('questionAnswer');
        questionAnswer.innerHTML = "";
        questionAnswer.appendChild(answerContainer);
  
        await MathJax.typesetPromise();
  
        const img = document.getElementById('questionImage');
        img.src = data.imageLink;
  
        this.questionID = data.id; // Store the question ID
        this.answerHolders = data.answerHolders; // Store the answer holders
  
        // Set question info
        const questionInfo = document.getElementById("questionInfo");
        if (data.generated) {
          questionInfo.innerText = 'Ez egy korábbi érettségik alapján generált, ellenőrzött feladat.';
        } else {
          const semesterText = data.semester === 1 ? 'tavaszi' : 'őszi';
          questionInfo.innerText = `Ez volt a ${data.number}. feladat a ${data.year}-s ${semesterText} érettségiben.`;
        }
  
      } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
      }
    }
  
    async checkAnswer() {
      const answerInputs = document.getElementsByClassName("answerInput");
      const postData = [];
  
      for (let i = 0; i < answerInputs.length; i++) {
        postData.push({
          answer_holder_id: this.answerHolders[i].id,
          answer: answerInputs[i].value,
        });
      }
  
      try {
        const response = await fetch(`/question/${this.questionID}/check_answers`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(postData),
        });
  
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
  
        const data = await response.json();
  
        for (let i = 0; i < data.length; i++) {
          const answerResult = document.getElementsByClassName("answerResult")[i];
          answerResult.textContent = data[i].answers.join(" vagy ");
          answerResult.style.backgroundColor = data[i].correct ? 'limegreen' : 'red';
        }
      } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
      }
    }
  
    createAnswerContainer(data) {
      const answerHolders = data.answerHolders;
      const container = document.createElement('div');
      container.className = 'container';
  
      for (let i = 0; i < answerHolders.length; i++) {
        const row = document.createElement('div');
        row.className = 'row';
  
        const col1 = document.createElement('div');
        col1.className = 'col-md-3';
        const p1 = document.createElement('p');
        p1.textContent = answerHolders[i].prefix;
        col1.appendChild(p1);
        row.appendChild(col1);

        const col2 = document.createElement('div');
        col2.className = 'col-md-3';
        const input = document.createElement('input');
        input.placeholder = 'válasz';
        input.className = 'answerInput';
        col2.appendChild(input);
        row.appendChild(col2);
  
        const col3 = document.createElement('div');
        col3.className = 'col-md-3';
        const p2 = document.createElement('p');
        p2.textContent = answerHolders[i].suffix;
        col3.appendChild(p2);
        row.appendChild(col3);
  
        const col4 = document.createElement('div');
        col4.className = 'col-md-3';
        const answerResult = document.createElement('p');
        answerResult.className = 'answerResult';
        col4.appendChild(answerResult);
        row.appendChild(col4);
  
        container.appendChild(row);
      }
      return container;
    }
  }
  
  window.addEventListener("DOMContentLoaded", () => {
    const app = new App();
  });
  
       
  