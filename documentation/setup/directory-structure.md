# Directory Structure

This document outlines the structure of the project and provides a description of each directory and file.

---

## Project Root Structure

```
AUTH-MS/
├── .github/
│   └── workflows/
│       └── cicd_pipeline.yaml   # GitHub Actions workflow for CI/CD pipeline
├── config/
│   └── config.go                # Configuration file for environment variables
├── documentation/               # Project documentation
├── handlers/
│   ├── auth.go                  # Authentication-related handlers
│   └── health.go                # Health check endpoint
├── models/
│   ├── jwt-token.go             # JWT token model and logic
│   └── user.go                  # User model
├── utils/
│   ├── db.go                    # Database connection utility
│   └── generate-jwt-token.go    # JWT token generation utility
├── .env                         # Environment variables file
├── .gitignore                   # Specifies files to ignore in version control
├── Dockerfile                   # Docker configuration for building the application
├── go.mod                       # Go module dependencies
├── go.sum                       # Checksums for Go module dependencies
├── main.go                      # Application entry point
└── unit_test.go                 # Unit tests for the project                    
```

---

## Key Directories and Files

### 1. **Project Root**
- **`.github/`**: Contains GitHub Actions workflows for CI/CD.
  - **`cicd_pipeline.yaml`**: Defines the CI/CD pipeline for automated testing and deployment.
- **`config/`**: Contains environment-specific configuration files.
  - **`config.go`**: Manages application configuration variables.
- **`documentation/`**: Project documentation for setup and usage.

### 2. **Handlers (`handlers/`)**
- **`auth.go`**: Handles authentication-related logic, including login and token generation.
- **`health.go`**: Defines the health check endpoint to monitor the service status.

### 3. **Models (`models/`)**
- **`jwt-token.go`**: Implements the logic related to generating and validating JWT tokens.
- **`user.go`**: Defines the user model, including attributes and database interactions.

### 4. **Utilities (`utils/`)**
- **`db.go`**: Provides functions for connecting to the database.
- **`generate-jwt-token.go`**: Contains functions for generating JWT tokens used for authentication.

### 5. **Project Configuration**
- **`.env`**: Holds environment variables for configuration and secrets.
- **`.gitignore`**: Specifies files and directories to exclude from version control, ensuring sensitive or unnecessary files are not pushed to repositories.
- **`Dockerfile`**: Defines how to build the application container.
- **`go.mod`**: Lists Go module dependencies.
- **`go.sum`**: Contains checksums for Go module dependencies.
- **`main.go`**: The entry point for the application, starting the server and routing requests.
- **`unit_test.go`**: Contains unit tests for the application to ensure correct functionality.