# Microservices Project with Contract Testing using Keploy

## Architecture of the project

<img width="899" alt="Screenshot 2025-02-17 at 3 27 56 AM" src="https://github.com/user-attachments/assets/e9249b9e-9008-4e1e-92bf-ede159ccc490" />

This project consists of four microservices: `user-service`, `order-service`, `payment-service`, and `notification-service`. Each service has its own functionality and interacts with either a PostgreSQL or MongoDB database. Contract testing is implemented using Keploy to ensure the services work as expected.

---

## Table of Contents

1. [Project Structure](#project-structure)
2. [Services Overview](#services-overview)
3. [Prerequisites](#prerequisites)
4. [Setup and Running the Project](#setup-and-running-the-project)
5. [Endpoints](#endpoints)


---

---

## Services Overview

### 1. User-Service

- **Database**: PostgreSQL
- **Endpoints**:
  - `POST /user`: Create a new user in the PostgreSQL database.
  - `GET /user`: Retrieve user data from the PostgreSQL database.

### 2. Order-Service

- **Database**: MongoDB
- **Endpoints**:
  - `POST /order`: Create a new order in the MongoDB database after checking if the order ID exists in the PostgreSQL database.
  - `GET /order`: Retrieve order data from the MongoDB database.

### 3. Payment-Service

- **Functionality**: Updates the payment status of an order in the MongoDB database.
- **Endpoint**:
  - `POST /pay`: Updates the payment status by hitting the `localhost:8081/order/update-payment` endpoint.

### 4. Notification-Service

- **Functionality**: Updates the email status of an order in the MongoDB database.
- **Endpoint**:
  - `POST /pay`: Updates the email status by hitting the `localhost:8081/update-email` endpoint.

---

## Prerequisites

Before running the project, ensure you have the following installed:

- Docker
- Docker Compose
- Keploy (running via Docker, as it doesn't support macOS natively)

---

## Setup and Running the Project

1. **Clone the Repository**:
   ```bash
   git clone github.com/arpit529srivastava/contract_testing_keploy.git
   ```
2. **Start the Databases and Services**:<br>
   Use Docker Compose to start the PostgreSQL and MongoDB databases along with the microservices.
   ```bash
   docker-compose up -d
3. **Verify Services** :<br>
Ensure all services are running by checking their respective endpoints:
```
- User-Service: localhost:8080/user

- Order-Service: localhost:8081/order

- Payment-Service: localhost:8082/pay

- Notification-Service: localhost:8083/pay  
```

4. **Endpoints**: <br>
* User-Service: <br>
```
POST /users: Create a new user.<br>
GET /users: Retrieve user data.
```

* Order-Service: <br>
```
POST /orders: Create a new order.
GET /orders: Retrieve user data.
```

* Payment-Service: <br>
```
POST /pay: Update payment status.
```
* Notification-Service: <br>
```
POST /pay: Update email status.
```

