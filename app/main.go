package main

import "example.com/m/infrastructure"

func main() {
	defer infrastructure.DB.Close()
	infrastructure.EchoRouter.Run()
}