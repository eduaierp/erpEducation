Perfect, Tushar — I’ve got you 💯. Here's a **complete `README.md`** that includes:

- ✅ Project Overview  
- ✅ PostgreSQL Setup (Mac & Windows)  
- ✅ Database Schema + Required Queries  
- ✅ Project Setup & Dependencies  
- ✅ How to Run  
- ✅ How to Test via Postman (gRPC)

---

### 📄 `README.md`

```markdown
# 📝 Attendance Service (gRPC - Golang)

This is a gRPC microservice built with Golang for managing student attendance records.

---

## 📦 Tech Stack

- Language: **Go (Golang)**
- Protocol: **gRPC**
- Database: **PostgreSQL**
- Communication: **Protobuf**
- Dev Tooling: **Postman (gRPC tab)** or **BloomRPC**

---

## 🖥️ Supported gRPC APIs

| Method            | Description                        |
|------------------|------------------------------------|
| MarkAttendance    | Add new attendance for a student   |
| GetAttendance     | Get a specific attendance record   |
| ListAttendance    | List all attendance of a student   |

---

## 🐘 PostgreSQL Setup

### 🔧 Installation

#### On macOS:
```bash
brew install postgresql
brew services start postgresql
```

#### On Windows:
1. Download & install PostgreSQL from: https://www.postgresql.org/download/windows/
2. Use pgAdmin to manage DB or `psql` CLI.

---

### 📂 Database Creation & Configuration

1. Login to PostgreSQL:
```bash
psql -U postgres
```

2. Create DB and user:
```sql
CREATE DATABASE eduaierp;
CREATE USER ravikigf WITH PASSWORD 'root';
GRANT ALL PRIVILEGES ON DATABASE eduaierp TO ravikigf;
```

3. Connect to database:
```bash
\c eduaierp
```

4. Create attendance table:
```sql
CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    student_id INTEGER NOT NULL,
    date VARCHAR(50),
    status VARCHAR(20) -- Present / Absent / Leave
);
```

---

## 🚀 Project Setup

### 📁 Project Structure

```
attendance-service/
├── go.mod
├── main.go
├── proto/
│   └── attendance.proto
├── model/
│   └── attendance.go
├── handler/
│   └── attendance_handler.go
├── repository/
│   └── attendance_repository.go
└── server/
    └── server.go
```

### 🛠️ Commands

#### 1. Initialize Go module (if not already done):
```bash
go mod init attendance-service
```

#### 2. Install dependencies:
```bash
go mod tidy
```

#### 3. Generate gRPC code:
```bash
protoc --go_out=. --go-grpc_out=. proto/attendance.proto
```

> Ensure you have `protoc` and Go plugins installed.

#### 4. Run the service:
```bash
go run main.go
```

✅ You should see:
```
✅ Attendance Service running on port 50054...
```

---

## 🧪 Test via Postman (gRPC)

### 🧰 Requirements:
- **Postman Desktop App**
- Use **gRPC tab** (not REST)

### 🧬 Steps:
1. Open Postman and create a new **gRPC Request**
2. Connect to server:
   ```
   localhost:50054
   ```
3. Import the file `proto/attendance.proto`
4. Select the service and method (e.g. `AttendanceService/MarkAttendance`)
5. Sample payloads:

#### ➕ MarkAttendance
```json
{
  "student_id": 1,
  "date": "2025-04-05",
  "status": "Present"
}
```

#### 📄 GetAttendance
```json
{
  "id": 1
}
```

#### 📋 ListAttendance
```json
{
  "student_id": 1
}
```

---

## 📌 Notes

- Database connection in `main.go` uses:
  ```go
  connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
  ```

- If you face error like `missing mustEmbedUnimplementedAttendanceServiceServer`, ensure this is included in handler:
  ```go
  pb.UnimplementedAttendanceServiceServer
  ```

---

