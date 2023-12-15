package main

import (
	"fmt"
	"rest_server/config"
)

func main() {
	config := config.MustLoad()

	fmt.Print(config)

	// TODO: create database entity struct
	// TODO: create MustLoad() for database
	// TODO: check if table exist + ping Database
}
