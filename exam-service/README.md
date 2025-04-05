Here's a **detailed `README.md`** for your `exam-service`, modeled after the structure and completeness of your `attendance-service` docs:

---

```markdown
# 📘 Exam Service – eduErp

The **Exam Service** is a gRPC-based microservice in the `eduErp` system, responsible for managing exam-related data such as exam creation, retrieval, and listing. Built with **Go (Golang)**, it uses **PostgreSQL** for data storage and is structured for maintainability, scalability, and clarity.

---

## 📁 Project Structure

```
exam-service/
├── go.mod
├── main.go                     # Entry point
├── handler/
│   └── exam_handler.go         # gRPC service logic
├── model/
│   └── exam.go                 # Exam data model
├── proto/
│   └── exam.proto              # gRPC service definitions
├── repository/
│   └── exam_repository.go      # DB queries
└── server/
    └── server.go               # gRPC server bootstrap
```

---

## 🧠 Features

- Create new exams
- Retrieve a single exam by ID
- List all exams
- Clean modular structure
- Secure PostgreSQL integration

---

## 🧱 Tech Stack

- **Language**: Go 1.21+
- **Protocol**: gRPC
- **Database**: PostgreSQL
- **Dependencies**:
  - `google.golang.org/grpc`
  - `github.com/lib/pq`

---

## 🔌 gRPC Methods

### ✅ CreateExam
- **Request**:
```json
{
  "subject": "Mathematics",
  "date": "2025-04-10",
  "type": "Final"
}
```
- **Response**:
```json
{
  "exam": {
    "id": 1,
    "subject": "Mathematics",
    "date": "2025-04-10",
    "type": "Final"
  }
}
```

---

### ✅ GetExam
- **Request**:
```json
{
  "id": 1
}
```
- **Response**:
```json
{
  "exam": {
    "id": 1,
    "subject": "Mathematics",
    "date": "2025-04-10",
    "type": "Final"
  }
}
```

---

### ✅ ListExams
- **Request**:
```json
{}
```
- **Response**:
```json
{
  "exams": [
    {
      "id": 1,
      "subject": "Mathematics",
      "date": "2025-04-10",
      "type": "Final"
    },
    {
      "id": 2,
      "subject": "Science",
      "date": "2025-04-11",
      "type": "Midterm"
    }
  ]
}
```

---

## 🗄️ Database Schema

```sql
CREATE TABLE IF NOT EXISTS exams (
    id SERIAL PRIMARY KEY,
    subject VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    type VARCHAR(100) NOT NULL
);
```

Ensure the connected user has the correct privileges:

```sql
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE exams TO ravikigf;
GRANT USAGE, SELECT, UPDATE ON SEQUENCE exams_id_seq TO ravikigf;
```

---

## 🚀 Running the Service

### 1. **Start PostgreSQL**
Make sure your `eduaierp` database is up and running.

### 2. **Update DB connection (if needed)**
Edit `main.go`:
```go
connStr := "user=ravikigf dbname=eduaierp sslmode=disable password=root"
```

### 3. **Run the gRPC server**
```bash
go run main.go
```

Service starts on:
```
localhost:50057
```

---

## 🧪 gRPC Testing via Postman

1. Open Postman → Create new **gRPC Request**
2. Enter:
   ```
   localhost:50057
   ```
3. Import `proto/exam.proto`
4. Choose methods (`CreateExam`, `GetExam`, `ListExams`) and use the sample payloads above.

