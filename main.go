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
	win       bool
}

type Board struct {
	tiles [][]*Tile
}

type Tile struct {
	Img   *ebiten.Image
	Value int
}

func (g *Game) Update() error {
	if g.moves > 0 && !g.player.win {
		handleMouseClick(g)
	}
	handleKeyPress(g)

	return nil
}

func handleMouseClick(g *Game) {
	if inpututil.IsMouseButtonJustPressed(0) {
		x, y := ebiten.CursorPosition()
		fmt.Printf("X: %d\nY: %d\n", x, y)
		px, py := whereWasClicked(x, y)
		fmt.Printf("X: %d\nY: %d\n", px, py)

		if !isMoveValid(g, px, py) {
			g.message = fmt.Sprintf(MESSAGE_PLAYAGAIN, g.player.IsPlaying)
		} else {
			g.moves--
			g.board.tiles[px][py].Value = g.player.IsPlaying
		}

		handleBoard(g)
		if g.player.win {
			g.message = fmt.Sprintf(MESSAGE_WINNER, g.player.IsPlaying)
		} else if g.moves == 0 {
			g.message = MESSAGE_DRAW
		} else {
			g.player.IsPlaying = nextToPlay(g.player.IsPlaying)
			g.message = fmt.Sprintf(MESSAGE_WHOPLAYS, g.player.IsPlaying)
		}
	}
}

func restartGame(g *Game) {
	g.moves = 9
	g.player.IsPlaying = rand.Intn(2) + 1
	g.message = fmt.Sprintf(MESSAGE_WHOPLAYS, g.player.IsPlaying)
	g.player.win = false
	for i := 0; i < ROW; i++ {
		for j := 0; j < COLUMN; j++ {
			g.board.tiles[i][j].Value = -1
		}
	}

}

func handleKeyPress(g *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		restartGame(g)
	}
}

func handleBoard(g *Game) {
	if g.board.tiles[0][0].Value == 1 && g.board.tiles[0][1].Value == 1 && g.board.tiles[0][2].Value == 1 || g.board.tiles[0][0].Value == 2 && g.board.tiles[0][1].Value == 2 && g.board.tiles[0][2].Value == 2 {
		g.player.win = true
	} else if g.board.tiles[1][0].Value == 1 && g.board.tiles[1][1].Value == 1 && g.board.tiles[1][2].Value == 1 || g.board.tiles[1][0].Value == 2 && g.board.tiles[1][1].Value == 2 && g.board.tiles[1][2].Value == 2 {
		g.player.win = true
	} else if g.board.tiles[2][0].Value == 1 && g.board.tiles[2][1].Value == 1 && g.board.tiles[2][2].Value == 1 || g.board.tiles[2][0].Value == 2 && g.board.tiles[2][1].Value == 2 && g.board.tiles[2][2].Value == 2 {
		g.player.win = true
	}

	if g.board.tiles[0][0].Value == 1 && g.board.tiles[1][0].Value == 1 && g.board.tiles[2][0].Value == 1 || g.board.tiles[0][0].Value == 2 && g.board.tiles[1][0].Value == 2 && g.board.tiles[2][0].Value == 2 {
		g.player.win = true
	} else if g.board.tiles[0][1].Value == 1 && g.board.tiles[1][1].Value == 1 && g.board.tiles[2][1].Value == 1 || g.board.tiles[0][1].Value == 2 && g.board.tiles[1][1].Value == 2 && g.board.tiles[2][1].Value == 2 {
		g.player.win = true
	} else if g.board.tiles[0][2].Value == 1 && g.board.tiles[1][2].Value == 1 && g.board.tiles[2][2].Value == 1 || g.board.tiles[0][2].Value == 2 && g.board.tiles[1][2].Value == 2 && g.board.tiles[2][2].Value == 2 {
		g.player.win = true
	}

	if g.board.tiles[0][0].Value == 1 && g.board.tiles[1][1].Value == 1 && g.board.tiles[2][2].Value == 1 || g.board.tiles[0][0].Value == 2 && g.board.tiles[1][1].Value == 2 && g.board.tiles[2][2].Value == 2 {
		g.player.win = true
	} else if g.board.tiles[0][2].Value == 1 && g.board.tiles[1][1].Value == 1 && g.board.tiles[2][0].Value == 1 || g.board.tiles[0][2].Value == 2 && g.board.tiles[1][1].Value == 2 && g.board.tiles[2][0].Value == 2 {
		g.player.win = true
	}
}

func isMoveValid(g *Game, px int, py int) bool {
	if px != -1 && py != -1 {
		return g.board.tiles[px][py].Value == -1
	}
	return false
}

func whereWasClicked(x int, y int) (int, int) {
	if x >= 0 && x <= FRAME_WIDTH*3 && y >= 0 && y <= FRAME_HEIGTH*3 {
		if x >= 0 && x <= FRAME_WIDTH {
			if y >= 0 && y <= FRAME_HEIGTH {
				return 0, 0
			} else if y >= FRAME_HEIGTH && y <= FRAME_HEIGTH*2 {
				return 0, 1
			} else if y >= FRAME_HEIGTH*2 && y <= FRAME_HEIGTH*3 {
				return 0, 2
			}
		}
		if x >= FRAME_HEIGTH && x <= FRAME_HEIGTH*2 {
			if y >= 0 && y <= FRAME_WIDTH {
				return 1, 0
			} else if y >= FRAME_WIDTH && y <= FRAME_WIDTH*2 {
				return 1, 1
			} else if y >= FRAME_WIDTH*2 && y <= FRAME_WIDTH*3 {
				return 1, 2
			}
		}
		if x >= FRAME_HEIGTH*2 && x <= FRAME_HEIGTH*3 {
			if y >= 0 && y <= FRAME_WIDTH {
				return 2, 0
			} else if y >= FRAME_WIDTH && y <= FRAME_WIDTH*2 {
				return 2, 1
			} else if y >= FRAME_WIDTH*2 && y <= FRAME_WIDTH*3 {
				return 2, 2
			}
		}
	}
	return -1, -1
}

func nextToPlay(whosPlaying int) int {
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
		moves:   9,
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
