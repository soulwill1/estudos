package main

import (
	"github.com/soulwill1/estudos/configs"
	"github.com/soulwill1/estudos/router"
)

func main() {
	configs.Init()

	router.RouterInit()
}
