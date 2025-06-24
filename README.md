---

# 🚀 URL Shortener – Go + Gin + Redis + Docker

A powerful and lightweight URL shortener service built with **Golang**, **Gin Web Framework**, **Redis** for fast storage, and **Docker** for containerized deployment. It provides RESTful APIs to shorten, retrieve, tag, edit, and delete URLs seamlessly.

---

## 📁 Project Structure

```
url-shortener/
│
├── api/
│   ├── database/     # Code to connect and interact with Redis
│   ├── utils/        # Utility functions (e.g., short URL generator)
│   └── router/       # API route handlers
│       ├── get.go          # Get original URL from short ID
│       ├── shorten.go      # Create a shortened URL
│       ├── editurl.go      # Edit existing short URL
│       ├── tag.go          # Add tags to URLs
│       └── deleteurl.go    # Delete a shortened URL
│
├── DB/
│   ├── redis.Dockerfile       # Dockerfile to run Redis server
│   ├── app.Dockerfile         # Dockerfile to build the Go application
│   └── docker-compose.yml     # Compose file to run app + Redis together
│
├── main.go        # Entry point of the application
└── README.md
```

---

## ✨ Features

* ✂️ **Shorten URLs** – Convert long URLs to short, unique IDs.
* 🔍 **Get Original URL** – Retrieve long URLs from their short versions.
* 📝 **Edit URLs** – Update existing short URL records.
* 🏷️ **Tag URLs** – Add tags to categorize your links.
* ❌ **Delete URLs** – Remove shortened URLs from the database.

---

## 📡 API Endpoints

| Method   | Endpoint                | Description                       |
| -------- | ----------------------- | --------------------------------- |
| `GET`    | `/api/v1/:shortid`      | Fetch original URL using short ID |
| `POST`   | `/api/v1/shorten`       | Shorten a new URL                 |
| `PUT`    | `/api/v1/editurl`       | Edit an existing short URL        |
| `POST`   | `/api/v1/tag`           | Add tag(s) to a shortened URL     |
| `DELETE` | `/api/v1/deleteurl/:id` | Delete a shortened URL            |

> Note: Prefix all routes with `/api/v1/` for versioning support.

---

## 🐳 Dockerized Setup

1. **Redis Dockerfile (`redis.Dockerfile`)**
   Sets up a Redis container for storing shortened URLs.

2. **App Dockerfile (`app.Dockerfile`)**
   Builds the Go application into a Docker container.

3. **Docker Compose (`docker-compose.yml`)**
   Runs both the app and Redis container, allowing them to communicate.

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

### 2. Start Using Docker Compose

```bash
docker-compose up --build
```

> This builds both Redis and the Go app and starts them in connected containers.

---

## 🧰 Tech Stack

* 🧠 **Go (Golang)** – High-performance backend logic
* 🌐 **Gin** – Lightweight HTTP web framework
* ⚡ **Redis** – Fast in-memory key-value store
* 📦 **Docker & Docker Compose** – For easy deployment and orchestration

---

## 🔮 Future Enhancements

* 📈 URL click analytics and tracking
* 🔒 Authentication & rate-limiting
* ⏳ Expiry for temporary links
* 🌐 Support for custom aliases and domains

---

## 🤝 Contributing

Contributions are always welcome! Feel free to fork the project, open issues, or submit pull requests.

---

## 📄 License

This project is licensed under the **MIT License**.

---

