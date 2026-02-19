Task Management REST API (Golang)

Overview

A production-style Task Management REST API built using Golang, Gin,
PostgreSQL, and JWT authentication.

Features

-   User Registration
-   Login with JWT authentication
-   Password hashing using bcrypt

-   Create, Get, Delete Tasks
-   Background worker for auto-completion
-   Configurable environment variables
-   Clean layered architecture

Project Structure

main.go router/router.go controller/controller.go services/service.go
services/auth_service.go model/model.go model/db.go
middleware/middleware.go

Tech Stack

-   Go (Golang)
-   Gin Web Framework
-   PostgreSQL
-   GORM
-   JWT (golang-jwt)
-   bcrypt

Environment Variables (.env)

DB_DSN=host=localhost user=postgres password=postgres dbname=tasks
port=5432 sslmode=disable JWT_SECRET=mysecretkey
TASK_AUTO_COMPLETE_MINUTES=10

How To Run

1.  Install dependencies: go mod tidy

2.  Create database in PostgreSQL: CREATE DATABASE tasks;

3.  Run application: go run main.go

Server runs at: http://localhost:8080

API Endpoints

Public Routes:

POST /register Body: { “email”: “user@test.com”, “password”: “123456”,
“role”: “user” }

POST /login Body: { “email”: “user@test.com”, “password”: “123456” }

Response: { “token”: “JWT_TOKEN” }

Protected Routes (Require JWT Header): Authorization: Bearer YOUR_TOKEN

POST /tasks GET /tasks DELETE /tasks/{id}

Background Worker

Tasks automatically change to “completed” after
TASK_AUTO_COMPLETE_MINUTES if they remain in “pending” or “in_progress”.

Concurrency Design

-   Task IDs are pushed to a channel
-   Worker listens on channel
-   Worker sleeps for configured duration
-   Updates task status if still incomplete
-   Non-blocking and thread-safe

Assignment Requirements Covered

-   REST APIs
-   SQL Persistence
-   Clean Architecture
-   JWT Authentication
-   Role-based Authorization
-   Background Worker using Goroutines
-   Configurable Environment
-   Proper Error Handling

Author

Developed as part of CashInvoice Golang Assignment.
