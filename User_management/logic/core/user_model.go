package core

import (
	"log"
	"time"
	"user_manager/fs"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type User struct {
	Id        string    `json:"id"`
	Owner_id  string    `json:"owner_id"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	// Role      string    `json:"role"`
	CreateAt  time.Time `json:"created_at"`
	UPdatedAt time.Time `json:"updated_at"`
}

// super-admin/admin/user
func InitSuperAdminUser() error {
	//hashedPassword,_ := HashPassword("Password123!")
	super_admin := User{
		Id:        uuid.NewString(),
		Owner_id:  "",
		Name:      "Tsegay Negassi",
		UserName:  "Tsegay@super-admin",
		Password:  "Password123!",
		Email:     "tsegaynega11@gmail.com",
		//Role:      "super-admin",
		CreateAt:  time.Now(),
		UPdatedAt: time.Now(),
	}
	//super_admin.Owner_id = super_admin.Id
	if err := super_admin.Create_user(); err != nil {
		return err
	}
	// create super admin role via authz proxy api
		
	return nil

}

func CreateUserSchema() error {
	queryCreateTable := `CREATE TABLE IF NOT EXISTS task_app_users (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		owner_id VARCHAR(255),
		name VARCHAR(255) NOT NULL,
		user_name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL,
		updated_at TIMESTAMPTZ NOT NULL
	);`
	err := fs.Fs_create_relational_table(queryCreateTable)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Create_user() error {
	//write sql query to insert the dot channel data
	//query := `INSERT INTO Dot_channel (thing_id,owner_id,thing_key,channel_id) VALUES ($1,$2,$3,$4);`
	hashedPassword,err:= HashPassword(user.Password)
	if err!= nil{
		log.Println("Unable to hash the password")
	}
	ds := goqu.Insert("task_app_users").Rows(
		goqu.Record{
			"id":         user.Id,
			"owner_id":   user.Owner_id,
			"name":       user.Name,
			"user_name":  user.UserName,
			"password":   hashedPassword,
			"email":      user.Email,
			"created_at": user.CreateAt,
			"updated_at": user.UPdatedAt,
		},
	)
	insertQuery, _, err := ds.ToSQL()
	if err != nil {
		return err
	}

	//log.Println(insertQuery)
	err = fs.Fs_write(insertQuery)
	if err != nil {
		return err
	}
	
	return nil
}
//func (user *User) Login_user() error{}
func (user *User) Get_user_by_uid() ([]byte, error){
	table := goqu.T("task_app_users")
        //column := goqu.C("thing_id")
        //value := "p0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"

        query := goqu.From(table).
                Select(goqu.Func("row_to_json", table)).
                Where(goqu.Ex{"id": user.Id, "owner_id": user.Owner_id}).Limit(1)

        // Generate the SQL command
        get_query, _, _ := query.ToSQL()
        data, err := fs.Fs_read(get_query)
        if err != nil {
                return nil, err
        }
        return data, nil
}
// get user by email and user_name 

func (user *User) Get_user_by_email_userName()([]byte, error){
	table := goqu.T("task_app_users")
        //column := goqu.C("thing_id")
        //value := "p0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"

        query := goqu.From(table).
                Select(goqu.Func("row_to_json", table)).
                Where(goqu.Ex{"email":user.Email,"user_name":user.UserName}).Limit(1)

        // Generate the SQL command
        get_query, _, _ := query.ToSQL()
        data, err := fs.Fs_read(get_query)
        if err != nil {
                return nil, err
        }
        return data, nil
}

func (user *User) Get_user_by_owner_id() ([]byte, error)  {
	table := goqu.T("task_app_users")
        //column := goqu.C("thing_id")
        //value := "p0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"

        query := goqu.From(table).
                Select(goqu.Func("row_to_json", table)).
                Where(goqu.Ex{"owner_id": user.Owner_id})

        // Generate the SQL command
        get_query, _, _ := query.ToSQL()
        data, err := fs.Fs_read(get_query)
        if err != nil {
                return nil, err
        }
        return data, nil
}
func (user *User) Remove_user_by_id() error {
	table := goqu.T("task_app_users")
        condition := goqu.Ex{"owner_id": user.Owner_id, "id":user.Id }

        // Build the query
        query := goqu.Delete(table).
                Where(condition)

        // Generate the SQL command
        sql_query, _, _ := query.ToSQL()

        err := fs.Fs_delete(sql_query)
        if err != nil {
                return err
        }
        return nil
}
func (user *User) Remove_user_by_owner() error {
	table := goqu.T("task_app_users")
        condition := goqu.Ex{"owner_id": user.Owner_id}

        // Build the query
        query := goqu.Delete(table).
                Where(condition)

        // Generate the SQL command
        sql_query, _, _ := query.ToSQL()

        err := fs.Fs_delete(sql_query)
        if err != nil {
                return err
        }
        return nil
}
//func (user *User) List_users_by_owner() ([]byte, error){}
// func AuthenticateUser(token string)(string,error){

// }
