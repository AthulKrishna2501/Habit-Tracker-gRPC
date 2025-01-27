# Habit Tracker gRPC Microservices

This project demonstrates a Habit Tracker application implemented using gRPC microservices in Golang. The system consists of two services: the **User Service** and the **Habit Service**, both interacting with a PostgreSQL database. The application is designed to help users manage and track their daily habits effectively.

## Features

### User Service
- Create a new user
- Retrieve user details

### Habit Service
- Create a new habit for a user
- Mark habits as completed
- Retrieve habits for a specific user

## Architecture

The project follows a microservices architecture, with each service having its own responsibilities and database interactions. Communication between services is handled via gRPC.

### Project Structure
```
├── user-service
│   ├── db
│   │   ├── user-repository.go
│   ├── proto
│   │   └── user-service.proto
│   ├── main.go
│   ├── server
│       └── user-server.go
├── habit-service
│   ├── db
│   │   ├── habit-repository.go
│   ├── proto
│   │   └── habit-service.proto
│   ├── main.go
│   ├── server
│       └── habit-server.go
├── README.md
```

## Prerequisites

- Go 1.20+
- PostgreSQL
- Protocol Buffers Compiler (`protoc`)
- gRPC plugins for Go (`protoc-gen-go` and `protoc-gen-go-grpc`)
- BloomRPC (optional, for testing gRPC services)

## Setup

### Clone the Repository
```bash
git clone https://github.com/your-repo/habit-tracker-grpc.git
cd habit-tracker-grpc
```

### Install Dependencies
```bash
go mod tidy
```

### Generate gRPC Code
Navigate to each service directory (e.g., `user-service/proto` or `habit-service/proto`) and run:
```bash
protoc --go_out=. --go-grpc_out=. service-name.proto
```

Replace `service-name.proto` with the corresponding `.proto` file.

### Set Up the Database
Create PostgreSQL databases for both services:
```sql
CREATE DATABASE user_service_db;
CREATE DATABASE habit_service_db;
```

Apply the necessary migrations. Ensure the database credentials are updated in the respective service's `main.go` file.

### Run the Services
Start the **User Service**:
```bash
cd user-service
go run main.go
```

Start the **Habit Service**:
```bash
cd habit-service
go run main.go
```

## Usage

### Testing with BloomRPC
1. Open BloomRPC and load the `.proto` files for each service.
2. Use the gRPC endpoints to test the functionality:
   - `CreateUser`
   - `GetUser`
   - `CreateHabit`
   - `MarkHabitComplete`
   - `GetHabits`

### Sample gRPC Requests

#### Create User
```json
{
  "name": "Athul",
  "email": "athul@example.com"
}
```

#### Get User
```json
{
  "user_id": "<user-id>"
}
```

#### Create Habit
```json
{
  "user_id": "<user-id>",
  "title": "Morning Exercise",
  "description": "30 minutes of running"
}
```

#### Mark Habit as Complete
```json
{
  "habit_id": "<habit-id>",
  "completion_date": "2025-01-27"
}
```

## Troubleshooting

### Common Issues

#### Error: `protoc-gen-go: program not found or is not executable`
Ensure `protoc-gen-go` and `protoc-gen-go-grpc` are installed and available in your PATH:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### Error: `duplicate key value violates unique constraint`
Check if the email is already registered in the database. Delete the existing entry or use a different email.

## Future Improvements
- Implement authentication and authorization.
- Add unit tests and integration tests.
- Use Docker for containerization.
- Deploy using Kubernetes.

## License
This project is licensed under the MIT License.

## Author
Athul

