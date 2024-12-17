### **Pull Request Process**

#### **Submitting a Pull Request**

1. **Ensure Your PR Meets the Following Criteria**:
   - Passes all automated checks (e.g., linting, tests, and pipeline workflows).
   - Addresses an existing issue (if applicable) or includes a clear purpose.
   - Includes a detailed description of the changes, why they are necessary, and their expected impact.

2. **Open a Pull Request**:
   - Go to the repository’s **Pull Requests** tab.
   - Click **New Pull Request**.
   - Select the base branch (e.g., `main`) and compare it with your feature branch.

3. **Review Process**:
   - A project maintainer will review your PR.
   - Feedback or change requests may be added. Please address these promptly.

4. **Approval and Merge**:
   - Once approved, the maintainer will merge your changes into the base branch.

---

### **Code Style and Standards**

- **Linting**: Ensure your code passes all linting checks using `ruff`:
  ```bash
  poetry run ruff .
  ```
- **Formatting**: Use `ruff --fix` to format your code automatically:
  ```bash
  poetry run ruff . --fix
  ```
- **Tests**: Verify all tests pass before submitting your PR:
  ```bash
  poetry run pytest
  ```

---

### **Best Practices**

- **Keep It Focused**: Submit one feature or bug fix per pull request.
- **Write Tests**: Add or update tests for any code you write.
- **Follow Commit Standards**: Use meaningful, concise commit messages.

By following these guidelines, you help maintain a high standard of quality for the project and its contributions. Thank you for your efforts!