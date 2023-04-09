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
│       └── answer_holder.go
│       └── answer.go
│       └── question.go
└── frontend/
    ├── assets/
    │   ├── css/
    │   │   └── main.css
    │   └── js/
    │       └── main.js
    └── templates/
        └── index.html