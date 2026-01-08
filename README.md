# HEX-HRMS

HEX-HRMS is a Human Resource Management System built with Go, following the **Hexagonal Architecture** (Ports and Adapters) pattern to ensure a clean separation of concerns, high testability, and flexibility in changing external dependencies.

## 🏗 Architecture

The project is structured into three main layers:

1.  **Core**: Contains the business logic, domain models, and port definitions (contracts). It is independent of any infrastructure or external frameworks.
    *   `core/models`: Domain entities (User, Department, etc.).
    *   `core/contracts`: Interfaces (Ports) for repositories and use cases.
    *   `core/usecases`: Implementation of business rules.
2.  **Infrastructure (Adapters)**: Implements the contracts defined in the core.
    *   `infra/api`: REST API adapter using the [Gin](https://github.com/gin-gonic/gin) framework.
    *   `infra/repository`: Data persistence adapter using [GORM](https://gorm.io/) and [PostgreSQL](https://www.postgresql.org/).
3.  **CMD**: Entry points for the application.
    *   `infra/api/main.go`: Main entry point for the REST API.

## 🛠 Tech Stack

*   **Language:** [Go](https://go.dev/) (v1.23+)
*   **Web Framework:** [Gin Gonic](https://gin-gonic.com/)
*   **ORM:** [GORM](https://gorm.io/)
*   **Database:** [PostgreSQL](https://www.postgresql.org/)
*   **API Documentation:** (Add tool if applicable, e.g., Swagger/OpenAPI)

## 🚀 Getting Started

### Prerequisites

*   Go 1.23 or higher
*   Docker and Docker Compose (for the database)

### Installation

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    cd HEX-HRMS
    ```

2.  Set up environment variables:
    ```bash
    cp sample.env .env
    # Edit .env with your local configuration
    ```

3.  Start the database using Docker:
    ```bash
    docker build -t hrms-db -f db.Dockerfile .
    # Or use docker-compose if available
    ```

4.  Install dependencies:
    ```bash
    go mod download
    ```

5.  Run the application:
    ```bash
    go run infra/api/main.go
    ```

## 🛣 API Endpoints

### Health Check
*   `GET /health`: Check if the service is running.

### User Management
*   `GET /api/users`: List all users (supports filtering).
*   `POST /api/users`: Create a new user.
*   (Add more as they are implemented)

## 📁 Project Structure

```text
.
├── boundaries/      # Concept boundaries (if applicable)
├── cmd/             # CLI and other entry points
├── core/            # Domain logic and ports
│   ├── contracts/   # Interface definitions
│   ├── models/      # Domain entities
│   └── usecases/    # Business logic implementation
├── docs/            # Documentation
├── infra/           # Infrastructure adapters
│   ├── api/         # REST API implementation (Gin)
│   └── repository/  # Persistence layer (GORM/Postgres)
├── go.work          # Go workspace configuration
└── README.md        # This file
```

## 🤝 Contributing

1.  Fork the project.
2.  Create your feature branch (`git checkout -b feature/AmazingFeature`).
3.  Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4.  Push to the branch (`git checkout -b feature/AmazingFeature`).
5.  Open a Pull Request.

## 📄 License

(Specify license here, e.g., MIT)
