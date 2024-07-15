package main

import (
	"github.com/soulwill1/estudos/config"
	"github.com/soulwill1/estudos/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}


	// Initialize router
	router.Initialize()
}