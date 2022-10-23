package main

import (
	"os"

	"clydra.io/cmd/clydra-server/command"
)

func main() {
	os.Exit(command.RunServer())
}
