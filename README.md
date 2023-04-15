# Domain
erettsegizteto.hu

# Frameworks:
- platform: heroku
- backend: go
- orm: gorm
- db: postgresql
- frontend: react-bootstrap

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
│   ├── api/
│   │   ├── models/
│   │   │   └── api_answer.go
│   │   │   └── api_question.go
│   │   └── conversion/
│   │       └── conversion.go
│   ├── config/
│   │   └── config.go
│   ├── db/
│   │   └── answer_holder.go
│   │   └── question.go
│   │   └── storage.go
│   ├── handlers/
│   │   └── handler.go
│   │   └── index.go
│   │   └── question.go
│   ├── routers/
│   │   └── router.go
│   └── models/
│       └── answer_holder.go
│       └── answer.go
│       └── question.go
└── frontend/
    ├── assets/
    │   ├── css/
    │   │   └── main.css
    │   ├── img/
    │   │   └── roof.ico
    │   └── js/
    │       └── main.js
    └── templates/
        └── index.html