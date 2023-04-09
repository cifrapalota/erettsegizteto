class App {
    constructor() {
      this.correctAnswer = '';
  
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
  
        const table = this.createAnswerTable(data);
        const questionAnswer = document.getElementById('questionAnswer');
        questionAnswer.innerHTML = "";
        questionAnswer.appendChild(table);
  
        await MathJax.typesetPromise();
  
        const img = document.getElementById('questionImage');
        img.src = data.imageLink;
    
        this.questionID = data.id; // Store the question ID
        this.answerHolders = data.answerHolders; // Store the answer holders
        
        this.correctAnswer = data.answer;
  
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
            answerResult.textContent = data[i].answers.join(", ");
            answerResult.style.backgroundColor = data[i].correct ? 'limegreen' : 'red';
          }
        } catch (error) {
          console.error('There has been a problem with your fetch operation:', error);
        }
      }
  
    createAnswerTable(data) {
      const answerHolders = data.answerHolders;
      const table = document.createElement('table');
      for (let i = 0; i < answerHolders.length; i++) {  
        const tr = document.createElement('tr');
        const td1 = document.createElement('td');
        td1.className = 'answerPrefix';
        const p1 = document.createElement('p');
        p1.textContent = answerHolders[i].prefix;
        td1.appendChild(p1);
        tr.appendChild(td1);
        const td2 = document.createElement('td');
        td2.className = 'answerText';
        const input = document.createElement('input');
        input.placeholder = 'válasz';
        input.className = 'answerInput';
        td2.appendChild(input);
        tr.appendChild(td2);
        const td3 = document.createElement('td');
        td3.className = 'answerSuffix';
        const p2 = document.createElement('p');
        p2.textContent = answerHolders[i].suffix;
        td3.appendChild(p2);
        tr.appendChild(td3);
        const td4 = document.createElement('td');
        td4.className = 'answerResult';
        tr.appendChild(td4);
        table.appendChild(tr);
      }
      return table;
    }
  };
  
  window.addEventListener("DOMContentLoaded", () => {
    const app = new App();
  });
    