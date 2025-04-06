Here’s a complete `README.md` file for your `finance-service`, following the same format as your `attendance-service`:

---

```md
# 📊 Finance Service (gRPC) - eduErp

The **Finance Service** is a microservice in the eduErp ecosystem responsible for managing financial transactions such as credits and debits. It provides full gRPC-based APIs for creating, retrieving, and listing transactions.

---

## 🚀 Features

- Create new financial transactions
- Fetch transaction by ID
- List all transactions

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **RPC Framework**: gRPC
- **Database**: PostgreSQL
- **Protobuf**: For service definitions

---

## 📁 Project Structure

```
finance-service/
├── go.mod
├── handler/                # gRPC method implementations
│   └── finance_handler.go
├── main.go                 # Entry point
├── model/                  # Data models
│   └── transaction.go
├── proto/                  # Protobuf definitions
│   └── finance.proto
├── repository/             # DB operations
│   └── finance_repository.go
└── server/                 # gRPC server setup
    └── server.go
```

---

## ⚙️ Setup & Run Locally

### ✅ Prerequisites

- Go 1.20+
- PostgreSQL running locally
- Protobuf compiler + `protoc-gen-go` plugin
- Postman (v10+) or `grpcurl` for testing

### 📦 Install Dependencies

```bash
cd finance-service
go mod tidy
```

### 🗃️ Create Table

Ensure the following table exists in your PostgreSQL `erp` database:

```sql
CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  type TEXT,
  amount REAL,
  date TEXT,
  description TEXT
);
```

### 🧠 Run the Service

```bash
go run main.go
```

Service runs on: `localhost:50055`

---

## 🧪 Testing with Postman (gRPC)

### 1. Open Postman > New > gRPC Request  
### 2. Import `proto/finance.proto`  
### 3. Connect to: `localhost:50055`  
### 4. Call these methods:

#### ➕ CreateTransaction

```json
{
  "type": "Credit",
  "amount": 250.75,
  "date": "2025-04-06",
  "description": "Initial funding"
}
```

#### 🔍 GetTransaction

```json
{
  "id": 1
}
```

#### 📋 ListTransactions

```json
{}
```

---

## 📘 Protobuf Reference

```proto
service FinanceService {
  rpc CreateTransaction (CreateTransactionRequest) returns (TransactionResponse);
  rpc GetTransaction (GetTransactionRequest) returns (TransactionResponse);
  rpc ListTransactions (ListTransactionsRequest) returns (ListTransactionsResponse);
}
```

---

