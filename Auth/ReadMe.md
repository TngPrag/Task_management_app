-> Init postgresql database
-> Init authorization policies for super-admin, admin and user

Authorization service
======================
Domain Entities
-----------------
 Policy {
    sub:     string
    Object:  string
    Action:  string
 }
 Role  {
    use_id:     string
    role:       string
 }

Key features
-------------
   -> Enables super admin to create policy on how users or admins uses the app and specifically each api endpoints.
   -> Enables admins or superadmins to define roles to users or admins respectively.
   -> So generally its used as a single source of truth for how and who should use the resources of the app.
API specification
-----------------
1. Identity: super-admin  Method: POST  url_route: '/task_app/authz_service/api/v0.1/policy/write'
   Description: Enables super-admin to define a policy for admin and user.
   Header: {
          'content-type': application/json
          'Authoriation': Bearer TOken
   } 
   req body {
        sub       string
        Object    string
        Action    string
   }
   resp {
       status 200: "Policy successfully created"
   }
2. Identity: super-admin  Method: GET  url_route: '/task_app/authz_service/api/v0.1/policy/read:sub'
    Description: Enables super-admin to read policies defined to specific subject
    Header: {
             'content-type': application/json
             'Authorization': Bearer Token
    }
    req param: sub
    resp {
        {list of policies for subject}
    }
3. Identity: super-admin Method: DELETE   url_route: '/task_app/authz_service/api/v0.1/policy/remove'
       Header: {
             'content-type': application/json
             'Authorization': Bearer Token
       }
       req_body {
            subject     string
            Object      string
            Action      string
       }
       resp {
            status: policy removed succeffully
       }
4. Identity: super-admin/admin/user  Method: GET   url_route: '/task_app/authz_service/api/v0.1/policy/check_permission'
           Header: {
             'content-type': application/json
             'Authorization': Bearer Token
       }
            req_body {
                subject     string
                Object      string
                Action      string
            }
            resp {
                True|False
            }
5. Identity: super-admin   Method: GET url_route: '/task_app/authz_service/api/v0.1/policy'
       DEscription: Enabes super-admin to list all policies
       Header: {
             'content-type': application/json
             'Authorization': Bearer Token
       }
       resp {
          {list of policies}
       }

Policy and authorication Definition
-----------------------------------
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

