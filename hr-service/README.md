Here’s a clean and professional `README.md` file for your **`hr-service`**, modeled on the `attendance-service` style:

---

### 📄 `README.md`

```md
# 👨‍💼 HR Service - ERP System

The **HR Service** is a gRPC microservice responsible for managing employee records as part of the eduErp backend system.

---

## 🚀 Features

- Add a new employee
- Get employee by ID
- List all employees

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** gRPC
- **Database:** PostgreSQL
- **Proto:** Protocol Buffers v3

---

## 📁 Project Structure

```
hr-service/
├── go.mod
├── main.go
├── handler/
│   └── hr_handler.go
├── model/
│   └── employee.go
├── proto/
│   └── hr.proto
├── repository/
│   └── hr_repository.go
├── server/
│   └── server.go
└── README.md
```

---

## 🧪 Proto Services

```proto
service HRService {
  rpc AddEmployee (AddEmployeeRequest) returns (EmployeeResponse);
  rpc GetEmployee (GetEmployeeRequest) returns (EmployeeResponse);
  rpc ListEmployees (ListEmployeesRequest) returns (ListEmployeesResponse);
}
```

---

## ⚙️ Setup & Run Instructions

### 1. 🐘 Create Database & Table

Make sure PostgreSQL is running:

```bash
createdb erp
psql erp
```

Run:

```sql
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    department TEXT NOT NULL,
    role TEXT NOT NULL,
    doj TEXT NOT NULL
);
```

---

### 2. 🧰 Install Dependencies

```bash
go mod tidy
```

---

### 3. 🔄 Generate gRPC Code

```bash
protoc --go_out=. --go-grpc_out=. proto/hr.proto
```

---

### 4. ▶️ Run the Service

```bash
go run main.go
```

Runs on: `localhost:50056`

---

## 🧪 Test with Postman

1. Open Postman > New > gRPC Request  
2. Connect to: `localhost:50056`  
3. Import the `proto/hr.proto` file  
4. Choose a method and provide request payload  
5. Click "Invoke"

### Example: AddEmployee
```json
{
  "name": "Alice Johnson",
  "department": "HR",
  "role": "Recruiter",
  "doj": "2024-03-15"
}
```

---

## 📬 gRPC Endpoints

| Method         | Request Type           | Description              |
|----------------|------------------------|--------------------------|
| AddEmployee    | `AddEmployeeRequest`   | Add a new employee       |
| GetEmployee    | `GetEmployeeRequest`   | Get employee by ID       |
| ListEmployees  | `ListEmployeesRequest` | List all employees       |

---

## 🔐 Security & Notes

- For now, no auth added (handled by `auth-service` in main system)
- SSL disabled for local PostgreSQL
- Not containerized (can be added later)

