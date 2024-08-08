# Account Project

## Overview

This project consists of two microservices: `account-manager` and `payment-manager`, implemented in Go and managed with Docker.

- **account-manager**: Handles account-related operations.
- **payment-manager**: Manages payment transactions.

## Setup and Running

### Prerequisites

- Docker
- Docker Compose

### Running the Services

1. **Build and Start Services**

   In the project root directory, run:

   ```sh
   docker-compose up --build
This command builds the Docker images and starts the services defined in docker-compose.yml.

2. Accessing the Services

- account-manager: Accessible at http://localhost:8080
- payment-manager: Accessible at http://localhost:8081
**Endpoints**

`account-manager` Endpoints
- POST /users/login: User login.

- POST /users/register: User registration.

- POST /accounts: Create a new account (Requires authentication).

- GET /accounts: List all accounts (Requires authentication).

- GET /accounts/:accountId/transactions: Retrieve transactions for a specific account (Requires authentication).

`payment-manager` Endpoints

- POST /send: Send a payment (Requires authentication).
- POST /withdraw: Withdraw funds (Requires authentication).

**Configuration**

Create a `.env` file in the root directory with the following content:
```
DB_HOST=host
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=database
DB_SSLMODE=disable
```
