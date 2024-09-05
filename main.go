package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 240
	SCREEN_HEIGHT = 320

	FRAME_WIDTH  = 16
	FRAME_HEIGHT = 16
)

type Game struct {
	*Player
}

type Player struct {
	isPlaying bool
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 180, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("TIC TAC TOE by Rafael Goulart")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
