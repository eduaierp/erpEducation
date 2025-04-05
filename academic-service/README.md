Academic Service - eduERP System

The Academic Service is a microservice within the eduERP system responsible for managing academic programs and courses. It exposes gRPC APIs to Create, Read, Update, Delete, and List both Programs and Courses. It is built using Go, PostgreSQL, and Protocol Buffers.

⚖️ Tech Stack

Language: Go (Golang)

API Communication: gRPC

Protobuf: Protocol Buffers (proto3)

Database: PostgreSQL

Port: :50052

📂 Project Structure

academic-service/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point to start the service
├── config/
│   └── config.go                   # Database connection logic
├── handler/
│   └── academic_handler.go         # gRPC service handler implementation
├── model/
│   └── academic.go                 # Models for Program and Course
├── proto/
│   └── academic.proto              # gRPC service and message definitions
├── repository/
│   └── academic_repository.go      # DB query and interaction logic
├── server/
│   └── server.go                   # Starts the gRPC server
├── go.mod
├── go.sum
└── README.md                      # Project documentation

📆 Database Schema

CREATE TABLE programs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    code VARCHAR(50),
    program_id INTEGER REFERENCES programs(id)
);

⚖️ Configuration

Update PostgreSQL credentials in cmd/server/main.go:

connStr := "user=postgres dbname=erp sslmode=disable password=postgres"

🤖 Protobuf Compilation

protoc --go_out=. --go-grpc_out=. proto/academic.proto

This will generate the required gRPC code for Go.

🚀 Running the Service

cd cmd/server
go run main.go

Output:

Academic Service running on port 50052

🔎 gRPC Endpoints (Postman / grpcurl)

Use academic.proto in Postman or with grpcurl to test endpoints.

Base URL

localhost:50052

📖 Program API Endpoints

1. CreateProgram

{
  "name": "Bachelor of Science",
  "description": "3-year degree program"
}

2. GetProgram

{
  "id": 1
}

3. UpdateProgram

{
  "id": 1,
  "name": "BSc Mathematics",
  "description": "Updated description"
}

4. DeleteProgram

{
  "id": 1
}

5. ListPrograms

{}

📗 Course API Endpoints

1. CreateCourse

{
  "title": "Linear Algebra",
  "code": "MATH101",
  "program_id": 1
}

2. GetCourse

{
  "id": 1
}

3. UpdateCourse

{
  "id": 1,
  "title": "Advanced Linear Algebra",
  "code": "MATH201",
  "program_id": 1
}

4. DeleteCourse

{
  "id": 1
}

5. ListCourses

{}

⚠️ Common Errors

Foreign Key Violation: Ensure that program_id exists before creating a course.

DB Connection Error: Check your DB credentials and make sure PostgreSQL is running.

gRPC Not Running: Make sure the server is running on port 50052.





