Task app services
================
1. Authz service
    -language: Go
    -webframework: fiber
    -database: postgresql

 * To start the server
   cmd:    go run main.go
 * To monitor the server
   http://localhost:8980/task_app/authz_service/api/v0.1/health
 * To view swagger docs
   http://localhost:8980/swagger
2. User manager service
     -language: Go
     -webframework: fiber
     -database: postgresql
 * TO start the server
    cmd: go run main.go
 * To monitor the server
   http://localhost:8981/task_app/user_manager_service/api/v0.1/health
 * To view swagger docs
   http://localhost:8981/swagger
3. Task manager service
      -language: Javascript
      -webframework: express
      -database: mongodb
 * To start the server
    cmd: node main.js
* To monitor the server
    http://localhost:3000/status
* To view swagger docs
  http://localhost:3000/swagger


