# go-grpc-template

A minimal Go microservice template using **gRPC**, designed with **Hexagonal Architecture (Ports and Adapters)** principles.

---

## 🧱 Project Architecture: Hexagonal Style

This project follows the **Hexagonal Architecture (aka Ports and Adapters)**, which helps in separating business logic from infrastructure and delivery mechanisms.

### 💡 What is Hexagonal Architecture?

Also known as **Ports and Adapters**, this architecture splits your app into core logic (inside the hexagon) and various adapters (outside the hexagon). This leads to:

- internal/port/: Defines interfaces (ports) for business logic.
- internal/service/: Implements business logic (adapters to ports).
- internal/handler/: gRPC handler, adapts transport to port interface.

---

### 📦 Folder Structure

```plaintext
go-grpc-template/
├── cmd/
│   └── server/            # Entry point (main.go)
├── config/
│   └── env.go             # Configuration loader
├── internal/
│   ├── handler/           # gRPC handler, adapts transport to port interface
│   ├── port/              # Defines interfaces (ports) for business logic
│   └── service/           # Implements business logic (adapters to ports)
├── proto/
│   └── user.proto         # gRPC contracts
├── bin/                   # Compiled binary
├── Dockerfile             # Docker setup
├── Makefile               # Build/run automation
└── go.mod
```
---
### 📦 Calling Server via Grpcurl list, describe by Enabling Reflection in main.go of Server

<img width="1484" height="527" alt="image" src="https://github.com/user-attachments/assets/bfa8aaad-c540-40f2-9fc9-aa98a3ca4b81" />
