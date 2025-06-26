

# 🚀 URL Shortener – Go + Gin + Redis + Docker

A lightweight and powerful URL shortening service built with **Golang**, **Gin Web Framework**, **Redis**, and **Docker**. It exposes RESTful APIs to shorten, retrieve, tag, edit, and delete URLs—ideal for modern web applications.

---

## 📁 Project Structure

```
url-shortener/
├── api/
│   ├── database/          # Redis connection and helper functions
│   ├── utils/             # Utility functions (e.g., short ID generator)
│   └── router/            # API route handlers
│       ├── get.go         # Get original URL by short ID
│       ├── shorten.go     # Create a shortened URL
│       ├── editurl.go     # Edit an existing short URL
│       ├── tag.go         # Tag a URL
│       └── deleteurl.go   # Delete a URL
│
├── DB/
│   ├── redis.Dockerfile       # Redis container configuration
│   ├── app.Dockerfile         # Go application Dockerfile
│   └── docker-compose.yml     # Compose file for Redis + App setup
│
├── main.go                # Entry point of the application
└── README.md              # Project documentation
```

---

## ✨ Features

- 🔗 **Shorten URLs** – Create unique short links from long URLs.
- 🧭 **Retrieve URLs** – Resolve short links to original URLs.
- ✏️ **Edit URLs** – Modify existing shortened URL entries.
- 🏷️ **Tag URLs** – Attach tags to your links for easy categorization.
- 🗑️ **Delete URLs** – Remove short URLs from the system.

---

## 📡 API Endpoints

| Method   | Endpoint                | Description                         |
|----------|-------------------------|-------------------------------------|
| `GET`    | `/api/v1/:shortid`      | Retrieve original URL from short ID |
| `POST`   | `/api/v1/shorten`       | Create a new shortened URL          |
| `PUT`    | `/api/v1/editurl`       | Edit an existing short URL          |
| `POST`   | `/api/v1/tag`           | Add tag(s) to a shortened URL       |
| `DELETE` | `/api/v1/deleteurl/:id` | Delete a shortened URL              |

> All endpoints are versioned under `/api/v1/`.

---

## 🐳 Dockerized Setup

**1. Redis Dockerfile**
- Sets up a Redis instance to store shortened URL mappings.

**2. App Dockerfile**
- Builds and packages the Go URL shortener app.

**3. Docker Compose**
- Boots up both Redis and the Go app in linked containers.

---

## 🚀 Getting Started

### 🔧 Prerequisites
- Docker & Docker Compose installed
- Go installed (optional, if running locally without Docker)

### 📦 Clone the Repository

```bash
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

### 🐳 Run with Docker Compose

```bash
docker-compose up --build
```

> This builds and starts both the Redis container and the Go app together.

---

## 🧰 Tech Stack

- ⚙️ **Golang** – Robust and performant backend language
- 🌐 **Gin** – Fast, minimalistic HTTP web framework for Go
- ⚡ **Redis** – High-speed in-memory database
- 🐳 **Docker** – Seamless containerization
- 📦 **Docker Compose** – Multi-container orchestration

---

## 🔮 Future Enhancements

- 📊 Track URL click stats
- 🔐 Authentication & rate limiting
- ⏱️ Expiry for time-bound URLs
- ✍️ Custom aliases and branded short domains

---

## 💡 Developer Notes

### Passing Values vs Pointers in Go

#### ❗️ By Value:
```go
func changeValue(val int) {
    val = 10
}
```
- Creates a **copy** of `val`.  
- Changes are **not reflected** in the original variable.

#### ✅ By Pointer:
```go
func changeValue(val *int) {
    *val = 10
}
```
- Works with the **original memory reference**.  
- Changes are **reflected globally**.

### 🔍 Why use `*gin.Engine`?

```go
func setupRouters(router *gin.Engine)
```
- You’re passing a **pointer** to the actual Gin engine.
- Allows you to **add routes** to the original instance in `main.go`.

> If passed **by value**, routes would be added to a copy—not the running server.

🧠 Think of `*gin.Engine` like editing a **shared Google Doc**, whereas `gin.Engine` alone is like editing a **printed copy**.

---

## 🤝 Contributing

Contributions, suggestions, and forks are always welcome!  
Please feel free to open issues or submit pull requests.

---

## 📄 License

Licensed under the **MIT License** – use it freely!

---


## 💡 Developer Notes

## Basics:
1. `go mod init github.com/yourusername/gin-demo` is used to initialize go module. If one plans to publish module in GitHub it is usefull
2. `*gin.Context` is the heart of each request in Gin. You need it to read the request and send the response. That's why you always pass it into your handler functions.


## Passing Values vs Pointers in Go

#### ❗️ By Value:
```go
func changeValue(val int) {
    val = 10
}
```
- Creates a **copy** of `val`.  
- Changes are **not reflected** in the original variable.

#### ✅ By Pointer:
```go
func changeValue(val *int) {
    *val = 10
}
```
- Works with the **original memory reference**.  
- Changes are **reflected globally**.

### 🔍 Why use `*gin.Engine`?

```go
func setupRouters(router *gin.Engine)
```
- You’re passing a **pointer** to the actual Gin engine.
- Allows you to **add routes** to the original instance in `main.go`.

> If passed **by value**, routes would be added to a copy—not the running server.

🧠 Think of `*gin.Engine` like editing a **shared Google Doc**, whereas `gin.Engine` alone is like editing a **printed copy**.

Perfect! Let me show you examples of how to use `*gin.Context` to:

1. Get **query parameters**
2. Get **JSON body**
3. Get **headers**
4. Get **URL parameters**
5. Send different types of **responses**

---
---
# GIN
## ✅ 1. Get Query Parameters

### 📥 Example URL:

```
GET /hello?name=Abdul
```

### 🔧 Code:

```go
r.GET("/hello", func(c *gin.Context) {
    name := c.Query("name") // ?name=Abdul
    c.JSON(200, gin.H{
        "message": "Hello " + name,
    })
})
```

---

## ✅ 2. Get JSON Body from POST request

### 📥 POST JSON:

```json
{
  "username": "abdul",
  "email": "abdul@example.com"
}
```

### 🔧 Code:

```go
type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
}

r.POST("/register", func(c *gin.Context) {
    var user User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": "Invalid JSON"})
        return
    }
    c.JSON(200, gin.H{"status": "Registered", "user": user})
})
```

---

## ✅ 3. Get Headers

### 🔧 Code:

```go
r.GET("/check-header", func(c *gin.Context) {
    token := c.GetHeader("Authorization")
    c.JSON(200, gin.H{"token_received": token})
})
```

---

## ✅ 4. Get Path Parameters

### 📥 Example URL:

```
GET /user/42
```

### 🔧 Code:

```go
r.GET("/user/:id", func(c *gin.Context) {
    userID := c.Param("id") // grabs '42'
    c.JSON(200, gin.H{"user_id": userID})
})
```

---

## ✅ 5. Send Different Responses

| Task          | Code                                                |
| ------------- | --------------------------------------------------- |
| JSON response | `c.JSON(200, gin.H{"message": "Hi"})`               |
| Plain text    | `c.String(200, "Hello")`                            |
| HTML response | `c.HTML(200, "index.tmpl", gin.H{"title": "Home"})` |
| Status only   | `c.Status(204)` (No Content)                        |

---

Let me know if you want to try handling **form submissions**, **file uploads**, or **cookies** with `*gin.Context` too!

---
----


## 🧱 Using Redis in Go

This project uses [go-redis v9](https://github.com/redis/go-redis) as the Redis client.

### 📦 Installation

Install the Redis client library:

```bash
go get github.com/redis/go-redis/v9
```

---

### 📁 Setup (`database/redis.go`)

```go
package database

import (
    "context"
    "os"

    "github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("DB_ADDR"), // e.g., "localhost:6379"
        Password: os.Getenv("DB_PASS"), // leave blank if no password
        DB:       dbNo,                 // Redis DB number (0–15)
    })
    return rdb
}
```

---

### 🔧 Store & Retrieve Data

#### ➕ Set a Key

```go
rdb := database.CreateClient(0)
err := rdb.Set(database.Ctx, "name", "Abdul", 0).Err()
```

#### 🔍 Get a Key

```go
val, err := rdb.Get(database.Ctx, "name").Result()
fmt.Println("name is", val)
```

#### ⛔ Delete a Key

```go
rdb.Del(database.Ctx, "name")
```

#### ⏳ Set a Key with Expiry

```go
rdb.Set(database.Ctx, "session:123", "loggedin", time.Minute*10)
```

---

### 🧪 Testing with Redis CLI (Optional)

If Redis is installed locally:

```bash
redis-cli
```

```bash
127.0.0.1:6379> SET name Abdul
OK
127.0.0.1:6379> GET name
"Abdul"
```

---
