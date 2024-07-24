package main

import (
	"user_manager/app"
)

// @title Task User management Service
// @version 0.1
// @description This is used to manage users of the app
// @termsOfService http://swagger.io/terms/
// @contact.name Tsegay Negassi
// @contact.email tng.nat2023@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8981
// @BasePath /task_app/user_manager_service/api/v0.1
func main() {

	// setup and run app fiber app
	err := app.SetupANDRun()
	if err != nil {
		panic(err)
	}

}
