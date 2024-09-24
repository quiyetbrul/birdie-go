package logic

import (
	"birdie-go/src/data"
	"birdie-go/src/entities"
	"math/rand"
)

func ApplyGravity(state *data.GameState) {
	state.Bird.VelY += state.Gravity
	state.Bird.Y += state.Bird.VelY
}

func BirdJump(state *data.GameState) {
	state.Bird.VelY = -5
}

func CheckCollision(state *data.GameState) bool {
	bird := state.Bird
	for _, pipe := range state.Pipes {
		if birdCollidesWithPipe(bird, pipe, state.PipeGap) {
			return true
		}
	}
	return bird.Y > 600 || bird.Y < 0
}

func birdCollidesWithPipe(bird *entities.Bird, pipe entities.Pipe, pipeGap float64) bool {
	pipeWidth := 50.0
	if pipe.X < 100+bird.Radius && pipe.X+pipeWidth > 100-bird.Radius {
		if bird.Y-bird.Radius < pipe.Y || bird.Y+bird.Radius > pipe.Y+pipeGap {
			return true
		}
	}
	return false
}

func SpawnPipe(state *data.GameState) {
	topPipeHeight := rand.Float64()*(600/2) + 50
	state.AddPipe(entities.NewPipe(400, topPipeHeight))
}

func UpdatePipes(state *data.GameState) {
	for i := range state.Pipes {
		state.Pipes[i].X -= state.PipeSpeed
	}

	// Remove off-screen pipes
	if len(state.Pipes) > 0 && state.Pipes[0].X < -50 {
		state.RemovePipe()
		state.Score++
	}
}
