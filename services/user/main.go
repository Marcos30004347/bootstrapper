package main

import (
	"user/database"
	"user/network"
)

func run() {
	// Connect to the database
	database.InitDatabase("mongodb://userdb:27017")

	// Init network services (REST, GRPC)
	network.InitNetword(30051, 30052)
	network.Run()
}

func main() {
	run()
}
