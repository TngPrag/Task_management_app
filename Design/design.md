Project Description
===================
Client Description
------------------
-> Develop a web based task maangement application where users can register,log in,
   create tasks, assign tasks to other users, and mark tasks as complete. The application
   should have a both front-end and backend components.

Developer Description
---------------------

-> Architectural paradignm: Micro-service with clean-architecture
-> Design pattern: Domain Driven Design (DDD)
-> API type: RESTFULL HTTP API

conceptual description
 -> The Task management app will have multiple admins on which each of them could have their
    own users so the admin could create,assign and update tasks to those users, and users 
    which are under this admin could also view the tasks given by the admin, update task
    task status, and be notified when deadline schedule time reaches.
 -> Admin could also view status of over all tasks by all users.
      * number of completed tasks
      * number of in-progress tasks
      * number of uncompleted tasks
      * number of users managed by admin
      * status of each tasks on a table
Problem modelling framework
  -> Identity: describes user behavior and type of the task management app.
  -> Domain entities: intends to describe the atomics objects of the app.
  -> capabilities: clarifies the capability provided to each identities.
  -> Micro-services: a bounded and welframed context with clear separation of concern between 
                      defined domain entities.
  -> values: useful features provided to each identities.

1. Identities
   -> Admins: are users who manages tasks of specific group of users.
   -> Users: are individuals who implements the tasks and adminstered under under their admins
   -> super-admins: are owners of the app who creates/registers admins
2. Values
   -> Admins
          * Enables to define and assign tasks to their users.
          * Enables to view over all task status and performance of their users via dashboard.
          * Enables to be notified to email/alerts if tasks are not completed as per the 
             schedule.
  -> Users
        * Enables to view tasks assigned to them.
        * Could update status of tasks.
        * be notified via email when task is assigned to them.
3. Domain Entities
      * User {
                user_id       string
                Owner_id      string
                FirstName     string
                LastName      string
                Email         string
             } 

      * Task {
               task_id      string
               User_id      string
               Owner_id     string
               Title        string 
               Description  string
               Status       string
               Deadline     string                
             }
      * Admin {
                admin_id       string
                owner_id       string
                FirstName      string
                LastName       string
                Email          string
              }
      *Super-Admin {
                     super_admin_id   string
                     FirstName      string
                     LastName       string
                     Email          string
                   }


4. Capabilities:
      * Super-admin 
      * Super-admin | Admin        :POST: auth_service/auth/register
      * SUper-admin | Admin | user :POST: auth_service/auth/login
      * super-admin | Admin | user :POST: auth_service/auth/verify
      
      * Super-admin:Method:POST: api_url_route:  user_management_service/admin/register
         Description: Enables to create admins  
         Authorization: Bearer Token
              req {
                      FirstName  string
                      LastName   string
                      Email      string
                    } 
                resp {
                       {id: string status: created successfully!}
                     }
     * Super-admin:Method: GET : api_url_route: user_management_service/admin/read/
       Description: Enables to read all admins
       Authorization: Bearer Token 
            req {
                  ''
                }
             resp :
                 {
		  "admins": [
		    {
		      "id": "string",
		      "first_name": "string",
		      "last_name": "string",
		      "email": "string"
		    }
		  ]
		}

     * Super-admin:Method: DELETE: API_URL_ROUTE: user_management_service/admin/remove/{id}
        Description: Enables to delete an admin.
        Authorization: Bearer Token
             req {
                   id string
                 }
              resp {
                     {"status":"admin identified by {id} is deleted successfully"}
                   }
     * Admin: Method: POST: api_url_route:user_management_service/user/register
             Description: Enables users to register their usres.
             Authorization: Bearer Token
             req {
                    FirstName    string
                    FatherName   string
                    Email        string
                  }
            resp {
                    "successfully created"
                 }
     * Admin: Method: GET: api_url_route: user_management_service/user/read_all
            Description: Enables to read all users owned by the admin
            Authorization: Bearer Token
            req {
                  ''
                }
            resp :
                 {
			  "users": [
			    {
			      "id": "string",
			      "first_name": "string",
			      "last_name": "string",
			      "email": "string"
			    }
			  ]
		}

                  
     * Admin: Method: DELETE: api_url_route: user_management_service/user/remove/{id}
             Description: Enables to remove a user identified by user_id
             Authorization: Bearer Token
             req {
                   id
                 }
            resp :
                  {
  			"status": "successfully deleted: {id}"
		  }

                 
     * Admin: Method: POST: api_url_route:task_management_service/task/write
          Description: Enables to create a task
            Authorization: Bearer Token 
            req {
                  TaskTitle     string
                  Description   string
                  Schedule      string 
                } 
            resp {
                   {"status":"Ok", resp: "task created successfully"}
                 }
     * Admin: Method: GET: api_url_route:task_management_service/task/read_all
          Description: Enables to read all tasks created by this admin
             Authorization: Bearer Token
             
             resp :
                   {
  			"tasks": [
    				{
      					"id": "string",
      					"title": "string",
     					"description": "string",
      					"status": "string",
      					"deadline": "string"
    				}
  				]
		}

     * Admin: Method: Get: api_url_route:task_management_service/task/remove/{id}
           Description: Enables to remove a task by id
           Authorization: Bearer TOken
           req {
                task_id
               }
           resp :
              	{
  			"status": "successfully removed: {id}"
		}

     *Admin: Method: DELETE: api_url_route: task_management_service/task/remove_all
           Description: Enables to remove all tasks of the admin
           Authorization: Bearer: Token
           
           resp :
              {
  		"status": "removed all"
	      }

     * Admin: Method: GET: api_url_route: task_management_service/task/analytics/read
            Description: Enables to read analytics of tasks
            Authorization: Bearer Token
            
           resp :
               {
  		"analytics": {
    				"completed_tasks": "number",
    				"in_progress_tasks": "number",
  			        "not_completed_tasks": "number"
  		         	}
		}

     *Admin: Method: POST:API_URL_ROUTE: task_management_service/task/assign
          Description: Enables to assign a task identified by task_id to user identified by id
          Authorization: Bearer Token
          req :
                {
  			"task_id": "string",
  			"user_id": "string"
		}

          resp :
                 {
  			"status": "assigned successfully"
		 }

     *user: Method: Get: api_url_route: task_management_service/task/read_all
           Description: Enables to read all tasks given by the admin
           Authorization: Bearer Token
           req {
                 ''
               } 
           resp {
                 "returns all tasks given to the user "
                }
     *user: Method: PUT: api_url_route: task_management_service/task/update_status
            Decription: Enables to update status of a task
            Authorization: Bearer Token
            req :
                 {
  		      "task_id": "string",
  		      "status": "completed" | "in_progress" | "not_completed"
                 }
            resp :
                {
  			"status": "updated successfully"
		}


     * User: Method: GET: api_url_route: task_management_service/task/analytics/read
            Description: Enables to read analytics of tasks of the given user
            Authorization: Bearer Token
            req {
                 ''
               }
           resp {
                  {analytics_data}
                }

5. Microservices
    A. Authentication and authorization service
            database: Postgresql
            Transport: http protocol
            programming language: Golang
            Web Framework: Fiber or go-chi
            Monitoring tool: Grafana and prometheus
            Container: Docker
    B. User management service
            database: Postgrsql
            Transport: http protocol
            Programming language: Golang 
            web framework: Fiber or go-chi
            Container: Docker
    C. Task management service
            database: Mongodb
            Transport: http protocol
            Programing language: Golang/javascript
            Web-framework: fiber, go-chi or node.js
            Container: Docker
   D. UI
        Transport: http 
        Programming language: javascript
        Framework and other tools: React.js, html5, css tailwind
        Container: Docker
        Multiple-constainer: Docker-compose
        
        
    
             

