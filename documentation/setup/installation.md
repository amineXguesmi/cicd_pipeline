# Installation

This guide provides step-by-step instructions to install and run the application both locally and in a Dockerized environment.

---

## Local Installation

To set up the application on your local machine, follow these steps:

### 1. Clone the Repository
First, clone the project repository to your local system:
```bash
git clone https://github.com/amineXguesmi/cicd_pipeline.git

cd cicd_pipeline
```

### 2. Install Dependencies
This project uses **Go Modules** for dependency management. Ensure you have Go installed (v1.22 or higher).

To download and install the required dependencies:
```bash
go mod tidy
```

### 3. Set Up Environment Variables
Create a `.env` file in the project root directory and configure the required environment variables:
```
MONGODB_URI=your_mongodb_uri
JWT_SECRET=your_jwt_secret
DB_NAME=your_database_name
```

### 4. Run the Application Locally
Start the Go application:
```bash
go run main.go
```

The application will be accessible at:
```
http://localhost:8080
```

### 5. (Optional) Run Tests Locally
To verify that everything works correctly, run the unit tests:
```bash
go test -v ./...
```

---

## Dockerized Installation

For a production-ready setup or easier deployment, you can run the application in a Dockerized environment.

### 1. Build and Start the Application
Navigate to the project root directory and run the following command:
```bash
docker-compose up --build
```

This command will:
- Build the Docker image for the Go application.
- Start the **Go app** and any additional services (e.g., MongoDB) as defined in the `docker-compose.yml`.

### 2. Access the Application
Once the containers are running, you can access the application at:
```
http://localhost:8080
```

### 3. Stopping the Application
To stop the running containers, execute:
```bash
docker-compose down
```

This will:
- Stop the running containers.
- Preserve database data in Docker volumes (if configured).

---

## Notes on Dockerized Installation

- **Persistent Data**:
  - If the application uses a database, ensure volumes are configured in the `docker-compose.yml` to persist data across restarts.

- **Configuration Files**:
  - Modify the `docker-compose.yml` file to adjust environment variables, ports, or service configurations.

- **Ports**:
  - The default application port is `8080`.
  - Ensure the port is not being used by other services.

- **Logs**:
  - View application logs using:
    ```bash
    docker-compose logs -f
    ```

---