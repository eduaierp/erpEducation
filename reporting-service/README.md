Here’s a clean and complete `README.md` for your **reporting-service**, modeled after the style of your attendance-service:

---

### 📄 `README.md`

```markdown
# 📊 Reporting Service (ERP Module)

This is the `reporting-service` module of the ERP system. It handles generation, retrieval, and listing of system reports via gRPC.

---

## 🚀 Features

- Generate reports with title, content, and timestamp.
- Retrieve report by ID.
- List all generated reports.

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Protocol**: gRPC
- **Protobuf**: Protocol Buffers v3

---

## 📦 Project Structure

```
reporting-service/
├── go.mod
├── main.go
├── handler/
│   └── report_handler.go
├── model/
│   └── report.go
├── proto/
│   └── reporting.proto
├── repository/
│   └── report_repository.go
└── server/
    └── server.go
```

---

## 🧑‍💻 Local Development Setup

### 1. Clone the repository
```bash
git clone https://github.com/your-org/reporting-service.git
cd reporting-service
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Generate gRPC code from proto
```bash
protoc --go_out=. --go-grpc_out=. proto/reporting.proto
```

### 4. Setup PostgreSQL
Ensure your database `erp` is created with the following table:
```sql
CREATE TABLE reports (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TEXT NOT NULL
);
```

### 5. Run the service
```bash
go run main.go
```
> Service will start on `localhost:50060`

---

## 🧪 gRPC Testing with Postman

### Example 1: GenerateReport
```json
{
  "title": "Monthly Finance Summary",
  "content": "All transactions are accounted for.",
  "created_at": "2025-04-06"
}
```

### Example 2: GetReport
```json
{
  "id": 1
}
```

### Example 3: ListReports
```json
{}
```

> You can test using [Postman](https://www.postman.com/downloads) or [Evans CLI](https://github.com/ktr0731/evans)

---

## 📫 gRPC Endpoints

| RPC Method         | Request Type             | Response Type         | Description                      |
|--------------------|--------------------------|------------------------|----------------------------------|
| `GenerateReport`   | `GenerateReportRequest`  | `ReportResponse`       | Create a new report              |
| `GetReport`        | `GetReportRequest`       | `ReportResponse`       | Get a report by its ID           |
| `ListReports`      | `ListReportsRequest`     | `ListReportsResponse`  | Retrieve all available reports   |

---

