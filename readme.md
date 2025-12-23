# Project base in roadmap.sh todo list api
- Link: https://roadmap.sh/projects/todo-list-api
    
# Estrutura de Projeto - Clean Architecture + DDD com Go

```bash
todo-list-api-roadmapsh/
├── cmd/
│   └── api/
│       └── main.go                           # Ponto de entrada da aplicação
├── internal/
│   ├── domain/
│   │   ├── entities/
│   │   │   ├── user.go                       # Entidade User rica com comportamentos
│   │   │   └── todo.go                       # Entidade Todo rica com comportamentos
│   │   ├── valueobjects/
│   │   │   ├── user_id.go                    # Value Object para UserID
│   │   │   ├── todo_id.go                    # Value Object para TodoID
│   │   │   ├── email.go                      # Value Object para Email
│   │   │   ├── todo_status.go                # Value Object para Status do Todo
│   │   │   └── priority.go                   # Value Object para Prioridade
│   │   ├── repositories/                     # Apenas interfaces
│   │   │   ├── user_repository.go
│   │   │   └── todo_repository.go
│   │   └── services/                         # Apenas serviços de domínio puros
│   │       ├── auth_domain_service.go        # Serviço de autenticação de domínio
│   │       └── todo_domain_service.go        # Lógica de negócio dos todos
│   ├── usecase/
│   │   ├── auth/
│   │   │   ├── register_user.go              # Use Case de registro
│   │   │   ├── login_user.go                 # Use Case de login
│   │   │   ├── register_user_test.go
│   │   │   ├── login_user_test.go
│   │   │   └── dto.go                        # DTOs específicos
│   │   └── todo/
│   │       ├── create_todo.go                # Use Case criar todo
│   │       ├── update_todo.go                # Use Case atualizar todo
│   │       ├── delete_todo.go                # Use Case deletar todo
│   │       ├── list_todos.go                 # Use Case listar todos
│   │       ├── get_todo.go                   # Use Case obter todo
│   │       ├── create_todo_test.go
│   │       ├── update_todo_test.go
│   │       ├── delete_todo_test.go
│   │       ├── list_todos_test.go
│   │       ├── get_todo_test.go
│   │       └── dto.go                        # DTOs específicos
│   └── infrastructure/
│       ├── db/
│       │   ├── postgres/
│       │   │   ├── user_repository.go       # Implementação Postgres de user.Repository
│       │   │   └── todo_repository.go       # Implementação Postgres de todo.Repository
│       │   └── memory/
│       │       ├── user_repository.go       # Implementação Memory de user.Repository
│       │       └── todo_repository.go       # Implementação Memory de todo.Repository
│       └── handler/                         # Controladores / Handlers (HTTP)
├── pkg/                                     # Pacotes utilitários
├── scripts/
│   └── migrations/
│       ├── 001_create_users_table.up.sql   # Migration para criação da tabela users
│       ├── 001_create_users_table.down.sql # Rollback da migration users
│       ├── 002_create_todos_table.up.sql   # Migration para criação da tabela todos
│       └── 002_create_todos_table.down.sql # Rollback da migration todos
├── go.mod
└── go.sum
```

## Principais decisões de design seguindo Clean Architecture + DDD

### 🏛️ **Domain-Driven Design (DDD)**
- **Entities**: Entidades ricas com comportamentos e validações
- **Value Objects**: Conceitos importantes como `Email`, `TodoStatus`, `Priority`
- **Domain Services**: Lógica de negócio pura sem dependências externas
- **Repositories**: Apenas interfaces no domínio, implementações na infraestrutura

### 🏗️ **Clean Architecture**
- **Use Cases**: Casos de uso específicos organizados por funcionalidade
- **DTOs**: Objetos de transferência específicos para cada use case
- **Separação de camadas**: Domínio independente de infraestrutura
- **Inversão de dependência**: Infraestrutura depende do domínio

### 📁 **Organização por responsabilidade**
- **entities/**: Entidades do negócio com comportamentos
- **valueobjects/**: Value objects imutáveis e validados
- **repositories/**: Interfaces centralizadas de persistência
- **services/**: Serviços de domínio puros
- **usecase/**: Casos de uso organizados por contexto (auth/, todo/)

### 🎯 **Nomenclatura específica e clara**
- `postgres_user_repo.go` - indica tecnologia e responsabilidade
- `auth_domain_service.go` - indica que é serviço de domínio
- `register_user.go` - use case específico e auto-explicativo

### ⚡ **Benefícios**
- **Testabilidade**: Domínio isolado, fácil de testar
- **Escalabilidade**: Estrutura suporta crescimento do projeto
- **Manutenibilidade**: Responsabilidades bem definidas
- **Flexibilidade**: Fácil troca de implementações de infraestrutura
- **DDD**: Linguagem ubíqua e modelagem rica do domínio

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