package tetris

import (
	"github.com/nsf/termbox-go"
)

const (
	Left = iota
	Right
	Rotate
	Down
)

const (
	Width  = 10
	Height = 20
	block  = 1
)

type board [][]bool
type Game struct {
	Board        board
	PointColors  [][]termbox.Attribute
	LandingPiece *Piece
}

func New() *Game {
	return &Game{
		Board:       initBoard(),
		PointColors: initPointsColors(),
	}
}

func initPointsColors() [][]termbox.Attribute {
	slice := make([][]termbox.Attribute, Height)
	for i := range slice {
		slice[i] = make([]termbox.Attribute, Width)
	}
	return slice
}

func (g *Game) IsLandingPieceLanded() bool {
	p := g.LandingPiece
	return p.Top >= Height-p.Height() || !g.boardHasSpace(Down)
}

func (g *Game) GameOver() bool {
	return g.IsLandingPieceLanded() && g.LandingPiece.Top == 0
}

func (g *Game) WriteLandingPieceToBoard() {
	p := g.LandingPiece
	for i := p.Top; i < p.Top+p.Height(); i++ {
		for j := p.Left; j < p.Left+p.Width(); j++ {
			if p.Shape[i-p.Top][j-p.Left] && !g.Board[i][j] {
				g.Board[i][j] = p.Shape[i-p.Top][j-p.Left]
				g.PointColors[i][j] = p.Color
				//g.PointColors[Point{i, j}] = p.Color
			}
		}
	}
}

func (g *Game) MoveLandingPiece(dir int) {
	p := g.LandingPiece
	if g.willBeInBoard(dir) && g.boardHasSpace(dir) {
		switch dir {
		case Left:
			p.Left--
		case Right:
			p.Left++
		case Down:
			p.Top++
		case Rotate:
			p.Rotate()
		}
	}
}

func (g *Game) boardHasSpace(dir int) bool {
	var secHeight, secWidth int
	p := *g.LandingPiece
	switch dir {
	case Left:
		secHeight, secWidth = p.Top, p.Left-block
	case Right:
		secHeight, secWidth = p.Top, p.Left+block
	case Down:
		secHeight, secWidth = p.Top+block, p.Left
	case Rotate:
		p.Rotate()
		secHeight, secWidth = p.Top, p.Left
	default:
		panic("out of range input")
	}
	for i := secHeight; i < secHeight+p.Height(); i++ {
		for j := secWidth; j < secWidth+p.Width(); j++ {
			if g.Board[i][j] && p.Shape[i-secHeight][j-secWidth] {
				return false
			}
		}
	}
	return true
}

func (g *Game) willBeInBoard(dir int) bool {
	p := *g.LandingPiece
	h, w := p.Height(), p.Width()
	t, l := p.Top, p.Left
	switch dir {
	case Left:
		l--
	case Right:
		l++
	case Down:
		t++
	case Rotate:
		h, w = w, h
	}
	return l >= Width-Width &&
		l+w <= Width &&
		t >= Height-Height &&
		t+h <= Height
}

func (g *Game) CheckForFilledRows() {
	var filled []int
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if !g.Board[i][j] {
				break
			}
			if j == Width-block {
				filled = append(filled, i)
			}
		}
	}
	g.removeFilledRows(filled)
}

func (g *Game) removeFilledRows(rows []int) {
	for _, row := range rows {
		g.Board = append(g.Board[:row], g.Board[row+1:]...)
		g.Board = append([][]bool{make([]bool, Width)}, g.Board...)
		g.PointColors = append(g.PointColors[:row], g.PointColors[row+1:]...)
		g.PointColors = append([][]termbox.Attribute{make([]termbox.Attribute, Width)}, g.PointColors...)
	}
}

func initBoard() board {
	slice := make([][]bool, Height)
	for i := range slice {
		slice[i] = make([]bool, Width)
	}
	return slice
}
