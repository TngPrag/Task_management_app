User management Service
=======================
Key functionality
  -> Enables admin,super-admin and user to login into the app
  -> Enables an admin/super-admin to signup a user or admin.
  -> Enables an admin to read a user by user_id
  -> Enables an admin to remove a user by user_id
  -> Enables an admin/super-admin to remove all of its users or admins respectively.
  -> Enables an admin/super-admin to read all of its users or admins. 
  -> Enables an admin/user/super-admin to be authenticated.
Domain entity
-------------
User {
       Owner_id     string
       Name         string
       UserName     string
       Password     string
       Email        string
       CreatedAt    time.Time 
       UpdatedAt    time.Time
     }
API specs
---------
1. Identity->User/admin/super-admin: Method:GET url_route: '/task_app/user_manager_service/api/v0.1/user/verify'
   Description: Enables an authentication api/verification to all users
   Header: {
              'control-type':application/json
              'Authorization':Bearer Token
            }
   req body {}
   resp {
          (user_profile)
        }
2. Identity ->User/admin/super-admin: Method:GET url_route: '/task_app/user_manager_service/api/v0.1/user/login'
      Description: Enables user to login to the app
      Header: {
                'control-type':application/json
                'Authorization':Bearer Token
              }
3. Identity -> admin/super-admin: Method: POST url_route: '/task_app/user_manager_service/api/v0.1/user/signup'
      Description: Enables admin/super-admin to register users or admins respectively.
       Header: {
                 'control-type':application/json
                 'Authorization':Bearer Token
               }
4. Identity -> super-admin/admin: Method: GET url_route: '/task_app/user_manager_service/api/v0.1/user'
      Description: Enables an admin or super user to retrieve one user or admin using user_id.
      Header: {
                'control-type':application/json
         	'Authorization':Bearer Token
              }
      req_param: user_id
      resp {
             {user}
           }
5. Identity -> super-admin/admin: Method: DELETE url_route: '/task_app/user_manager_service/api/v0.1/user'
    Descritpion: Enables an admin or super-admin to remove a user or an admin using user_id
     Header: {
              'control-type':application/json
              'Authorization':Bearer TOken
             }
     req_para: user_id
     resp {
            (user_id)
          }
6. Identity -> super-admin/admin: Method:DELETE url_route: '/task_app/user_manager_service/api/v0.1/user'
    Description: Enables to remove all users or admins by admin or super-admin respectively.
    Header: {
               'Control-type':application/json
               'Authorization':Bearer TOken
            } 
    req_param: user_id
    resp {
          {list of users}
         }
7. Identity -> admin: Method: POST url_route: '/task_app/user_manager_service/api/v0.1/user/notify'
     Description: Enables admins to notify the user when the user is assigned a task.
     Header: {

            'control-type': application/json
            'Authorization': Bearer Token
     } 
     req Body {
           "email" : string,
           "title" : string,
           "description": string,
           "deadline": string,
     }
     resp {
      status: 200
         "user has been notified successfully about the task assignment"
     }

Policy and authorization definition
--------/---------------------------
admin,task_app/user_manager_service/api/v0.1/user, POST
super-admin,task_app/user_manager_service/api/v0.1/user,POST 

admin/user/super-admin,task_app/user_manager_service/api/v0.1/user,GET



super-admin,task_app/user_manager_service/api/v0.1/user, GET

admin,task_app/user_manager_service/api/v0.1/user,DELETE

super-admin,task_app/user_manager_service/api/v0.1/user,DELETE







 

