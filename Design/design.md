# Project Description

## Client Description
Develop a web-based task management application where users can:
- Register and log in.
- Create tasks and assign them to other users.
- Mark tasks as complete.
- Utilize both front-end and back-end components.

## Developer Description
- **Architectural Paradigm:** Microservices with Clean Architecture
- **Design Pattern:** Domain-Driven Design (DDD)
- **API Type:** RESTful HTTP API

## Conceptual Description
The task management app will support multiple admins, each managing their own set of users. Admins can:
- Create, assign, and update tasks for their users.
- View task statuses and overall progress.
- Receive notifications when task deadlines approach.

Users can:
- View tasks assigned to them.
- Update task statuses.
- Receive notifications when assigned a task.

Admins will also have access to task analytics, including:
- Number of completed tasks.
- Number of in-progress tasks.
- Number of uncompleted tasks.
- Number of users managed.
- Task statuses displayed in a table.

## Problem Modeling Framework
- **Identity:** Defines user behavior and roles.
- **Domain Entities:** Represents atomic objects of the app.
- **Capabilities:** Outlines permissions for each identity.
- **Microservices:** Establishes well-defined, separate concerns.
- **Values:** Features provided to each user role.

### Identities
- **Super-Admins:** Owners of the application who register admins.
- **Admins:** Manage tasks and oversee specific groups of users.
- **Users:** Execute assigned tasks under an admin’s supervision.

### Values
#### Admins
- Create and assign tasks to users.
- View task performance and analytics.
- Receive email/alert notifications for incomplete tasks.

#### Users
- View assigned tasks.
- Update task statuses.
- Receive email notifications when assigned a task.

## Domain Entities
```json
User {
  "user_id": "string",
  "owner_id": "string",
  "first_name": "string",
  "last_name": "string",
  "email": "string"
}

Task {
  "task_id": "string",
  "user_id": "string",
  "owner_id": "string",
  "title": "string",
  "description": "string",
  "status": "string",
  "deadline": "string"
}

Admin {
  "admin_id": "string",
  "owner_id": "string",
  "first_name": "string",
  "last_name": "string",
  "email": "string"
}

SuperAdmin {
  "super_admin_id": "string",
  "first_name": "string",
  "last_name": "string",
  "email": "string"
}
```

## Capabilities
### Super-Admin
- Register admins.
- View all admins.
- Delete an admin.

### Admin
- Register users.
- View all users under their management.
- Delete users.
- Create, read, update, and delete tasks.
- Assign tasks to users.
- View task analytics.

### User
- View assigned tasks.
- Update task statuses.
- View task analytics.

## API Endpoints
### Authentication Service
| Method | Endpoint | Description |
|--------|---------|-------------|
| POST | `/auth_service/auth/register` | Register super-admin/admin |
| POST | `/auth_service/auth/login` | Login for all users |
| POST | `/auth_service/auth/verify` | Verify authentication |

### User Management Service
| Method | Endpoint | Description |
|--------|---------|-------------|
| POST | `/user_management_service/admin/register` | Create an admin |
| GET | `/user_management_service/admin/read` | View all admins |
| DELETE | `/user_management_service/admin/remove/{id}` | Delete an admin |
| POST | `/user_management_service/user/register` | Register a user |
| GET | `/user_management_service/user/read_all` | View all users under admin |
| DELETE | `/user_management_service/user/remove/{id}` | Delete a user |

### Task Management Service
| Method | Endpoint | Description |
|--------|---------|-------------|
| POST | `/task_management_service/task/write` | Create a task |
| GET | `/task_management_service/task/read_all` | View all tasks under admin |
| GET | `/task_management_service/task/remove/{id}` | Remove a specific task |
| DELETE | `/task_management_service/task/remove_all` | Remove all tasks under an admin |
| GET | `/task_management_service/task/analytics/read` | View task analytics |
| POST | `/task_management_service/task/assign` | Assign a task to a user |
| GET | `/task_management_service/task/read_all` | View tasks assigned to a user |
| PUT | `/task_management_service/task/update_status` | Update task status |
| GET | `/task_management_service/task/analytics/read` | View user task analytics |

## Microservices Architecture
### Authentication & Authorization Service
- **Database:** PostgreSQL
- **Transport:** HTTP
- **Programming Language:** Go
- **Web Framework:** Fiber or Go-Chi
- **Monitoring:** Grafana & Prometheus
- **Containerization:** Docker

### User Management Service
- **Database:** PostgreSQL
- **Transport:** HTTP
- **Programming Language:** Go
- **Web Framework:** Fiber or Go-Chi
- **Containerization:** Docker

### Task Management Service
- **Database:** MongoDB
- **Transport:** HTTP
- **Programming Language:** Go/JavaScript
- **Web Framework:** Fiber, Go-Chi, or Node.js
- **Containerization:** Docker

### UI
- **Transport:** HTTP
- **Programming Language:** JavaScript
- **Framework & Tools:** React.js, HTML5, Tailwind CSS
- **Containerization:** Docker
- **Multi-Container:** Docker-Compose

