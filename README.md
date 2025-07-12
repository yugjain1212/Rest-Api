# Student Management REST API 🚀

This is a simple **REST API** built using **Go (Golang)** and **SQLite** for managing student records.  
It demonstrates how to build a clean, idiomatic Go web service with persistent storage, supporting full CRUD (Create, Read, Update, Delete) operations for students.

I also used **TablePlus** to easily view and manage the SQLite database.

---

## 📋 Features
✅ Manage student records with RESTful API endpoints  
✅ SQLite database for data persistence  
✅ JSON request & response handling  
✅ Modular and clean code structure  
✅ Uses Go’s standard library + `database/sql` + `github.com/mattn/go-sqlite3` driver  
✅ Easy to set up & run

---

## 📂 Project Structure

/student-api
├── main.go            # Entry point
├── handlers/          # HTTP handlers for students
├── models/            # Student data models & DB operations
├── routes/            # Routing setup
├── database/          # SQLite connection
├── mydb.sqlite        # SQLite database file
└── README.md


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
