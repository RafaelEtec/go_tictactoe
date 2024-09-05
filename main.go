package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"

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

	MESSAGE_WHOPLAYS  = "P%d turn"
	MESSAGE_MOVES     = "Moves: %d"
	MESSAGE_PLAYAGAIN = "P%d, move invalid!"
	MESSAGE_DRAW      = "It's a Draw!"
	MESSAGE_WINNER    = "P%d Wins!"
)

type Game struct {
	moves   int
	message string
	player  *Player
	board   *Board
}

type Player struct {
	IsPlaying int
}

type Board struct {
	tiles [][]*Tile
}

type Tile struct {
	Img   *ebiten.Image
	Value int
}

func (g *Game) Update() error {
	handleClick(g)
	//handleMatch(g)

	return nil
}

func handleClick(g *Game) {
	if inpututil.IsMouseButtonJustPressed(0) {
		x, y := ebiten.CursorPosition()
		px, py := whereWasClicked(x, y)

		if !isMoveValid(g, px, py) {
			g.message = fmt.Sprintf(MESSAGE_PLAYAGAIN, g.player.IsPlaying)
		} else {
			g.moves++
			g.player.IsPlaying = nextToPlay(g, g.player.IsPlaying)
			g.message = fmt.Sprintf(MESSAGE_WHOPLAYS, g.player.IsPlaying)
		}
	}
}

func isMoveValid(g *Game, px int, py int) bool {
	if px != -1 && py != -1 {
		return g.board.tiles[px][py].Value == -1
	}
	return false
}

func whereWasClicked(x int, y int) (int, int) {
	if x > 0 && x < FRAME_HEIGTH*3 && y > 0 && y < FRAME_WIDTH*3 {
		// row 0
		if y > 0 && y < FRAME_HEIGTH {
			if x > 0 && x < FRAME_WIDTH {
				return 0, 0
			} else if x > FRAME_WIDTH && x < FRAME_WIDTH*2 {
				return 0, 1
			} else if x > FRAME_WIDTH*2 && x < FRAME_WIDTH*3 {
				return 0, 2
			}
		}
		// row 1
		if y > FRAME_HEIGTH && y < FRAME_HEIGTH*2 {
			if x > 0 && x < FRAME_WIDTH {
				return 1, 0
			} else if x > FRAME_WIDTH && x < FRAME_WIDTH*2 {
				return 1, 1
			} else if x > FRAME_WIDTH*2 && x < FRAME_WIDTH*3 {
				return 1, 2
			}
		}
		// row 2
		if y > FRAME_HEIGTH*2 && y < FRAME_HEIGTH*3 {
			if x > 0 && x < FRAME_WIDTH {
				return 2, 0
			} else if x > FRAME_WIDTH && x < FRAME_WIDTH*2 {
				return 2, 1
			} else if x > FRAME_WIDTH*2 && x < FRAME_WIDTH*3 {
				return 2, 2
			}
		}
	}
	return -1, -1
}

func nextToPlay(g *Game, whosPlaying int) int {
	switch whosPlaying {
	case 1:
		return 2
	}
	return 1
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
	firstToPlay := rand.Intn(2) + 1

	tile, _, err := ebitenutil.NewImageFromFile("tiles.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		message: fmt.Sprintf(MESSAGE_WHOPLAYS, firstToPlay),
		moves:   0,
		player: &Player{
			IsPlaying: firstToPlay,
		},
		board: &Board{
			tiles: [][]*Tile{
				{
					{
						Img:   tile,
						Value: -1,
					},
					{
						Img:   tile,
						Value: -1,
					},
					{
						Img:   tile,
						Value: -1,
					},
				}, {
					{
						Img:   tile,
						Value: -1,
					},
					{
						Img:   tile,
						Value: -1,
					},
					{
						Img:   tile,
						Value: -1,
					},
				}, {
					{
						Img:   tile,
						Value: -1,
					},
					{
						Img:   tile,
						Value: -1,
					},
					{
						Img:   tile,
						Value: -1,
					},
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
	moves := fmt.Sprintf(MESSAGE_MOVES, g.moves)

	ebitenutil.DebugPrintAt(screen, g.message, 0, 192)
	ebitenutil.DebugPrintAt(screen, moves, 142, 192)
}

func drawTiles(g *Game, opts ebiten.DrawImageOptions, screen *ebiten.Image) {
	for i := 0; i < ROW; i++ {
		for j := 0; j < COLUMN; j++ {
			tile := g.board.tiles[i][j]
			switch tile.Value {
			case -1:
				opts.GeoM.Translate(float64(i)*64, float64(j)*64)
				screen.DrawImage(
					tile.Img.SubImage(
						image.Rect(FRAME_OX, FRAME_OY, FRAME_WIDTH, FRAME_HEIGTH),
					).(*ebiten.Image),
					&opts,
				)
			case 1:
				opts.GeoM.Translate(float64(i)*64, float64(j)*64)
				screen.DrawImage(
					tile.Img.SubImage(
						image.Rect(FRAME_OX, FRAME_OY+64, FRAME_WIDTH, FRAME_HEIGTH*2),
					).(*ebiten.Image),
					&opts,
				)
			case 2:
				opts.GeoM.Translate(float64(i)*64, float64(j)*64)
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
}
