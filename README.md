
# Microservices with NGINX API Gateway and Rate Limiting

This project demonstrates a simple microservices architecture using Golang for backend services, Docker for containerization, and NGINX as an API gateway with rate limiting and load balancing.

## Directory Structure

- **`gateway/`**: Contains NGINX configuration (`nginx.conf`) and Dockerfile for the API gateway.
- **`internal/users/`**: Golang microservice for managing user data.
- **`internal/orders/`**: Golang microservice for managing orders data.

## Features

- **API Gateway**: NGINX configured as an API gateway to route requests to appropriate microservices.
- **Load Balancing**: NGINX configured with load balancing to distribute traffic among multiple instances of backend services.
- **Rate Limiting**: Configured rate limiting to protect backend services from excessive requests.

## Setup

### Prerequisites

- Docker installed on your machine.
- Ensure port `8080` is not in use.


### Steps to Run

1. **Clone Repository:**

   ```bash
   git clone https://github.com/PGabriel20/api-gateway.git
   ```
1. **Build and Run Services:**

   ```bash
   cd api-gateway
   docker compose up -d
   ```
