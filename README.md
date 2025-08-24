<a href="https://opensource.org/" target="_blank" rel="noreferrer">
  <img src="https://img.shields.io/badge/Open%20Source-Initiative-green?style=for-the-badge&logo=open-source-initiative&logoColor=white" alt="Open Source" height="40"/>
</a>

# Go Initial Project 🚀
This project is a **Go-based API** built with **Gin Framework** and **GORM**.
It includes JWT Authentication, a Generic Repository/Service structure, and Swagger documentation.

---

## 📂 Project Structure

```
.
├── config/             # Database & JWT configuration
├── controller/         # HTTP Controllers
├── docs/               # Swagger documentation (generated via swag init)
├── entity/             # Database models
├── middleware/         # JWT & Activity Logger middleware
├── repository/         # Repository layer
├── service/            # Service layer
├── router/             # Router definitions
└── main.go             # Entry point
```

---

## ⚡ Setup

### 1. Clone the Repository
```bash
git clone https://github.com/imcanugur/go-initial-project.git
cd go-initial-project
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Swagger Setup
Install `swag` CLI for Swagger:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate Swagger files:
```bash
swag init -g main.go
```

### 4. Run the Project
```bash
go run main.go
```

---

## 🔑 API Authentication

JWT Authentication is used.
- `/api/auth/login` → login and get token
- `/api/auth/register` → create a new user
- `/api/auth/me` → get user info with token

Header:
```
Authorization: Bearer <token>
```

---

## 📖 Swagger Documentation

After running the server, visit:
👉 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## ✅ Features

- [x] JWT Authentication
- [x] Generic Repository & Service
- [x] Middleware (Auth + Activity Logger)
- [x] Swagger Integration
- [x] Docker Support

---

## 📜 License

MIT License © 2025
