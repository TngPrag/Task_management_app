Task Manager Service 
====================
domain entity
------------
task: is a logical entity which contains a user who implements the task and 
      the one who owns the task which is the admin, but it also have 
      status, Deadline, title and description
    Task {
            Id           string,
            Title        string,
            Description  string,
            Status       string,
            Deadline     string,
            
         }

Key functionalities
   * Enables Admin to perform CRUD operations on Task object
   * Enables Admin to assign task to its users, then eventually users could be notified about the task assignment to their email.
   * Enables users to view their tasks and update the status of the tasks.
   * It also enables Admin to visualize task status on a dashboard like
        - number of tasks pending
        - number of tasks In-progress
        - number of tasks completed
  
  API specifications
  ===================
  1. Identity-> Admin:  Method->POST  URL_route: '/task_app/task_manager_service/api/v0.1/task/write' 
           Header: {
                      'control-type':application/json
                      'Authoriation':Bearer Token
                   }
           req Body {
                        Title         string,
                        Description   string,
                        status        string,
                        Deadline      string,
                    } 
          resp {
                 status: task created successfully!
               }
  2. Identity -> Admin/User: Method->GET   URL_route: '/task_app/task_manager_service/api/v0.1/task/read/:id'
           Header: {
                      'control-type':application/json
                      'Authorization':Bearer Token
                   } 
          resp {
                  Id            string,
                  Title         string,
                  Description   string,
                  status        string,
                  Deadline      string,
               }            
  3. Identity -> Admin: Method: PUT    URL_route: '/task_app/task_manager_service/api/v0.1/task/update_status/:id'

            Header: {
                      'control-type':application/json
                      'Authorization': Bearer Token
                    }
            req {
                    status: "Pending/In-progress/completed"            
        
                 }
            resp {
                   UpdatedTask
                 }
  4. Identity -> Admin: Method: PUT   url_route: '/task_app/task_manager_service/api/v0.1/task/update_schedule/:id'
            Header: {
                     'control-type':application/json
                     'Authorization': Bearer Token
                   }
            req {
                 schedule: "Pending/In-progress/compelted"
            } 
            resp {
                  UpdatedTask
            }
  5. Identity -> Admin: Method: DELETE url_route: '/task_app/task_manager_service/api/v0.1/task/remove/:id'
      Description: Enables an admin to remove a task identified by task_id 
            Header: {
                    'control-type': application/json
                    'Authorization': Bearer Token
            }
            resp {
              status: "a task identitifed by {task_id} removed successfully"
            }
  6. Identity -> Admin:  Method: DELETE url_route: '/task_app/task_manager_service/api/v0.1/tasks/remove_by_user'
     Description: Enables an admin to remove all of its tasks assigned to specific user  identified as user_id.
             Header: {
                    'control-type': application/json
                    'Authorization': Bearer Token
            }
            resp {
              status: "all  tasks assigned to a user identified by {task_id} removed successfully"
            }
  7. Identity -> Admin: Method: DELETE url_route: '/task_app/task_manager_service/api/v0.1/tasks/remove_by_owner'
      Description: Enables an admin to remove all its task he/she created
            Header: {
                  'control-type': application/json
                  'Authorization': Bearer Token
            }
            resp {
                  status: "all tasks created by admin are removed successfully"
            }
  8. Identity -> Admin: Method: GET url_route: '/task_app/task_manager_service/api/v0.1/tasks/list_by_admin'
     Description: Enables an admin to list all of its tasks he/she owned/created.
         Header: {
              'control-type': application/json
              'Authorization': Bearer Token
         }
         resp {
                {list of Task object owned by admin}
         }
  9. Identity -> User: Method: GET url_route: '/task_app/task_manager_service/api/v0.1/tasks/list_by_user'
      Description: Enables a user to read all of its tasks assigned by his/her admin
             Header: {
                 'control-type': application/json
                 'Authorization': Bearer Token
             }
             
             resp {
                  {list of Task object owned by user}
             }
             
Policy and authorization definition
------------------------------------
admin, task_app/task_manager_service/api/v0.1/task, POST
admin, task_app/task_manager_service/api/v0.1/task, GET
user,  task_app/task_manager_service/api/v0.1/task, GET
admin, task_app/task_manager_service/api/v0.1/task, PUT
admin, task_app/task_manager_service/api/v0.1/task, DELETE
