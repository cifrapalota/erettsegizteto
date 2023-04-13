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

      document.getElementById("showSolution").addEventListener("click", () => {
        this.displaySolution();
      });      
    }
  
    async getRandomQuestion() {
      try {
        const response = await fetch('/question/random');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
    
        const data = await response.json();
    
        // Convert Markdown to HTML using the convertMarkdownToHTML method
        const questionHTML = this.convertMarkdownToHTML(data.question);
    
        const questionDiv = document.getElementById('question');
        questionDiv.innerHTML = questionHTML;
    
        const answerContainer = this.createAnswerContainer(data);
        const questionAnswer = document.getElementById('questionAnswer');
        questionAnswer.innerHTML = "";
        
        questionAnswer.appendChild(answerContainer);
    
        await MathJax.typesetPromise();
    
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

      // Hide the "showSolution" button and clear the questionSolution div
      const showSolutionButton = document.getElementById("showSolution");
      showSolutionButton.style.display = "none";
      const questionSolution = document.getElementById("questionSolution");
      questionSolution.innerHTML = "";

        // Add this line to hide the questionSolution div when fetching a new question
  questionSolution.style.display = "none";
    
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

        await MathJax.typesetPromise();
  
        for (let i = 0; i < data.checkedAnswers.length; i++) {
          const answerResult = document.getElementsByClassName("answerResult")[i];
          answerResult.textContent = data.checkedAnswers[i].answers.join(" vagy ");
          answerResult.style.backgroundColor = data.checkedAnswers[i].correct ? 'limegreen' : 'red';
        }


        // Store the solution data
        this.solutionData = data.solution;

        // Show the "showSolution" button
        const showSolutionButton = document.getElementById("showSolution");
        showSolutionButton.style.display = "inline-block";

      } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
      }
    }

    async displaySolution() {
      const questionSolution = document.getElementById("questionSolution");
    
      // Add this line to initially hide the questionSolution div
      questionSolution.style.display = "none";
    
      const solution = this.convertMarkdownToHTML(this.solutionData);
      questionSolution.innerHTML = solution;
    
      await MathJax.typesetPromise();
    
      // Add this line to show the questionSolution div when the button is pressed
      questionSolution.style.display = "block";
    }  

    convertMarkdownToHTML(rawContent) {
      // Add this line to handle undefined rawContent
      if (!rawContent) return "";

      const classMap = {
        img: 'img-fluid mx-auto d-block',
        table: 'table table-bordered',
        thead: 'd-none',
        // Add the wrapper for responsive tables
        'table-responsive': 'table-responsive',
      };
    
      const bindings = Object.keys(classMap)
        .map(key => ({
          type: 'output',
          regex: new RegExp(`<${key}(.*)>`, 'g'),
          replace: `<${key} class="${classMap[key]}" $1>`
        }));
    
      const converter = new showdown.Converter({
        extensions: [...bindings],
        tables: true,
        strikethrough: true
      });
    
      const htmlContent = converter.makeHtml(rawContent);
    
      // Wrap the table inside a div with the class 'table-responsive'
      return htmlContent.replace(/<table([^>]*)>/, '<div class="table-responsive"><table$1>').replace(/<\/table>/, '</table></div>');
    }    
  
    createAnswerContainer(data) {
      const answerHolders = data.answerHolders;
      const container = document.createElement("div");
      container.className = "row";
    
      for (let i = 0; i < answerHolders.length; i++) {
        const col = document.createElement("div");
        col.className = "col-6 col-md-3";
    
        const inputGroup = document.createElement("div");
        inputGroup.className = "input-group";
    
        if (answerHolders[i].prefix) {
          const prefix = document.createElement("span");
          prefix.textContent = answerHolders[i].prefix;
          prefix.className = "prefix";
          inputGroup.appendChild(prefix);
        }
    
        const input = document.createElement("input");
        input.type = "text";
        input.className = "form-control answerInput";
        input.placeholder = "válasz";
    
        if (answerHolders[i].help && answerHolders[i].help.trim() !== "") {
          input.setAttribute("data-bs-placement", "top");
          input.setAttribute("data-bs-trigger", "focus");
          input.setAttribute("title", "Segítség");
          input.setAttribute("data-bs-content", answerHolders[i].help);
        }
    
        inputGroup.appendChild(input);
    
        if (answerHolders[i].suffix) {
          const suffix = document.createElement("span");
          suffix.textContent = answerHolders[i].suffix;
          suffix.className = "suffix";
          inputGroup.appendChild(suffix);
        }
    
        col.appendChild(inputGroup);
    
        const result = document.createElement("p");
        result.className = "answerResult";
        col.appendChild(result);
    
        container.appendChild(col);
      }
      return container;
    }

  }
  
  window.addEventListener("DOMContentLoaded", () => {
    const app = new App();
  });