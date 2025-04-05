Thanks Tushar — let’s get this perfect.

I'll now:
1. **Align the README** with the same style as your `auth-service`.
2. **Fix Proto generation instructions** with the correct command and structure.
3. **Correct the project structure formatting** in the README to show the tree clearly and cleanly.
4. **Include all the actual Postman instructions** (gRPC tab config).
5. Keep the tone clean and sharp, just how you'd expect in a real team repo.

---

## ✅ Final `README.md` (Fully Correct & Aligned)

```markdown
# 📡 Communication Service

This is the **Communication Microservice** in the eduERP system. It handles internal messaging between users such as notifications and alerts. Built using **Go + gRPC + PostgreSQL**.

---

## 📁 Project Structure

```
communication-service/
├── cmd/
│   └── main.go                   # Entry point
├── internal/
│   ├── handler/
│   │   └── message_handler.go    # gRPC handler logic
│   ├── model/
│   │   └── message.go            # Message entity
│   ├── repository/
│   │   └── message_repository.go # DB interaction
│   └── server/
│       └── grpc_server.go        # gRPC server setup
├── proto/
│   ├── communication.proto         # Proto definition
│   ├── communication.pb.go         # Generated types
│   └── communication_grpc.pb.go    # Generated gRPC code
├── go.mod
└── README.md
```

---

## ⚙️ Tech Stack

- **Language**: Go 1.18+
- **Protocol**: gRPC (Protocol Buffers)
- **Database**: PostgreSQL
- **Transport**: TCP on port `:50059`

---

## 🛠️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/your-org/communication-service.git
cd communication-service
```

### 2. Setup PostgreSQL

Create the database and table:

```sql
CREATE DATABASE erp;

\c erp

CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  recipient VARCHAR(255),
  subject TEXT,
  body TEXT,
  sent_at VARCHAR(100)
);
```

### 3. Update `.env` or connection string (if applicable)

In `cmd/main.go` (or use environment variables securely):

```go
sql.Open("postgres", "user=postgres dbname=erp sslmode=disable password=postgres")
```

---

## 📦 Dependencies

Run:

```bash
go mod tidy
```

---

## 🔧 Proto Generation

Make sure you have the `protoc` compiler and plugins installed:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

Generate the gRPC files:

```bash
protoc --go_out=. --go-grpc_out=. proto/communication.proto
```

> This will generate:
> - `proto/communication.pb.go`
> - `proto/communication_grpc.pb.go`

---

## 🚀 Running the Service

```bash
go run cmd/main.go
```

Expected log:

```
✅ Communication Service running on port 50059...
```

---

## 🧪 Postman - gRPC API Testing

Use [Postman gRPC](https://blog.postman.com/postman-supports-grpc/) to test this service.

### Step-by-step:

1. Open Postman.
2. Go to **New > gRPC Request**.
3. Connect to: `localhost:50059`.
4. Click **"Select Proto file"** and upload `proto/communication.proto`.
5. Choose **Service**: `CommunicationService`.
6. Select the method you want to call.

---

### 📬 RPC: SendMessage

```json
{
  "to": "john@example.com",
  "subject": "Welcome!",
  "body": "Hello John, your account is activated.",
  "sent_at": "2025-04-05T15:04:00Z"
}
```

### 📨 RPC: GetMessage

```json
{
  "id": 1
}
```

### 📃 RPC: ListMessages

```json
{}
```

---

## 📌 Notes

- All `sent_at` values should be in **ISO 8601** string format.
- gRPC runs on **port 50059**.
- Make sure PostgreSQL is running before starting the service.

