package tetris

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

var (
	j = shape{{true, false, false}, {true, true, true}}
	l = shape{{false, false, true}, {true, true, true}}

	z = shape{{true, true, false}, {false, true, true}}
	s = shape{{false, true, true}, {true, true, false}}

	t = shape{{false, true, false}, {true, true, true}}
	i = shape{{true, true, true, true}}
	o = shape{{true, true}, {true, true}}

	shapes = [...]shape{j, l, z, s, t, i, o}
	colors = [...]termbox.Attribute{
		termbox.ColorWhite,
		termbox.ColorBlue,
		termbox.ColorCyan,
		termbox.ColorYellow,
		termbox.ColorGreen,
		termbox.ColorRed,
		termbox.ColorMagenta}
)

type shape [][]bool

func (s shape) clockwiseRotate() shape {
	h, w := len(s[0]), len(s)
	if h == w {
		return s
	}
	newSlice := make([][]bool, h)
	for i := range newSlice {
		newSlice[i] = make([]bool, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			newSlice[i][j] = s[w-1-j][i]
		}
	}
	return newSlice
}

type Piece struct {
	Shape     shape
	Color     termbox.Attribute
	Left, Top int
}

func (p *Piece) Rotate() {
	p.Shape = p.Shape.clockwiseRotate()
}

func (p *Piece) Width() int {
	return len(p.Shape[0])
}

func (p *Piece) Height() int {
	return len(p.Shape)
}

func PickRandom() *Piece {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	sh := shapes[r.Intn(len(shapes))]
	return &Piece{
		Shape: sh,
		Color: colors[r.Intn(len(colors))],
		Left:  0,
		Top:   0,
	}
}
