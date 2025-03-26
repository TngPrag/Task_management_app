Authorization Service
======================

### Init PostgreSQL Database
- Initialize the PostgreSQL database.
- Set up authorization policies for super-admin, admin, and user.

### Domain Entities
#### Policy
```
Policy {
    sub:     string
    Object:  string
    Action:  string
}
```
#### Role
```
Role  {
    user_id:     string
    role:       string
}
```

### Key Features
- Enables super-admin to create policies that define how users and admins use the app, including access to specific API endpoints.
- Allows super-admins and admins to define roles for users and other admins.
- Serves as a single source of truth for determining who can access which resources.

### API Specification

#### 1. Create Policy
**Identity:** super-admin  
**Method:** POST  
**URL Route:** `/task_app/authz_service/api/v0.1/policy/write`  
**Description:** Enables super-admin to define a policy for admin and user.  
**Header:**
```
{
    'Content-Type': 'application/json',
    'Authorization': 'Bearer Token'
}
```
**Request Body:**
```
{
    "sub": "string",
    "Object": "string",
    "Action": "string"
}
```
**Response:**
```
{
    "status": 200,
    "message": "Policy successfully created"
}
```

#### 2. Read Policy
**Identity:** super-admin  
**Method:** GET  
**URL Route:** `/task_app/authz_service/api/v0.1/policy/read/:sub`  
**Description:** Enables super-admin to read policies assigned to a specific subject.  
**Header:**
```
{
    'Content-Type': 'application/json',
    'Authorization': 'Bearer Token'
}
```
**Request Parameter:** `sub`  
**Response:**
```
{
    "policies": [ {list of policies for the subject} ]
}
```

#### 3. Remove Policy
**Identity:** super-admin  
**Method:** DELETE  
**URL Route:** `/task_app/authz_service/api/v0.1/policy/remove`  
**Header:**
```
{
    'Content-Type': 'application/json',
    'Authorization': 'Bearer Token'
}
```
**Request Body:**
```
{
    "subject": "string",
    "Object": "string",
    "Action": "string"
}
```
**Response:**
```
{
    "status": "Policy removed successfully"
}
```

#### 4. Check Permission
**Identity:** super-admin/admin/user  
**Method:** GET  
**URL Route:** `/task_app/authz_service/api/v0.1/policy/check_permission`  
**Header:**
```
{
    'Content-Type': 'application/json',
    'Authorization': 'Bearer Token'
}
```
**Request Body:**
```
{
    "subject": "string",
    "Object": "string",
    "Action": "string"
}
```
**Response:**
```
{
    "permission": true | false
}
```

#### 5. List All Policies
**Identity:** super-admin  
**Method:** GET  
**URL Route:** `/task_app/authz_service/api/v0.1/policy`  
**Description:** Enables super-admin to list all policies.  
**Header:**
```
{
    'Content-Type': 'application/json',
    'Authorization': 'Bearer Token'
}
```
**Response:**
```
{
    "policies": [ {list of policies} ]
}
```

### Policy and Authorization Definitions
```
super-admin, task_app/authz_service/api/v0.1/policy, POST
super-admin, task_app/authz_service/api/v0.1/policy, GET
super-admin, task_app/authz_service/api/v0.1/policy, DELETE
admin, task_app/authz_service/api/v0.1/policy, Verify
user, task_app/authz_service/api/v0.1/policy, Verify
super-admin, task_app/authz_service/api/v0.1/policy, Verify
super-admin, task_app/authz_service/api/v0.1/role, POST
super-admin, task_app/authz_service/api/v0.1/role, GET
admin, task_app/authz_service/api/v0.1/role, POST
admin, task_app/authz_service/api/v0.1/role, GET
user, task_app/authz_service/api/v0.1/role, GET
```

