# Auth API

**Descrição**
API de autenticação construída com Go, Gin e PostgreSQL. Suporta registro de usuários, login com JWT e validação de tokens. Ideal para integração em sistemas web ou mobile.

---

## Tecnologias

- **Go** – linguagem principal
- **Gin** – framework web
- **PostgreSQL** – banco de dados relacional
- **JWT** – autenticação baseada em token
- **Docker** – para containerização (opcional)

---

## Funcionalidades

- Cadastro de usuário com validação de e-mail
- Login com geração de JWT
- Middleware de autenticação para rotas protegidas
- Recuperação de senha (opcional)

---

## Pré-requisitos

- Go 1.21+
- PostgreSQL 15+
- Docker e Docker Compose (opcional, para facilitar o setup)

---

## Configuração

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/auth-api.git
cd auth-api
```

2. Crie o arquivo `.env` na raiz:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=auth_db
JWT_SECRET=uma_chave_secreta
```

3. Inicialize o banco de dados (local ou via Docker):

```bash
docker-compose up -d
```

4. Rode a aplicação:

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

---

## Endpoints principais

| Método | Rota      | Descrição                                |
| ------ | --------- | ---------------------------------------- |
| POST   | /register | Registrar novo usuário                   |
| POST   | /login    | Realizar login                           |
| GET    | /profile  | Obter informações do usuário (protegido) |

---

## Estrutura do projeto (exemplo)

```
auth-api/
├─ cmd/               # Arquivo principal
├─ internal/
│  ├─ controllers/    # Handlers das rotas
│  ├─ models/         # Modelos do banco
│  ├─ repositories/   # Operações com banco
│  ├─ services/       # Lógica de negócio
│  └─ middlewares/    # JWT, autenticação, etc
├─ config/            # Configurações e .env
└─ Dockerfile
```

---

## Contribuição

Pull requests são bem-vindos! Para mudanças maiores, abra uma issue primeiro para discutir o que deseja modificar.

---

## Licença

MIT License
