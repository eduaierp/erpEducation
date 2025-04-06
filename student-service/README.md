Perfect, Tushar! Here's a **clean and professional `README.md`** for your `student-service`, modeled in the same style as your `attendance-service` with consistent formatting and developer-friendly instructions.

---

### 📄 `README.md`

```markdown
# 🎓 student-service

A gRPC-based microservice to manage student records as part of the eduErp system. This service allows you to Create, Retrieve, Update, Delete, and List student data. Built using Go, gRPC, and PostgreSQL.

---

## 🚀 Features

- Add a new student
- Fetch a student by ID
- Update student details
- Delete a student record
- List all students

---

## 📦 Tech Stack

- **Language:** Go (Golang)
- **RPC Framework:** gRPC with Protocol Buffers
- **Database:** PostgreSQL
- **ORM/Driver:** database/sql with lib/pq

---

## 🛠️ Project Structure

```
student-service/
├── go.mod
├── main.go
├── handler/
│   └── student_handler.go
├── model/
│   └── student.go
├── proto/
│   └── student.proto
├── repository/
│   └── student_repository.go
├── server/
│   └── server.go
```

---

## 🧪 gRPC Endpoints

| Method            | Request Message          | Description            |
|------------------|--------------------------|------------------------|
| CreateStudent     | `CreateStudentRequest`   | Create a new student   |
| GetStudent        | `GetStudentRequest`      | Get student by ID      |
| UpdateStudent     | `UpdateStudentRequest`   | Update student details |
| DeleteStudent     | `DeleteStudentRequest`   | Delete student by ID   |
| ListStudents      | `ListStudentsRequest`    | List all students      |

---

## ⚙️ Setup & Run Locally

### 1. Clone the repo
```bash
git clone https://github.com/yourname/student-service.git
cd student-service
```

### 2. Initialize Go module
```bash
go mod tidy
```

### 3. Create PostgreSQL DB & Table
```sql
CREATE DATABASE erp;

\c erp

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    course VARCHAR(100)
);
```

### 4. Compile `.proto` file
```bash
protoc --go_out=. --go-grpc_out=. proto/student.proto
```

### 5. Run the service
```bash
go run main.go
```

Server will start at `localhost:50051`

---

## 🧪 Testing with Postman (gRPC)

### 1. Open Postman v10+
### 2. Create a new gRPC request
- Import the file `proto/student.proto`
- Connect to `localhost:50051`
- Choose a method like `student.StudentService/CreateStudent`

### Example Payloads:

#### ✅ CreateStudent
```json
{
  "name": "Alice",
  "email": "alice@example.com",
  "course": "Mathematics"
}
```

#### ✅ GetStudent
```json
{
  "id": 1
}
```

#### ✅ UpdateStudent
```json
{
  "id": 1,
  "name": "Alice Smith",
  "email": "alice.smith@example.com",
  "course": "Physics"
}
```

#### ✅ DeleteStudent
```json
{
  "id": 1
}
```

#### ✅ ListStudents
```json
{}
```

