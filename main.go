package main

import (
	"github.com/mnajj/gTetris/screen"
	"github.com/mnajj/gTetris/tetris"
)

func main() {
	screen.New(tetris.New())
}
