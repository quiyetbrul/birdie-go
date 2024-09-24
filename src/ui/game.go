package ui

import (
	"birdie-go/src/data"
	logic "birdie-go/src/processor"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	state *data.GameState
}

func NewGame() *Game {
	return &Game{
		state: data.NewGameState(),
	}
}

func (g *Game) Update() error {
	if g.state.GameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.state.ResetGame()
		}
		return nil
	}

	// If the game hasn't started, wait for space to start
	if !g.state.GameStarted {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.state.GameStarted = true // Start the game
		}
		return nil
	}

	// Bird physics and input
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		logic.BirdJump(g.state)
	}

	logic.ApplyGravity(g.state)
	logic.UpdatePipes(g.state)

	// Spawn pipes
	if len(g.state.Pipes) == 0 || g.state.Pipes[len(g.state.Pipes)-1].X < 400-250 {
		logic.SpawnPipe(g.state)
	}

	// Check for collision
	if logic.CheckCollision(g.state) {
		g.state.GameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Set color transformation matrix for drawing the bird
	var clrM colorm.ColorM
	clrM.Scale(1, 1, 1, 1) // Default white color

	// Draw bird
	ebitenutil.DrawCircle(screen, 100, g.state.Bird.Y, g.state.Bird.Radius, clrM.Apply(color.White))

	for _, p := range g.state.Pipes {
		// Top pipe
		ebitenutil.DrawRect(screen, p.X, 0, 50, p.Y, clrM.Apply(color.White))

		// Bottom pipe
		ebitenutil.DrawRect(screen, p.X, p.Y+g.state.PipeGap, 50, 600-p.Y-g.state.PipeGap, clrM.Apply(color.White))
	}

	// Draw score
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.state.Score))

	// If the game hasn't started, show "Press Space to Start" message
	if !g.state.GameStarted {
		ebitenutil.DebugPrintAt(screen, "Press Space to Start", 100, 300)
	}

	// Game over message
	if g.state.GameOver {
		ebitenutil.DebugPrintAt(screen, "Game Over! Press Space to Restart", 100, 300)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 600
}
