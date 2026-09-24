# 🔗 High-Performance Go URL Shortener

A lightweight, concurrent, and containerized URL Shortener built with **Go** and optimized using **Docker Multi-Stage Builds**.

---

## 📌 Features

- **Concurrent-Safe In-Memory Store**: Uses `sync.Mutex` to handle concurrent reads and writes safely.
- **Lightweight Standard Library**: Built purely with Go's standard `net/http` package (no heavy external dependencies).
- **SEO-Friendly Redirection**: Responds with `302 Found` (`http.StatusFound`) for effective redirect handling.
- **Ultra-Small Docker Image**: Multi-stage build slashes the final container image size from **~800MB down to ~10MB**.

---

## 🏗️ Architecture & Docker Optimization

This project uses a **Multi-Stage Docker Build** pattern to produce a minimal, highly secure production image:

1. **Build Stage (`golang:1.25-alpine`)**: Compiles the source code (`main.go`) into a static binary executable.
2. **Final Stage (`alpine:latest`)**: Copies *only* the compiled binary from the build stage into a minimal Alpine Linux container.
+------------------------------------+      +-----------------------------------+
|      BUILD STAGE (golang:1.25)     |      |       FINAL STAGE (alpine)        |
| - Full Go Toolchain & Source Code  | ---> | - Zero Go Toolchain / Source Code |
| - Compiles main.go -> binary     |      | - ONLY Executable Binary (~10MB)  |
+------------------------------------+      +-----------------------------------+
## 🚀 Getting Started

### Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- [Docker](https://www.docker.com/)

---

### 🛠️ Running Locally (Without Docker)

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/SuronjitSharma/url-shortener.git](https://github.com/SuronjitSharma/url-shortener.git)
   cd url-shortener
Run the Go server:

Bash
go run main.go
Access the application:
The server runs at http://localhost:8080.

🐳 Running with Docker
Build the Docker image:

Bash
docker build -t url-shortener .
Run the Docker container:

Bash
docker run -p 8080:8080 url-shortener
The app is now running at http://localhost:8080.

📂 Project Structure
Plaintext
.
├── main.go            # Core application logic & HTTP handlers
├── Dockerfile         # Multi-stage Docker build configuration
├── go.mod             # Go module definition
└── README.md          # Project documentation
🤝 Contributing
Contributions, issues, and feature requests are welcome! Feel free to check the issues page.

📜 License
This project is open-source and available under the MIT License.
