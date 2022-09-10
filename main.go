package main

import (
	"clydra.io/app/handler"
	"clydra.io/app/system"
)

func main() {

	// Init system is required
	system.InitConfig()
	system.InitDbConnection()

	handler.StartHttpHandler()
}
