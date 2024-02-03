package main

import (
	"github.com/flevin58/staticserver/handlers"
	"github.com/flevin58/staticserver/server"
)

func main() {
	server := server.New().WithAddress(":8080").WithStaticDir("static-www")
	server.AddRoute("/pippo", "GET", handlers.Pippo)
	server.Run()
}
