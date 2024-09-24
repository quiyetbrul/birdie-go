package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quiyetbrul/birdie-go/ui"
)

func main() {
	ebiten.SetWindowSize(400, 600)
	ebiten.SetWindowTitle("Flappy Bird")
	game := ui.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
