# Prerequisites

To get started with this project, ensure you have the following installed and configured:

---

## 1. **Docker**
   - Install Docker to containerize and run services:
     - [Get Docker](https://www.docker.com/get-started)

---

## 2. **Docker Compose**
   - Install Docker Compose to orchestrate multi-container setups:
     - [Install Docker Compose](https://docs.docker.com/compose/install/)

---

## 3. **SonarQube Server**

This project integrates **SonarQube** for code quality analysis. You can deploy it locally or on a remote server. Below are the deployment options:

### Option 1: Deploying SonarQube on a Remote VM (e.g., Azure)

#### Why?
Hosting SonarQube on a cloud VM ensures it's always accessible for your CI/CD pipeline.

#### Steps:
1. **Create a virtual machine (VM)**:
   - Use Azure (or any other cloud provider) to create a VM with:
     - A public IP address.
     - Minimum 4 GB of RAM and sufficient disk space (SonarQube requires resources).
2. **Access the VM**:
   - SSH into your VM:
     ```bash
     ssh <username>@<vm-ip-address>
     ```
3. **Install Docker and Docker Compose** on the VM.

4. **Deploy SonarQube**:
   - Use the provided `docker-compose-sonarQube.yml` file to deploy SonarQube:
     ```bash
     docker-compose -f docker-compose-sonarQube.yml up -d
     ```

5. **Expose Port 6500**:
   - For Azure, go to the VM's **Networking settings** and add an **inbound rule** for port 6500.

6. **Access SonarQube**:
   - Visit SonarQube in your browser at:
     ```
     http://<vm-ip-address>:6500
     ```

7. **Generate a Token for CI/CD Integration**:
   - Log into SonarQube at `http://<vm-ip-address>:6500`.
   - Navigate to:
     **My Account** > **Security** > **Tokens**.
   - Generate a new token and **copy it** (you’ll need this for GitHub Actions).

---

### Option 2: Deploying Locally with Ngrok for CI/CD Integration

#### What is Ngrok?
**Ngrok** is a tool that creates a secure tunnel to expose your local services to the internet. It’s a quick way to make your local SonarQube instance accessible to your CI/CD pipeline.

#### Steps to Deploy SonarQube Locally with Ngrok:
1. **Run SonarQube Locally**:
   - Use the provided `docker-compose-sonarQube.yml` file:
     ```bash
     docker-compose -f docker-compose-sonarQube.yml up -d
     ```
   - Access SonarQube at `http://localhost:6500`.

2. **Install Ngrok**:
   - Download and install **Ngrok**: [Get Ngrok](https://ngrok.com/download).

3. **Expose SonarQube**:
   - Use Ngrok to expose your local SonarQube instance:
     ```bash
     ngrok http 6500
     ```
   - Copy the generated **public URL** (e.g., `https://randomstring.ngrok.io`).

4. **Use the Ngrok URL in CI/CD**:
   - Log into SonarQube using the Ngrok URL (e.g., `https://randomstring.ngrok.io`).
   - Generate a token as described above.
   - Use this **public URL** in your CI/CD configuration.

---

## 4. **Configure GitHub Actions to Use SonarQube**

In your GitHub Actions workflow, include the following secrets to connect to SonarQube:

### Environment Variables in GitHub Actions
- **`SONAR_HOST_URL`**: This is the URL of your SonarQube instance (e.g., `http://<vm-ip-address>:6500` or `https://randomstring.ngrok.io`).
- **`SONAR_TOKEN`**: The token you generated in SonarQube.

#### Steps to Add Secrets in GitHub:
1. Go to your GitHub repository.
2. Navigate to:
   **Settings** > **Secrets and Variables** > **Actions** > **New Repository Secret**.
3. Add the following secrets:
   - **`SONAR_ORGANIZATION`**: URL of your SonarQube instance.
   - **`SONAR_PROJECT_KEY`**: The unique identifier for your project in SonarQube
   - **`SONAR_TOKEN`**: The authentication token generated in SonarQube.


---

## Why the SonarQube URL is Needed

The SonarQube server must be accessible to your **GitHub Actions** CI/CD pipeline. Whether hosted on a remote VM or exposed locally via Ngrok, the server URL ensures that the `sonarsource/sonarqube-scan-action` can connect to your SonarQube instance to analyze code quality.