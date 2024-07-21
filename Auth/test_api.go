package main

import (
	"tele_auth/app"
)

func main() {

	// setup and run app fiber app
	err := app.SetupANDRun()
	if err != nil {
		panic(err)
	}

}
