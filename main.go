package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	SCREEN_WIDTH  = 192
	SCREEN_HEIGHT = 208

	FRAME_OX     = 0
	FRAME_OY     = 0
	FRAME_WIDTH  = 64
	FRAME_HEIGTH = 64

	ROW    = 3
	COLUMN = 3
)

type Game struct {
	moves  int
	player *Player
	board  *Board
}

type Player struct {
	IsPlaying int
}

type Board struct {
	tiles []*Tile
}

type Tile struct {
	Img    *ebiten.Image
	Value  int
	Row    int
	Column int
}

func (g *Game) Update() error {
	handleClick(g)

	return nil
}

func handleClick(g *Game) {
	if inpututil.IsMouseButtonJustPressed(0) {
		g.moves++
		x, y := ebiten.CursorPosition()
		fmt.Printf("Plays: %d\nX: %d\nY: %d\n", g.moves, x, y)

		g.player.IsPlaying = nextToPlay(g.player.IsPlaying)
	}
}

func nextToPlay(whosPlaying int) int {
	switch whosPlaying {
	case 0:
		return 1
	}
	return 0
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 1})

	opts := ebiten.DrawImageOptions{}

	drawTiles(g, opts, screen)
	drawStats(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {

	tile, _, err := ebitenutil.NewImageFromFile("tiles.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		player: &Player{
			IsPlaying: 0,
		},
		board: &Board{
			tiles: []*Tile{
				{
					Img:    tile,
					Value:  0,
					Row:    0,
					Column: 0,
				},
				{
					Img:    tile,
					Value:  0,
					Row:    0,
					Column: 1,
				},
				{
					Img:    tile,
					Value:  0,
					Row:    0,
					Column: 2,
				},
				{
					Img:    tile,
					Value:  0,
					Row:    1,
					Column: 0,
				},
				{
					Img:    tile,
					Value:  0,
					Row:    1,
					Column: 1,
				},
				{
					Img:    tile,
					Value:  0,
					Row:    1,
					Column: 2,
				},
				{
					Img:    tile,
					Value:  0,
					Row:    2,
					Column: 0,
				},
				{
					Img:    tile,
					Value:  0,
					Row:    2,
					Column: 1,
				},
				{
					Img:    tile,
					Value:  0,
					Row:    2,
					Column: 2,
				},
			},
		},
	}

	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("TIC TAC TOE by Rafael Goulart")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func drawStats(g *Game, screen *ebiten.Image) {
	var whosPlaying string
	if g.player.IsPlaying == 0 {
		whosPlaying = "Player 1 turn"
	} else {
		whosPlaying = "Player 2 turn"
	}
	moves := fmt.Sprintf("Moves: %d", g.moves)

	ebitenutil.DebugPrintAt(screen, whosPlaying, 0, 192)
	ebitenutil.DebugPrintAt(screen, moves, 142, 192)
}

func drawTiles(g *Game, opts ebiten.DrawImageOptions, screen *ebiten.Image) {

	for _, tile := range g.board.tiles {
		switch tile.Value {
		case 0:
			opts.GeoM.Translate(float64(tile.Row)*64, float64(tile.Column)*64)
			screen.DrawImage(
				tile.Img.SubImage(
					image.Rect(FRAME_OX, FRAME_OY, FRAME_WIDTH, FRAME_HEIGTH),
				).(*ebiten.Image),
				&opts,
			)
		case 1:
			opts.GeoM.Translate(float64(tile.Row)*64, float64(tile.Column)*64)
			screen.DrawImage(
				tile.Img.SubImage(
					image.Rect(FRAME_OX, FRAME_OY+64, FRAME_WIDTH, FRAME_HEIGTH*2),
				).(*ebiten.Image),
				&opts,
			)
		case 2:
			opts.GeoM.Translate(float64(tile.Row)*64, float64(tile.Column)*64)
			screen.DrawImage(
				tile.Img.SubImage(
					image.Rect(FRAME_OX, FRAME_OY+128, FRAME_WIDTH, FRAME_HEIGTH*3),
				).(*ebiten.Image),
				&opts,
			)
		}
		opts.GeoM.Reset()
	}
}
