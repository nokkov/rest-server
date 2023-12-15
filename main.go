package main

import (
	"fmt"
	"rest_server/config"
)

func main() {
	config := config.MustLoad()

	fmt.Print(config)
}
