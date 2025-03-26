# Task Manager Service

## Domain Entity

### Task
A task is a logical entity that contains the user responsible for implementing the task and the admin who owns the task. It also includes attributes such as status, deadline, title, and description.

#### Task Object Structure:
```json
{
    "Id": "string",
    "Title": "string",
    "Description": "string",
    "Status": "string",
    "Deadline": "string"
}
```

## Key Functionalities
- Enables Admin to perform CRUD operations on Task objects.
- Allows Admin to assign tasks to users; users receive email notifications upon task assignment.
- Enables users to view their tasks and update task statuses.
- Provides Admin with a dashboard visualization for:
  - Number of pending tasks
  - Number of tasks in progress
  - Number of completed tasks

## API Specifications

### 1. Create a Task (Admin)
**Method:** `POST`  
**URL:** `/task_app/task_manager_service/api/v0.1/task/write`  
**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Request Body:**
```json
{
    "Title": "string",
    "Description": "string",
    "Status": "string",
    "Deadline": "string"
}
```
**Response:**
```json
{
    "status": "Task created successfully!"
}
```

### 2. Read a Task (Admin/User)
**Method:** `GET`  
**URL:** `/task_app/task_manager_service/api/v0.1/task/read/:id`  
**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Response:**
```json
{
    "Id": "string",
    "Title": "string",
    "Description": "string",
    "Status": "string",
    "Deadline": "string"
}
```

### 3. Update Task Status (Admin)
**Method:** `PUT`  
**URL:** `/task_app/task_manager_service/api/v0.1/task/update_status/:id`  
**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Request Body:**
```json
{
    "status": "Pending | In-progress | Completed"
}
```
**Response:**
```json
{
    "UpdatedTask": "Task updated successfully"
}
```

### 4. Update Task Schedule (Admin)
**Method:** `PUT`  
**URL:** `/task_app/task_manager_service/api/v0.1/task/update_schedule/:id`  
**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Request Body:**
```json
{
    "schedule": "Pending | In-progress | Completed"
}
```
**Response:**
```json
{
    "UpdatedTask": "Task schedule updated successfully"
}
```

### 5. Remove a Task (Admin)
**Method:** `DELETE`  
**URL:** `/task_app/task_manager_service/api/v0.1/task/remove/:id`  
**Description:** Allows an Admin to remove a task identified by `task_id`.

**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Response:**
```json
{
    "status": "Task identified by {task_id} removed successfully"
}
```

### 6. Remove All Tasks Assigned to a User (Admin)
**Method:** `DELETE`  
**URL:** `/task_app/task_manager_service/api/v0.1/tasks/remove_by_user`  
**Description:** Allows an Admin to remove all tasks assigned to a specific user identified by `user_id`.

**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Response:**
```json
{
    "status": "All tasks assigned to user identified by {user_id} removed successfully"
}
```

### 7. Remove All Tasks Created by Admin
**Method:** `DELETE`  
**URL:** `/task_app/task_manager_service/api/v0.1/tasks/remove_by_owner`  
**Description:** Allows an Admin to remove all tasks they created.

**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Response:**
```json
{
    "status": "All tasks created by admin removed successfully"
}
```

### 8. List All Tasks Created by Admin
**Method:** `GET`  
**URL:** `/task_app/task_manager_service/api/v0.1/tasks/list_by_admin`  
**Description:** Allows an Admin to list all tasks they created.

**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Response:**
```json
{
    "tasks": [
        { "Id": "string", "Title": "string", "Description": "string", "Status": "string", "Deadline": "string" }
    ]
}
```

### 9. List All Tasks Assigned to a User
**Method:** `GET`  
**URL:** `/task_app/task_manager_service/api/v0.1/tasks/list_by_user`  
**Description:** Allows a user to list all tasks assigned by their Admin.

**Headers:**
```json
{
    "Content-Type": "application/json",
    "Authorization": "Bearer <Token>"
}
```
**Response:**
```json
{
    "tasks": [
        { "Id": "string", "Title": "string", "Description": "string", "Status": "string", "Deadline": "string" }
    ]
}
```

## Policy and Authorization Definition
| Role  | Endpoint | Method |
|-------|------------------------------------------------------------|--------|
| Admin | `/task_app/task_manager_service/api/v0.1/task` | `POST` |
| Admin | `/task_app/task_manager_service/api/v0.1/task` | `GET` |
| User  | `/task_app/task_manager_service/api/v0.1/task` | `GET` |
| Admin | `/task_app/task_manager_service/api/v0.1/task` | `PUT` |
| Admin | `/task_app/task_manager_service/api/v0.1/task` | `DELETE` |

---


