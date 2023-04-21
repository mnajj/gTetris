package screen

import (
	"github.com/mnajj/simple-tetris/tetris"
	"github.com/nsf/termbox-go"
)

const (
	blockHeight = 2
	blockWidth  = blockHeight * 2

	skip            = 1
	playgroundTop   = skip
	playgroundBot   = tetris.Height + skip
	playgroundLeft  = skip
	playgroundRight = tetris.Width + skip
)

func drawScene() {
	if err := clearTerminal(); err != nil {
		panic(err)
	}
	drawPlaygroundFrame()
	drawBoard()
	drawPiece(*game.LandingPiece)
	if err := termbox.Flush(); err != nil {
		panic(err)
	}
}

func drawBoard() {
	for y := 0; y < tetris.Height; y++ {
		for x := 0; x < tetris.Width; x++ {
			if game.PointColors[y][x] != 0 {
				drawBlock(x+skip, y+skip, game.PointColors[y][x])
			}
		}
	}
}

func drawPiece(p tetris.Piece) {
	for iy, y := range p.Shape {
		for ix, x := range y {
			if x {
				x := p.Left + ix + skip
				y := p.Top + iy + skip
				drawBlock(x, y, p.Color)
			}
		}
	}
}

func drawBlock(x, y int, c termbox.Attribute) {
	y *= blockHeight
	x *= blockWidth
	for i := 0; i < blockHeight; i++ {
		for j := 0; j < blockWidth; j++ {
			termbox.SetCell(x+j, y+i, ' ', termbox.ColorDefault, c)
		}
	}
}

func drawPlaygroundFrame() {
	for y := 0; y < tetris.Height+(skip*2); y++ {
		for x := 0; x < tetris.Width+(skip*2); x++ {
			if !isPointInPlayground(x, y) {
				drawBlock(x, y, termbox.ColorLightGreen)
			}
		}
	}
}

func isPointInPlayground(x, y int) bool {
	return y >= playgroundTop &&
		y < playgroundBot &&
		x >= playgroundLeft &&
		x < playgroundRight
}

func clearTerminal() error {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		return err
	}
	if err := termbox.Flush(); err != nil {
		return err
	}
	return nil
}
