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

# Endpoints:
/ -> returns index.html
/question with param id -> returns a question by id

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
    │   ├── img/
    │   └── js/
    │       └── main.js
    └── templates/
        ├── base.html
        └── index.html