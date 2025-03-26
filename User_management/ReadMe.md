## User Management Service

### Key Functionality
- Enables admin, super-admin, and user to log in to the app.
- Enables an admin/super-admin to sign up a user or admin.
- Enables an admin to read a user by `user_id`.
- Enables an admin to remove a user by `user_id`.
- Enables an admin/super-admin to remove all of their users or admins, respectively.
- Enables an admin/super-admin to read all of their users or admins.
- Enables an admin/user/super-admin to be authenticated.

### Domain Entity
```go
User {
   Owner_id     string
   Name         string
   UserName     string
   Password     string
   Email        string
   CreatedAt    time.Time 
   UpdatedAt    time.Time
}
```

### API Specifications

1. **User Authentication Verification**  
   - **Method:** GET  
   - **URL:** `/task_app/user_manager_service/api/v0.1/user/verify`
   - **Description:** Enables authentication/verification for all users.  
   - **Headers:**  
     ```json
     {
       "Content-Type": "application/json",
       "Authorization": "Bearer Token"
     }
     ```
   - **Request Body:** `{}`
   - **Response:** `{ user_profile }`

2. **User Login**  
   - **Method:** GET  
   - **URL:** `/task_app/user_manager_service/api/v0.1/user/login`
   - **Description:** Enables users to log in to the app.
   - **Headers:**  
     ```json
     {
       "Content-Type": "application/json",
       "Authorization": "Bearer Token"
     }
     ```

3. **User/Admin Signup**  
   - **Method:** POST  
   - **URL:** `/task_app/user_manager_service/api/v0.1/user/signup`
   - **Description:** Enables an admin/super-admin to register users or admins.
   - **Headers:**  
     ```json
     {
       "Content-Type": "application/json",
       "Authorization": "Bearer Token"
     }
     ```

4. **Retrieve User by ID**  
   - **Method:** GET  
   - **URL:** `/task_app/user_manager_service/api/v0.1/user`
   - **Description:** Enables an admin or super-admin to retrieve one user or admin using `user_id`.
   - **Headers:**  
     ```json
     {
       "Content-Type": "application/json",
       "Authorization": "Bearer Token"
     }
     ```
   - **Request Parameter:** `user_id`
   - **Response:** `{ user }`

5. **Remove User by ID**  
   - **Method:** DELETE  
   - **URL:** `/task_app/user_manager_service/api/v0.1/user`
   - **Description:** Enables an admin or super-admin to remove a user or admin using `user_id`.
   - **Headers:**  
     ```json
     {
       "Content-Type": "application/json",
       "Authorization": "Bearer Token"
     }
     ```
   - **Request Parameter:** `user_id`
   - **Response:** `{ user_id }`

6. **Remove All Users/Admins**  
   - **Method:** DELETE  
   - **URL:** `/task_app/user_manager_service/api/v0.1/user`
   - **Description:** Enables an admin or super-admin to remove all users or admins under their management.
   - **Headers:**  
     ```json
     {
       "Content-Type": "application/json",
       "Authorization": "Bearer Token"
     }
     ```
   - **Request Parameter:** `user_id`
   - **Response:** `{ list_of_users }`

7. **Notify User of Task Assignment**  
   - **Method:** POST  
   - **URL:** `/task_app/user_manager_service/api/v0.1/user/notify`
   - **Description:** Enables admins to notify users when they are assigned a task.
   - **Headers:**  
     ```json
     {
       "Content-Type": "application/json",
       "Authorization": "Bearer Token"
     }
     ```
   - **Request Body:**  
     ```json
     {
       "email": "string",
       "title": "string",
       "description": "string",
       "deadline": "string"
     }
     ```
   - **Response:**  
     ```json
     {
       "status": 200,
       "message": "User has been notified successfully about the task assignment"
     }
     ```

### Policy and Authorization Definitions

| Role          | Endpoint                                          | Method |
|--------------|--------------------------------------------------|--------|
| admin        | `/task_app/user_manager_service/api/v0.1/user`   | POST   |
| super-admin  | `/task_app/user_manager_service/api/v0.1/user`   | POST   |
| admin/user/super-admin | `/task_app/user_manager_service/api/v0.1/user` | GET   |
| super-admin  | `/task_app/user_manager_service/api/v0.1/user`   | GET    |
| admin        | `/task_app/user_manager_service/api/v0.1/user`   | DELETE |
| super-admin  | `/task_app/user_manager_service/api/v0.1/user`   | DELETE |
