package data

import (
	"birdie-go/src/entities"
	"math/rand"
	"time"
)

type GameState struct {
	Bird        *entities.Bird
	Pipes       []entities.Pipe
	PipeGap     float64
	Score       int
	GameStarted bool // New field to track if the game has started
	GameOver    bool
	Gravity     float64
	PipeSpeed   float64
	rng         *rand.Rand
}

func NewGameState() *GameState {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &GameState{
		Bird:        entities.NewBird(),
		Pipes:       []entities.Pipe{},
		PipeGap:     150,
		Score:       0,
		GameStarted: false,
		GameOver:    false,
		Gravity:     0.25,
		PipeSpeed:   2,
		rng:         rng,
	}
}

func (gs *GameState) AddPipe(pipe entities.Pipe) {
	gs.Pipes = append(gs.Pipes, pipe)
}

func (gs *GameState) RemovePipe() {
	if len(gs.Pipes) > 0 {
		gs.Pipes = gs.Pipes[1:]
	}
}

func (gs *GameState) ResetGame() {
	gs.Bird = entities.NewBird()
	gs.Pipes = nil
	gs.Score = 0
	gs.GameOver = false
	gs.GameStarted = false
}
