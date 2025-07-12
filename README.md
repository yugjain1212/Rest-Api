# Rest-Api
A simple and clean REST API built in Go (Golang) to practice HTTP routing, JSON handling, and modular code structure.

# Go REST API Project 🚀

This is a simple **REST API** built using **Go (Golang)** to practice web development concepts such as HTTP routing, JSON handling, and modular code structure.  

It demonstrates how to build a clean, idiomatic Go web service that supports basic CRUD operations.

---

## 📋 Features
✅ Clean and idiomatic Go code  
✅ RESTful endpoints (GET, POST, PUT, DELETE)  
✅ JSON request and response handling  
✅ Modular package structure  
✅ Uses Go’s standard library — no external frameworks required  
✅ Concurrency-ready and lightweight

---

## 📂 Project Structure

/your-project
├── main.go            # entry point
├── handlers/          # HTTP handlers
├── models/            # data models & DB operations
├── routes/            # routing setup
├── database/          # SQLite connection
├── utils/             # helper functions (if any)
├── mydb.sqlite        # SQLite database file
└── README.md
---

## 🚀 Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) (>= 1.20 recommended)
- [SQLite](https://www.sqlite.org/index.html) (optional, already included as a file)
- [TablePlus](https://tableplus.com/) (optional, to view/manage the database)

### Install SQLite driver
bash
go get github.com/mattn/go-sqlite3
## 🚀 Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) installed (version >= 1.20 recommended)

# Run the app
go run main.go

#Server start at
http://localhost:8081
