package main

import (
	"task/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":9090")
}

