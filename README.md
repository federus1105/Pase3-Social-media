# 🚀 SOCIAL MEDIA BACK-END API

This API provides basic endpoints for a social media application. The available features include creating posts, liking and commenting on posts, following other users, and viewing posts from followed users

![badge golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![badge postgresql](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![badge redis](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)

## 📸 Preview
> Swagger UI for Dokumentation: [`/swagger/index.html`](http://localhost:8080/swagger/index.html)

## DESIGN SYSTEM
![sistem Design](/SistemDesign/sistemDesign.png)

## 🚀 Features
- 🔐 JWT Authentication (Login & Register)
- 🧠 Redis Caching 
- 📘 Swagger Auto-Generated API Docs
- 📦 PostgreSQL integration
- 🐳 Dockerized (Redis + PostgreSQL)

## 🛠️ Tech Stack
![Go](https://img.shields.io/badge/-Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)
![Gin](https://img.shields.io/badge/-Gin-00ADD8?logo=go&logoColor=white&style=for-the-badge)
![PostgreSQL](https://img.shields.io/badge/-PostgreSQL-4169E1?logo=postgresql&logoColor=white&style=for-the-badge)
![Docker](https://img.shields.io/badge/-Docker-2496ED?logo=docker&logoColor=white&style=for-the-badge)
![Swagger](https://img.shields.io/badge/Swagger-UI-85EA2D?logo=swagger&logoColor=black&style=for-the-badge)
![Redis](https://img.shields.io/badge/Redis-Database-DC382D?logo=redis&logoColor=white&style=for-the-badge)
![Postman](https://img.shields.io/badge/Postman-FF6C37?logo=postman&logoColor=white&style=for-the-badge)


##  🔐 .env Configuration
```
DBUSER=youruser
DBPASS=yourpass
DBHOST=localhost
DBPORT=yourport
DBNAME=yourdbname

JWT_SECRET=your_jwt_secret

REDISUSER=youruser
REDISPASS=yourpass
REDISPORT=yourport

```

## 📦 How to Install & Run
First, clone this repository: 
https://github.com/federus1105/Pase3-Social-media.git
```bash
cd social_media
```
### Install Dependencies
```go
go mod tidy
```
### Run Project
```go
go run .\cmd\main.go 
```

## 📬 Postman Collection

You can try out the API using the Postman collection below:

🔗 [Social Media API Postman Collection](https://federusrudi-9486783.postman.co/workspace/federus-rudi's-Workspace~9cd45016-f25d-441e-8c5a-10f1070df09d/collection/48098195-212b3afa-0860-4558-be0f-4ef879380668?action=share&source=copy-link&creator=48098195)

> Make sure the server is running at `http://localhost:8080`


## 📄 LICENSE

MIT License

Copyright (c) 2025 Social Media API

## 👨‍💻 Made by
### 📬 fedeursrudi@gmail.com

