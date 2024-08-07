Multi-tenant and Microservice oriented Task Management Application
==================================================================
0. Key Outcomes
   -> To implement this app from Microservices architectural paradignm
   -> To practice Data first and multitenancy design patterns.
   -> To make the app implementation language aggnostic and let the business contract determine the language. 
1. Architectural Paradigm and descriptions
------------------------------------------
-> Microservice Architecture with a Data first and domain driven design pattern.
-> A multi-tenant architecture with clear RBAC policy definitions to enable mutliple user groups to manage their tasks and users.
-> Amplifies Data first design pattern which magnifies the data to be placed at the center of contextually bounded services having clear business contracts, and the logic and its transport mechanisms will be diffused to the boundary. The key value of making data at the center of a lightweight microservices is 
      * To build sacalable products.
      * To build products that leverages data as a currency
      * TO build products with artificial intelligence in mind and get in pace with today's 
        data driven products.
      * To enable maintenance of the software easy.
-> The app have a microservice development framework which eanbles the percistence layer/ file-system layer to be completley decoupled from the the data, buisness logic, its transport mechanisms this enables to boost productivity across teams and  the developer to focus on the its business contract. so the frame-work works well with Go and node.js with the above architectural paradignms and design patterns, its tested across different databases like timescale db, postgres database, mongodb, influxdb...

-> The mircoservices API's uses HTTP protocol as transportation mechanism but if the service demands to handle and process events it should use an inter-service communication specially if scalabilty and reliability demanded within the product roadmap, and the framework could work smoothly with NATS message broker and if we need reliability with kafka .

2. Project Description
 -> This is task management app that enables a super admin to create RBAC policies for different admins then each admins created by the super-admin with certain previlages can create their own users and assign tasks.
 -> Admins could create their own users could track the status of tasks with historical trends in their insight dashboard.
 -> Users could also view tasks assigned to them by their admin with clear deadlines and but also they are notified to their email upon the task assignment or deadline schedule updates. They could also track their historical trends in their inight dashboard.
 -> The multi-tenant philosophy/pattern is reflected within this app, as super-admin owns admins, and admins owns their tasks and users, so they can perform CRUD operations on their tasks and usrs with clear separation of concern.


>> Services Description
  Git Description:
  step-1: create a project directory for task app
  step-2: git clone https://github.com/TngPrag/tele_mini_project.git
  step-3: cd tele_mini_project/
  1. Authz Service

    > Key Features:
        * Create policies on user and admin app usage.
        * Define roles for users and admins.
        * Single source of truth for resource usage.
    > Technologies:
        * Language: Go
        * Framework: Fiber
        * Database: PostgreSQL
        * Authorization Library: Casbin
    > Setup:
        * cd Auth
        * Install PostgreSQL.
        * Create a database task_app with user and password postgres.
        * Install dependencies: go mod tidy
        * Start the server: go run main.go
    > Monitoring:
        * Metrics: http://localhost:8980/metrics
    > Documentation
        * Swagger Docs: http://localhost:8980/swagger

2. User Manager Service

    > Key Features:
        * User and admin login.
        * Signup, read, and remove users and admins.
        * Authentication.
    > Technologies:
        * Language: Go
        * Framework: Fiber
        * Database: PostgreSQL
    > Setup:
        * cd User_management
        * Install PostgreSQL.
        * Create a database task_app with user and password postgres.
        * Install dependencies: go mod tidy
        * Start the server: go run main.go
    > Monitoring:
        * Metrics: http://localhost:8981/metrics
    > Documentation
        * Swagger Docs: http://localhost:8981/swagger

  3. Task Manager Service

    > Key Features:
        * CRUD operations on Task object.
        * Task assignment and email notifications.
        * Visualize task status on a dashboard.
    > Technologies:
        * Language: JavaScript
        * Framework: Express
        * Database: MongoDB
    > Setup:
       *  Start the server: node main.js
    > Monitoring:
        Status: http://localhost:8903/status
    > Documentation
       * Swagger Docs: http://localhost:8903/api-docs

  4. UI Service

    > Key Features:
        * Login page for admin and user.
        * Separate dashboards for admin and user.
    > Technologies:
        * Language: JavaScript
        * Framework: React.js
        * Styling: Tailwind CSS
    > Setup:
        * cd UI
        * Install dependencies: npm install
        * Start the server: npm start
        * Access: http://localhost:3000/login

>> Deployment Architecture

    > Containerization: Each microservice is containerized using Docker for OS platform agnosticism.
    > Orchestration: Uses Docker Compose for simplicity but supports Kubernetes for orchestration.
    > Reverse Proxy and Load Balancing: Uses Nginx.
    > Network and Volumes: Each service container has dedicated networks and volumes for data persistence.
    > Build Steps:
        * cd Docker
        * Build services: sudo docker-compose build
        * Start containers: sudo docker-compose up --remove-orphans

Contact

For more information or inquiries, please contact tng.nat2023@gmail.com.





    
    
     
    
        
            





