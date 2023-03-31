# Domain
erettsegizteto.hu

# Frameworks:
- platform: heroku
- backend: go
- orm: gorm
- db: postgresql
- frontend: react-bootstrap

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

# Endpoints:
/
/question/:questionID
/question/random

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
│   │   └── handler.go
│   │   └── index.go
│   │   └── question.go
│   ├── routers/
│   │   └── router.go
│   └── models/
│       └── question.go
└── frontend/
    ├── assets/
    │   ├── css/
    │   │   └── main.css
    │   └── js/
    │       └── main.js
    └── templates/
        └── index.html