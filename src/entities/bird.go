package entities

type Bird struct {
	Y      float64
	VelY   float64
	Radius float64
}

func NewBird() *Bird {
	return &Bird{
		Y:      300, // Starting Y position
		VelY:   0,
		Radius: 15,
	}
}
