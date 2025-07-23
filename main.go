package main

import (
	fiberRouter "jastip/router/http"
	grpcRoute "jastip/router/tcp"
	"sync"

	"github.com/alfisar/jastip-import/config"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		grpcRoute.Start()
	}()

	go func() {
		defer wg.Done()
		config.Init()
		router := fiberRouter.NewRouter()

		router.Listen("0.0.0.0:8801")
	}()

	wg.Wait()
}
