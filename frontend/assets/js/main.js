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
      this.hideQuestionWorkings();

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
    
    // Check if MathJax.typesetPromise is available, and if not, wait for it
    const typesetMathJax = async () => {
      if (typeof MathJax.typesetPromise === 'function') {
        await MathJax.typesetPromise();
      } else {
        setTimeout(typesetMathJax, 500);
      }
    };
    await typesetMathJax();
    
        this.questionID = data.id; // Store the question ID
        this.answerHolders = data.answerHolders; // Store the answer holders
    
        // Set question info
        const questionInfo = document.getElementById("questionInfo");
        if (data.generated) {
          questionInfo.innerText = 'Ez egy korábbi érettségik alapján generált, ellenőrzött feladat.';
        } else {
          const semesterText = data.semester === 1 ? 'tavaszi' : 'őszi';
          questionInfo.innerText = `Ez a ${data.year}-es ${semesterText} érettségi ${this.ordinalSuffix(data.number)} feladata.`;
        }

        document.getElementById("submitAnswer").disabled = false;
    
      } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
      }
    }    

    async checkAnswer() {
      const answerInputs = document.getElementsByClassName("answerInput");
      const postData = [];
      let hasEmptyInputs = false;
    
      for (let i = 0; i < answerInputs.length; i++) {
        // Check if the input is empty
        if (answerInputs[i].value.trim() === "") {
          hasEmptyInputs = true;
      
          // Create an alert-primary div for the empty input warning
          const emptyInputAlert = document.createElement("div");
          emptyInputAlert.className = "alert alert-primary mt-2 alert-dismissible fade show"; // Add alert-dismissible and fade show classes
          emptyInputAlert.textContent = "Töltsd ki az összes válaszmezőt!";
      
          // Create a close button for the alert
          const closeButton = document.createElement("button");
          closeButton.type = "button";
          closeButton.className = "btn-close";
          closeButton.setAttribute("data-bs-dismiss", "alert");
          closeButton.setAttribute("aria-label", "Close");
          emptyInputAlert.appendChild(closeButton);
      
          const answerResult = document.getElementsByClassName("answerResult")[i];
          // Remove the previous alert-primary div if it exists
          const previousAlert = answerResult.querySelector(".alert");
          if (previousAlert) {
            answerResult.removeChild(previousAlert);
          }
      
          // Append the new alert-primary div
          answerResult.appendChild(emptyInputAlert);
        } else {
          postData.push({
            answerHolderId: this.answerHolders[i].id,
            answer: answerInputs[i].value,
          });
        }
      }
    
      if (hasEmptyInputs) {
        return; // Exit the checkAnswer method to avoid submitting an incomplete answer
      }

      this.hideQuestionWorkings();
    
      try {
        const response = await fetch(`/question/${this.questionID}/check-answers`, {
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
          const input = answerInputs[i]; // Get the answer input
        
          if (data.checkedAnswers[i].correct) {
            answerResult.textContent = "";
            input.style.backgroundColor = '#d4edda';
            input.classList.remove('strikethrough');
          } else {
            input.style.backgroundColor = '#f8d7da';
            input.classList.add('strikethrough');
        
            // Create an alert-warning div for the correct answers
            const correctAnswerDiv = document.createElement("div");
            correctAnswerDiv.className = "alert alert-warning mt-2";
            correctAnswerDiv.textContent = "A helyes válasz: " + data.checkedAnswers[i].answers.join(" vagy ");
        
            // Remove the previous alert-warning div if it exists
            const previousAlert = answerResult.querySelector(".alert");
            if (previousAlert) {
              answerResult.removeChild(previousAlert);
            }
        
            // Append the new alert-warning div
            answerResult.appendChild(correctAnswerDiv);
          }
        
          // Disable the answer input when the submitAnswer button is disabled
          input.disabled = true;
        }

        document.getElementById("submitAnswer").disabled = true;
    
        // Store the workings data
        // this.workingsData = data.workings;

        const questionWorkings = document.getElementById("questionWorkings");
        const workings = this.convertMarkdownToHTML(data.workings);
        questionWorkings.innerHTML = workings;
      
        await MathJax.typesetPromise();
    
        // Enable the "showWorkings" button
        const showWorkingsButton = document.getElementById("showWorkings");
        showWorkingsButton.style.display = "inline-block";
        showWorkingsButton.disabled = false;
    
      } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
      }
    }       
    
    hideQuestionWorkings() {
      const questionWorkingsElement = document.getElementById('questionWorkings');
      questionWorkingsElement.innerHTML = "A levezetés megtekintéséhez először küldj be egy megoldást, majd nyomj újra a 'Levezetés!' gombra.";
      const collapseQuestionWorkingsElement = document.getElementById('collapseQuestionWorkings');
      collapseQuestionWorkingsElement.classList.remove('show');
    }

    convertMarkdownToHTML(rawContent) {
      // Add this line to handle undefined rawContent
      if (!rawContent) return "";

      const classMap = {
        img: 'img-fluid mx-auto d-block q-image',
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
      container.className = "row"; // Add the 'justify-content-center' class here
    
      // Create a single column and append it to the container
      const col = document.createElement("div");
      // col.className = "col-12 col-md-8 offset-md-2";
      container.appendChild(col);
    
      for (let i = 0; i < answerHolders.length; i++) {
        const inputGroup = document.createElement("div");
        inputGroup.className = "input-group my-2"; // Add margin for spacing between input groups
    
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
        inputGroup.appendChild(input);
    
        if (answerHolders[i].suffix) {
          const suffix = document.createElement("span");
          suffix.textContent = answerHolders[i].suffix;
          suffix.className = "suffix";
          inputGroup.appendChild(suffix);
        }
    
        col.appendChild(inputGroup);
    
        if (answerHolders[i].help && answerHolders[i].help.trim() !== "") {
          const helpDiv = document.createElement("div");
          helpDiv.className = "alert alert-info mt-2 d-flex justify-content-between";
          helpDiv.textContent = answerHolders[i].help;
    
          const closeButton = document.createElement("button");
          closeButton.type = "button";
          closeButton.className = "btn-close";
          closeButton.setAttribute("data-bs-dismiss", "alert");
          closeButton.setAttribute("aria-label", "Close");
    
          helpDiv.appendChild(closeButton);
          col.appendChild(helpDiv);
        }
    
        const result = document.createElement("p");
        result.className = "answerResult";
        col.appendChild(result);
      }
      return container;
    }      
    

    ordinalSuffix(i) {
      const suffixes = [
        '',
        'első',
        'második',
        'harmadik',
        'negyedik',
        'ötödik',
        'hatodik',
        'hetedik',
        'nyolcadik',
        'kilencedik',
        'tizedik',
        'tizenegyedik',
        'tizenkettedik',
        'tizenharmadik',
        'tizennegyedik',
        'tizenötödik',
        'tizenhatodik',
        'tizenhetedik',
        'tizennyolcadik',
        'tizenkilencedik',
        'huszadik'
      ];
    
      if (i >= 1 && i <= 20) {
        return `${suffixes[i]}`;
      }
      return `${i}.`;
    }    

  }
  
  window.addEventListener("DOMContentLoaded", () => {
    const app = new App();
  });