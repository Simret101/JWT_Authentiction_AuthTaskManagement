package main

import (
	"task/config"
	"task/router"
)

func main() {

	r := router.SetupRouter()

	r.Run(config.ServerPort)
}
