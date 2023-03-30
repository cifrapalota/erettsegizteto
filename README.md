# Domain
erettsegizteto.hu

# Frameworks:
- platform: heroku
- backend: go
- orm: gorm
- db: postgresql
- frontend: bootstrap

# Database structure
schema: db_erettsegizteto
table name:  questions
fields:
{
    "id":	        "serial4"
    "question":	    "text"
    "answer":	    "text"
    "generated":	"boolean"
    "year":	        "smallint"
    "semester":	    "smallint"
    "number":	    "smallint"
}

# BDD:
GIVEN: I am a user
WHEN: I open the main page
THEN: The main page displays a random question, which is selected from the database, with an ID between 1 and 3 (inclusive).

# Layout:
erettsegizteto/
├── .gitignore
├── go.mod
├── go.sum
├── Procfile
├── README.md
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── storage/
│   │   └── storage.go
│   ├── handlers/
│   │   └── question.go
│   ├── routers/
│   │   └── router.go
│   └── models/
│       └── question.go
└── frontend/
    ├── assets/
    │   ├── css/
    │   │   └── main.css
    │   ├── img/
    │   └── js/
    │       └── main.js
    └── templates/
        ├── base.html
        └── index.html