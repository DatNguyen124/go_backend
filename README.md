# 📚 Bookstore RESTful API (Golang + PostgreSQL)

A simple backend REST API built with Golang and PostgreSQL to manage a list of books.  
This project is currently tested via Postman and supports full CRUD operations.

> ✅ Tested with Postman  
> ✅ Uses PostgreSQL via GORM  
> ⏳ Docker support coming soon

---

## 📌 Features

- Add a new book
- Get all books
- Get book by ID
- Update book by ID
- Delete book by ID

---

## 🛠 Tech Stack

- **Language:** Go (Golang)
- **Framework:** net/http + Gorilla Mux
- **ORM:** GORM
- **Database:** PostgreSQL
- **API Testing:** Postman

---

## 📁 Project Structure
```
│   go.mod
│   go.sum
│   main.go
│
├───cmd
│   └───api
├───database
└───pkg
    ├───config
    │       app.go
    │
    ├───controllers
    │       book-controller.go
    │
    ├───models
    │       book.go
    │
    ├───routes
    │       bookstore-routes.go
    │
    └───utils
            utils.go
```
---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/DatNguyen124/go_backend.git
cd go_backend
```
---

### 2. 🧰 Install dependencies

```bash
go mod tidy
```

---

### 3. 🛠 Setup PostgreSQL

- Ensure PostgreSQL is running locally
- Create a database named:

```bash
bookstore
```

- Update DB credentials in ```pkg/config/app.go:```

```bash
gorm.Open("postgres", "host=localhost port=5432 user=postgres password=yourpassword dbname=bookstore sslmode=disable")
```

### 4. 🚀 Run the server

```bash
go run main.go
```
Server will start on:
```bash
http://localhost:8080
```

