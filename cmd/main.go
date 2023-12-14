package main

import "fmt"

type ServerPool struct {
	name    string
	message string
	port    uint
}

func main() {
	serverPool := ServerPool{name: "mad", message: "server", port: 8080}
	fmt.Printf("server pool name: %v, message: %v. port: %d \n", serverPool.name, serverPool.message, serverPool.port)
}
