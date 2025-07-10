# Go Bank Server API

## Overview
The Bank Server API is a secure, high-performance JSON API built using Go, designed for financial applications. It supports money transformation (e.g., currency conversion, transactions), user authentication, and JWT-based access control. Leveraging Go's concurrency and robust standard library, this API provides a scalable solution for fintech use cases.

## Features
- *Backend API*: RESTful JSON API built with Go for high performance and scalability.
- *Money Transformation*: Supports currency conversion or transaction processing with precise calculations.
- *Authentication*: Secure user authentication system with password hashing.
- *JWT Authentication*: Token-based authentication using JSON Web Tokens (JWT) for secure API access.
- *Request Handlers*: Modular HTTP handlers for routing and processing API requests.

## Prerequisites
- Go (version 1.20 or higher)
- A running instance of a database (e.g.PostgreSQL )
- Environment variables configured (see .env.example)

## Installation
1. Clone the repository:
   bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   

2. Install dependencies:
   bash
   go mod tidy
   

3. Set up environment variables:
   bash
   cp .env.example .env
   
   Edit .env to configure database credentials, JWT secret, and other settings.

4. Run the application:
   bash
   go run cmd/main.go
   
   The API will be available at http://localhost:8080 (or the port specified in .env).

## Usage
### Authentication
- *Register*: POST /api/register with JSON payload { "username": "string", "password": "string" }.
- *Login*: POST /api/login with JSON payload { "username": "string", "password": "string" }. Returns a JWT token.
- *Protected Routes*: Include the JWT token in the Authorization header as Bearer <token>.

### Money Transformation
- *Convert Currency*: POST /api/convert with JSON payload { "amount": float64, "from": "USD", "to": "EUR" }.
- *Process Transaction*: POST /api/transaction with JSON payload { "user_id": int, "amount": float64, "currency": "USD" }.

### Example Request
bash
curl -X POST http://localhost:8080/api/login \
-H "Content-Type: application/json" \
-d '{"username": "user", "password": "pass"}'


## API Endpoints
| Method | Endpoint              | Description                     | Authentication |
|--------|-----------------------|---------------------------------|----------------|
| POST   | /api/register       | Register a new user            | None           |
| POST   | /api/login          | Authenticate and get JWT token | None           |
| POST   | /api/convert        | Convert currency               | JWT            |
| POST   | /api/transaction    | Process a transaction          | JWT            |
| GET    | /api/health         | Check API health               | None           |

## JWT Authentication
- Tokens are signed using HS256 algorithm.
- Token expiration is set to 24 hours (configurable in .env).
- Invalid or expired tokens return a 401 Unauthorized response.

## Development
To contribute or modify the project:
1. Fork the repository.
2. Create a feature branch: git checkout -b feature/your-feature.
3. Commit changes: git commit -m "Add your feature".
4. Push to the branch: git push origin feature/your-feature.
5. Open a pull request.

## Testing
Run unit tests:
bash
go test ./...


## Dependencies
- [gorilla/mux](https://github.com/gorilla/mux): HTTP router and dispatcher
- [jwt-go](https://github.com/dgrijalva/jwt-go): JWT implementation
- [godotenv](https://github.com/joho/godotenv): Environment variable loader
- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto): Cryptographic utilities

## Contact
For questions or support, please open an issue or contact the maintainer at [your-email@example.com]