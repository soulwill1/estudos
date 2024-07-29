package main

import (
	"plane-game/models"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &models.Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
