Multi-tenant and Microservice oriented Task Management Application
--------------------------------------------------------------------
Key Outcomes
    Microservices Architecture: Implement the app using a microservices architectural paradigm.
    Data-First and Multitenancy Design Patterns: Practice these design patterns.
    Language Agnostic Implementation: Ensure the business contract determines the language, making the implementation language agnostic.

Architectural Paradigm and Descriptions
Microservice Architecture

    Data-First and Domain-Driven Design Pattern: Focus on placing data at the center of contextually bounded services with clear business contracts.
    Multi-Tenant Architecture: Implement a multi-tenant architecture with clear RBAC policy definitions to enable multiple user groups to manage their tasks and users.
    Benefits:
        Scalable products.
        Leverage data as a currency.
        AI-ready products.
        Easier software maintenance.

Framework and Design Patterns

    Microservice Development Framework: Decouple the persistence layer from the data, business logic, and transport mechanisms.
    Compatibility: Works well with Go and Node.js and tested across databases like TimescaleDB, PostgreSQL, MongoDB, InfluxDB.
    Inter-Service Communication: Uses HTTP protocol but supports NATS message broker and Kafka for scalability and reliability.

Project Description

    Task Management App: Enables a super admin to create RBAC policies for different admins. Admins can create their own users and assign tasks.
    Admin Features: Create and manage users, track task status, and view historical trends on their dashboard.
    User Features: View assigned tasks, receive email notifications, and track historical trends on their dashboard.
    Multi-Tenant Philosophy: Super-admin owns admins, and admins own their tasks and users, ensuring clear separation of concern.

Services Description
Authz Service

    Key Features:
        Create policies on user and admin app usage.
        Define roles for users and admins.
        Single source of truth for resource usage.
    Technologies:
        Language: Go
        Framework: Fiber
        Database: PostgreSQL
        Authorization Library: Casbin
    Setup:
        cd Auth
        Install PostgreSQL.
        Create a database task_app with user and password postgres.
        Install dependencies: go mod tidy
        Start the server: go run main.go
    Monitoring:
        Metrics: http://localhost:8980/metrics
        Swagger Docs: http://localhost:8980/swagger

User Manager Service

    Key Features:
        User and admin login.
        Signup, read, and remove users and admins.
        Authentication.
    Technologies:
        Language: Go
        Framework: Fiber
        Database: PostgreSQL
    Setup:
        cd User_management
        Install PostgreSQL.
        Create a database task_app with user and password postgres.
        Install dependencies: go mod tidy
        Start the server: go run main.go
    Monitoring:
        Metrics: http://localhost:8981/metrics
        Swagger Docs: http://localhost:8981/swagger

Task Manager Service

    Key Features:
        CRUD operations on Task object.
        Task assignment and email notifications.
        Visualize task status on a dashboard.
    Technologies:
        Language: JavaScript
        Framework: Express
        Database: MongoDB
    Setup:
        Start the server: node main.js
    Monitoring:
        Status: http://localhost:8903/status
        Swagger Docs: http://localhost:8903/api-docs

UI Service

    Key Features:
        Login page for admin and user.
        Separate dashboards for admin and user.
    Technologies:
        Language: JavaScript
        Framework: React.js
        Styling: Tailwind CSS
    Setup:
        cd UI
        Install dependencies: npm install
        Start the server: npm start
        Access: http://localhost:3000/login

Deployment Architecture

    Containerization: Each microservice is containerized using Docker for OS platform agnosticism.
    Orchestration: Uses Docker Compose for simplicity but supports Kubernetes for orchestration.
    Reverse Proxy and Load Balancing: Uses Nginx.
    Network and Volumes: Each service container has dedicated networks and volumes for data persistence.
    Build Steps:
        cd Docker
        Build services: sudo docker-compose build
        Start containers: sudo docker-compose up --remove-orphans

License

Include the license information here.
Contact

For more information or inquiries, please contact tng.nat2023@gmail.com.
