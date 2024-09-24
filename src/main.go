package main

import (
	"log"

	"birdie-go/src/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(400, 600)
	ebiten.SetWindowTitle("Flappy Bird")
	game := ui.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
