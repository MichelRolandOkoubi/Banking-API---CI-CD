# Architecture Overview

## Component Diagram
The system follows a simple microservice architecture.

- **API Layer**: Written in Go, handles REST requests.
- **CI/CD**: GitLab CI handles automated builds and tests.
- **Containerization**: Docker is used for consistent environments.

## Data Flow
1. User requests resource.
2. Go API processes request.
3. Response returned as JSON.
