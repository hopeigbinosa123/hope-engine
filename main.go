package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	// Update game logic here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render content here
	ebitenutil.DebugPrint(screen, "Welcome to Hope Engine!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480 // Window resolution
}

func main() {
	ebiten.SetWindowTitle("Hope Engine 2D")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}