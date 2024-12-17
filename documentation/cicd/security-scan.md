### **Security Scan**

#### **Trivy**
- **Trivy** is a powerful open-source tool for detecting vulnerabilities in your codebase and container images.
- It scans the project for critical and high-severity vulnerabilities that pose potential risks.
- If unresolved vulnerabilities are found during the scan, the pipeline will fail, ensuring that only secure code and images are deployed.

#### **Snyk**
- **Snyk** helps in identifying security vulnerabilities in your code, open-source dependencies, and Docker images.
- It continuously monitors for known vulnerabilities and alerts on critical issues that may affect the application.
- The pipeline will continue even if vulnerabilities are detected, with the results logged for review. Snyk ensures that you can track and manage vulnerabilities during the development process.
