package screen

import (
	"log"
	"os"
	"time"

	"github.com/mnajj/simple-tetris/tetris"

	"github.com/nsf/termbox-go"
)

var game *tetris.Game

func New(g *tetris.Game) {
	game = g
	initLandingPiece()
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	mainLoop()
}

func initLandingPiece() {
	newRandomLandingPiece()
	game.LandingPiece.Top = -1
}

func mainLoop() {
	inputCh := make(chan termbox.Key)
	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc:
					exit()
				case
					termbox.KeyArrowDown,
					termbox.KeyArrowUp,
					termbox.KeyArrowLeft,
					termbox.KeyArrowRight:
					inputCh <- ev.Key
				}
			}
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			game.MoveLandingPiece(tetris.Down)
		case input := <-inputCh:
			mapKeyToHandler(input)
		}
		checkGameState()
		drawScene()
	}
}

func checkGameState() {
	if game.GameOver() {
		exit()
	}
	if game.IsLandingPieceLanded() {
		game.WriteLandingPieceToBoard()
		game.CheckForFilledRows()
		newRandomLandingPiece()
	}
}

func newRandomLandingPiece() {
	game.LandingPiece = tetris.PickRandom()
	game.LandingPiece.Left = game.LandingPiece.Width()
}

func mapKeyToHandler(k termbox.Key) {
	var dir int
	switch k {
	case termbox.KeyArrowRight:
		dir = tetris.Right
	case termbox.KeyArrowLeft:
		dir = tetris.Left
	case termbox.KeyArrowUp:
		dir = tetris.Rotate
	case termbox.KeyArrowDown:
		dir = tetris.Down
	}
	game.MoveLandingPiece(dir)
}

func exit() {
	if err := clearTerminal(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
