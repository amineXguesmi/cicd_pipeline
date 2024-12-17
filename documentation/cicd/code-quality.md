### **Code Quality**

#### **SonarQube**
- Integrates with the pipeline to provide detailed code analysis.
- Connects to a SonarQube server using a generated token and host URL.
- Blocks builds if critical issues are found.

**Environment Variables for SonarQube:**
- `SONAR_TOKEN`: The token for authenticating with SonarQube.
- `SONAR_HOST_URL`: The URL of the SonarQube server.

---

