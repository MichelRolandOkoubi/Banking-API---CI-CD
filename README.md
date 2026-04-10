# Banking API - CI/CD Demo

## Description
Simple Banking API with automated CI/CD pipeline demonstrating DevOps best practices.

## Features
- ✅ RESTful API (Go)
- ✅ Docker containerization
- ✅ Automated testing
- ✅ CI/CD pipeline (GitLab CI)
- ✅ Health check endpoint
- ✅ Production-ready deployment

## Tech Stack
- **Backend:** Go 1.21
- **Containerization:** Docker, Docker Compose
- **CI/CD:** GitLab CI / GitHub Actions
- **Testing:** Automated health checks

## Quick Start

### Local Development
```bash
# Clone repository
git clone https://github.com/votre-username/banking-api-cicd.git
cd banking-api-cicd

# Run with Docker Compose
docker-compose up -d

# Test API
curl http://localhost:8080/api/balance
curl http://localhost:8080/health