# fullstack-auth-app
A full-stack application with Keycloak for authentication, Go for the backend API, and Vue.js served via Nginx for the frontend.

## Overview
This project sets up a development environment for a system consisting of a Keycloak authentication server, PostgreSQL database, backend, and frontend services. Docker Compose is used to orchestrate the services.

## Prerequisites
- Docker
- Docker Compose

## Project Setup

### 1. Running the Docker Compose
To start the services (Keycloak, PostgreSQL, Backend, and Frontend), run the following command:

```bash
docker-compose up --build
```

## Access the Keycloak Admin Console
Once the services are running, you can access the Keycloak Admin Console by navigating to:
http://localhost:8080

Admin Username: admin

Admin Password: admin
