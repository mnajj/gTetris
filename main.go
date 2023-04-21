package main

import (
	"github.com/mnajj/simple-tetris/screen"
	"github.com/mnajj/simple-tetris/tetris"
)

func main() {
	screen.New(tetris.New())
}
