package main

import (
	"user_manager/app"
)

func dmain() {

	// setup and run app fiber app
	err := app.SetupANDRun()
	if err != nil {
		panic(err)
	}

}
