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
│   │   │   └── answer_holder.go
│   │   │   └── checked_answer_holder.go
│   │   │   └── question.go
│   │   │   └── user_answer.go
│   │   └── conversion/
│   │       └── conversion.go
│   ├── config/
│   │   └── config.go
│   ├── db/
│   │   └── answer_holder.go
│   │   └── question.go
│   │   └── db.go
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