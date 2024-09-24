package presentation

import (
	"birdie-go/src/data"
	logic "birdie-go/src/processor"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	// If the game is over and space is pressed, reset the game
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

	// Game is running, handle bird movement and game logic
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

	// Draw pipes using vector.Path for rectangles
	for _, p := range g.state.Pipes {
		// Draw the top pipe
		drawFilledRect(screen, float32(p.X), 0, 50, float32(p.Y))

		// Draw the bottom pipe
		drawFilledRect(screen, float32(p.X), float32(p.Y+g.state.PipeGap), 50, float32(600-p.Y-g.state.PipeGap))
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

// drawFilledRect draws a filled rectangle using vector.Path
func drawFilledRect(screen *ebiten.Image, x, y, width, height float32) {
	var path vector.Path
	path.MoveTo(x, y)
	path.LineTo(x+width, y)
	path.LineTo(x+width, y+height)
	path.LineTo(x, y+height)
	path.Close()

	op := &ebiten.DrawTrianglesOptions{}
	op.FillRule = ebiten.FillAll

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	screen.DrawTriangles(vs, is, ebiten.NewImageFromImage(nil), op)
}
