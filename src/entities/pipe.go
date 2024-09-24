package entities

type Pipe struct {
	X float64
	Y float64
}

func NewPipe(x, y float64) Pipe {
	return Pipe{
		X: x,
		Y: y,
	}
}
