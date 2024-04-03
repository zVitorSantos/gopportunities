package main

import (
	"github.com/zVitorSantos/gopportunities.git/config"
	"github.com/zVitorSantos/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config inirialization error: %v", err)
		return
	}

	router.Initialize()
}
