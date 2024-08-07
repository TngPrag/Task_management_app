# Multi-tenant and Microservice Oriented Task Management Application

## Key Outcomes

- Implement the app using Microservices architectural paradigm
- Practice Data-first and multitenancy design patterns
- Make the app implementation language-agnostic and let the business contract determine the language

## Architectural Paradigm and Descriptions

- **Microservice Architecture** with a Data-first and domain-driven design pattern
- **Multi-tenant Architecture** with clear RBAC (Role-Based Access Control) policy definitions to enable multiple user groups to manage their tasks and users
- Emphasizes Data-first design pattern which places data at the center of contextually bounded services with clear business contracts. The logic and its transport mechanisms are diffused to the boundary.

### Key Values of Data-first Design

- Build scalable products
- Leverage data as a currency
- Develop products with artificial intelligence in mind to keep pace with today's data-driven products
- Facilitate easy maintenance of the software

- The app includes a microservice development framework that decouples the persistence layer/file-system layer from the data, business logic, and transport mechanisms. This boosts productivity across teams and allows developers to focus on the business contract.
- Compatible with Go and Node.js and tested with various databases like TimescaleDB, PostgreSQL, MongoDB, and InfluxDB
- Microservices APIs use HTTP protocol for transportation, but for services requiring event handling and processing, inter-service communication using NATS message broker or Kafka for reliability is recommended.

## Project Description

- A task management app enabling a super admin to create RBAC policies for different admins
- Admins, created by the super admin with certain privileges, can create their own users and assign tasks
- Admins can track the status of tasks with historical trends in their insight dashboard
- Users can view tasks assigned to them by their admin with clear deadlines and receive email notifications upon task assignment or deadline updates. They can also track their historical trends in their insight dashboard.
- Reflects a multi-tenant philosophy/pattern where the super admin owns admins, and admins own their tasks and users, enabling CRUD operations on their tasks and users with clear separation of concern.

## Services Description

### Git Description

1. **Step 1:** Create a project directory for the task app
2. **Step 2:** Clone the repository
    ```sh
    git clone https://github.com/TngPrag/tele_mini_project.git
    ```
3. **Step 3:** Navigate to the project directory
    ```sh
    cd tele_mini_project/
    ```

### Authz Service

- **Key Features:**
  - Create policies on user and admin app usage
  - Define roles for users and admins
  - Single source of truth for resource usage
- **Technologies:**
  - Language: Go
  - Framework: Fiber
  - Database: PostgreSQL
  - Authorization Library: Casbin
- **Setup:**
  ```sh
  cd Auth
  Install PostgreSQL.
  Create a database task_app with user and password postgres.
  Install dependencies: go mod tidy
  Start the server: go run main.go
  ```
- **Monitoring:**
  - Metrics: [http://localhost:8980/metrics](http://localhost:8980/metrics)
- **Documentation:**
  - Swagger Docs: [http://localhost:8980/swagger](http://localhost:8980/swagger)

### User Manager Service

- **Key Features:**
  - User and admin login
  - Signup, read, and remove users and admins
  - Authentication
- **Technologies:**
  - Language: Go
  - Framework: Fiber
  - Database: PostgreSQL
- **Setup:**
  ```sh
  cd User_management
  Install PostgreSQL.
  Create a database task_app with user and password postgres.
  Install dependencies: go mod tidy
  Start the server: go run main.go
  ```
- **Monitoring:**
  - Metrics: [http://localhost:8981/metrics](http://localhost:8981/metrics)
- **Documentation:**
  - Swagger Docs: [http://localhost:8981/swagger](http://localhost:8981/swagger)

### Task Manager Service

- **Key Features:**
  - CRUD operations on Task object
  - Task assignment and email notifications
  - Visualize task status on a dashboard
- **Technologies:**
  - Language: JavaScript
  - Framework: Express
  - Database: MongoDB
- **Setup:**
  ```sh
  Start the server: node main.js
  ```
- **Monitoring:**
  - Status: [http://localhost:8903/status](http://localhost:8903/status)
- **Documentation:**
  - Swagger Docs: [http://localhost:8903/api-docs](http://localhost:8903/api-docs)

### UI Service

- **Key Features:**
  - Login page for admin and user
  - Separate dashboards for admin and user
- **Technologies:**
  - Language: JavaScript
  - Framework: React.js
  - Styling: Tailwind CSS
- **Setup:**
  ```sh
  cd UI
  Install dependencies: npm install
  Start the server: npm start
  Access: http://localhost:3000/login
  ```

## Deployment Architecture

- **Containerization:** Each microservice is containerized using Docker for OS platform agnosticism
- **Orchestration:** Uses Docker Compose for simplicity but supports Kubernetes for orchestration
- **Reverse Proxy and Load Balancing:** Uses Nginx
- **Network and Volumes:** Each service container has dedicated networks and volumes for data persistence

### Build Steps

1. **Step 1:** Navigate to Docker directory
    ```sh
    cd Docker
    ```
2. **Step 2:** Build services
    ```sh
    sudo docker-compose build
    ```
3. **Step 3:** Start containers
    ```sh
    sudo docker-compose up --remove-orphans
    ```

## Contact

For more information or inquiries, please contact [tng.nat2023@gmail.com](mailto:tng.nat2023@gmail.com).
