package rover

type Rover struct {
	X int
	Y int
	Direction string
}


func New(x, y int, direction string) *Rover {
	return &Rover{
		X: x,
		Y: y,
		Direction: direction,
	}
}