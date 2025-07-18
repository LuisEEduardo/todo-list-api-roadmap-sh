# Project base in roadmap.sh todo list api
- Link: https://roadmap.sh/projects/todo-list-api
    
# Estrutura de Projeto - Clean Architecture com Go

```bash
go-clean-architecture/
├── cmd/
│   └── api/
│       └── main.go                     # Ponto de entrada da aplicação
├── internal/
│   ├── domain/
│   │   └── user.go                     # Entidades e interfaces do domínio
│   ├── usecase/
│   │   └── user_usecase.go            # Lógica de negócio (casos de uso)
│   ├── interface/
│   │   ├── handler/
│   │   │   └── user_handler.go        # Controladores / Handlers (HTTP)
│   │   └── repository/
│   │       └── user_repository.go     # Interfaces dos repositórios
│   └── infrastructure/
│       └── db/
│           └── postgres.go            # Implementação de acesso ao banco (PostgreSQL)
├── scripts/
│   └── migrations/
│       ├── 001_create_users_table.up.sql    # Migration para criação de tabela
│       └── 001_create_users_table.down.sql  # Rollback da migration
├── go.mod
└── go.sum
```

# Goals
The skills you will learn from this project include:
- User authentication
- Schema design and Databases
- RESTful API design
- CRUD operations
- Error handling
- Security

# Requirements
You are required to develop a RESTful API with following endpoints:

- User registration to create a new user
- Login endpoint to authenticate the user and generate a token
- CRUD operations for managing the to-do list
- Implement user authentication to allow only authorized users to access the to-do list
- Implement error handling and security measures
- Use a database to store the user and to-do list data (you can use any database of your choice)
- Implement proper data validation
- Implement pagination and filtering for the to-do list