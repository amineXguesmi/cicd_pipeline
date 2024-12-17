### **Code Quality**

#### **SonarQube**
- Integrates with the pipeline to provide detailed code analysis.
- Connects to a SonarQube server using a generated token and host URL.
- Blocks builds if critical issues are found.

**Environment Variables for SonarQube:**
- **`SONAR_ORGANIZATION`**: URL of your SonarQube instance.
- **`SONAR_PROJECT_KEY`**: The unique identifier for your project in SonarQube
- **`SONAR_TOKEN`**: The authentication token generated in SonarQube.

---

