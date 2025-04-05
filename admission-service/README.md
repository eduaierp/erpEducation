Sure Tushar! Here's your **complete `README.md`** for the `admission-service`, properly formatted and aligned with your existing services like the attendance service.

---

## 🎓 Admission Service - ERP Backend Microservice

The **Admission Service** is a gRPC-based microservice written in Go that handles student admission records as part of the `eduErp` system. It provides APIs to:

- Apply for admission
- Get admission by ID
- List all admissions

---

## 🗂️ Project Structure

```
admission-service/
├── go.mod
├── main.go
├── proto/
│   ├── admission.proto
│   ├── admission.pb.go
│   └── admission_grpc.pb.go
├── model/
│   └── admission.go
├── handler/
│   └── admission_handler.go
├── repository/
│   └── admission_repository.go
└── server/
    └── server.go
```

---

## 🧠 Data Model

### 📄 `admissions` table (PostgreSQL)

```sql
CREATE TABLE admissions (
    id SERIAL PRIMARY KEY,
    student_name TEXT,
    course TEXT,
    status TEXT,
    applied_date TEXT
);
```

### 🔐 Grant access to user

```sql
ALTER TABLE admissions OWNER TO ravikigf;
```

---

## 🚀 Running the Service

### ✅ Prerequisites

- Go installed
- PostgreSQL running
- Required Go tools:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  export PATH="$PATH:$(go env GOPATH)/bin"
  ```

---

### ⚙️ Steps to Run

1. **Generate Proto Files**
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/admission.proto
   ```

2. **Install Go Dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the Server**
   ```bash
   go run main.go
   ```

---

## 🧪 Postman - gRPC Requests

> Make sure you are using **Postman v10+** with gRPC support.

### 🛰️ gRPC Endpoint
```
localhost:50058
```

### 📦 Service Name
```
admission.AdmissionService
```

---

### 1️⃣ ApplyAdmission

- **Method**: `ApplyAdmission`
- **Request Body**:
```json
{
  "student_name": "Alice Sharma",
  "course": "MCA",
  "applied_date": "2025-04-05"
}
```

---

### 2️⃣ GetAdmission

- **Method**: `GetAdmission`
- **Request Body**:
```json
{
  "id": 1
}
```

---

### 3️⃣ ListAdmissions

- **Method**: `ListAdmissions`
- **Request Body**:
```json
{}
```

