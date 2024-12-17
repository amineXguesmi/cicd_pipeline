### **Build and Push Docker Image**

#### **Build**
- Creates a Docker image for the application after successful linting, testing, and scanning.
- Uses the `Dockerfile` in the project directory.

**Example Command:**
```bash
docker build -t <image-name>:latest .
```

#### **Push**
- Authenticates with DockerHub using GitHub Secrets.
- Pushes the built image to the DockerHub repository.

**Example Push Command:**
```bash
docker push <repository>/<image-name>:latest
```

---

