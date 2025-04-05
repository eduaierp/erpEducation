

---

## 📘 Auth Service – README

This is a gRPC-based authentication microservice built in Go. It handles user login and JWT generation.

---

### 🛠 Tech Stack

- **Go (Golang)**
- **gRPC**
- **PostgreSQL**
- **bcrypt for password hashing**
- **JWT for token generation**

---

## 📦 Project Structure

```
auth-service/
├── go.mod
├── main.go
├── handler/
│   └── auth_handler.go
├── model/
│   └── user.go
├── proto/
│   └── auth.proto
├── repository/
│   └── user_repository.go
├── server/
│   └── server.go
├── utils/
│   └── jwt.go
└── generate_hash.go
```

---

## 🚀 How to Run

### 1. 🧾 Prerequisites

- Go 1.20+ installed
- PostgreSQL running
- Your DB has a `users` table like this:

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  role TEXT NOT NULL
);
```

### 2. 🧪 Insert Sample User

To insert a user:

1. Run the hash generator:

```bash
go run generate_hash.go
```

2. Copy the output `Hashed password`.

3. Insert into PostgreSQL:

```sql
INSERT INTO users (username, password, role)
VALUES ('admin', 'your_hashed_password_here', 'admin');
```

Example:

```sql
INSERT INTO users (username, password, role)
VALUES ('admin', '$2a$10$G0c6YMre44E8jDOHjWV9outcPMwsu5TLIo15gTY.1cXFAA/lrMb46', 'admin');
```

### 3. 🧩 Install Dependencies

```bash
go mod tidy
```

### 4. ▶️ Run the Service

```bash
go run main.go
```

Expected output:

```
✅ Auth Service running on port 50061...
```

---

## 📬 How to Test with Postman

### 1. Install Postman gRPC client (native app required)

### 2. Add the `.proto` file

- Import `auth.proto` from the `proto/` directory into Postman.

### 3. Connect to gRPC server

- **gRPC server:** `localhost:50061`
- **Method:** `AuthService/Login`

### 4. Send the Request

In Postman, go to **Body**, choose `JSON`, and send:

```json
{
  "username": "admin",
  "password": "admin123"
}
```

> Make sure the password matches the one you hashed!

### 5. Expected Response

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "role": "admin"
}
```

---

## 🔐 Token Structure

The JWT contains:

```json
{
  "username": "admin",
  "role": "admin",
  "exp": 1712398897
}
```

It is signed using HS256 and the secret key `supersecretkey`.

---

## 📂 generate_hash.go

You can use this file to hash new passwords manually:

```bash
go run generate_hash.go
```

It prompts you for a password and prints the hash.

