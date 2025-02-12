package main

import (
	"jastip/router"

	"github.com/alfisar/jastip-import/config"
)

func main() {
	config.Init()
	router := router.NewRouter()

	router.Listen("0.0.0.0:8801")
}
