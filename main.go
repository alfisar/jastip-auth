package main

import (
	"jastip/config"
	"jastip/router"
)

func main() {
	config.Init()
	router := router.NewRouter()

	router.Listen("0.0.0.0:8801")
}
