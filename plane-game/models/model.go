package models

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
}

// Update method is responsible for update the game logic
func (g *Game) Update() error {
	return nil
}
// Draw method is responsible for draw objects in the screen
func (g *Game) Draw(screen *ebiten.Image) {
}
// Layout method is responsible for the screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screnWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
