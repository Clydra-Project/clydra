package main

import (
	"context"

	"clydra.io/services/server"
)

func main() {
	ctx := context.Background()
	srv := server.InitServer(ctx)
	srv.Run(ctx)
}
