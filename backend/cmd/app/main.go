package main

import (
	"pollparlor/internal/app/server"
)

func main() {
	println("Hello, world!")

	srv := server.NewServer(nil)
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
