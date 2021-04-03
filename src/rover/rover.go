package rover

import "newstore/src/landscape"

type Rover struct {
	X int
	Y int
	Direction string
	Landscape *landscape.Landscape
}


func New(x, y int, direction string, land *landscape.Landscape) *Rover {
	return &Rover{
		X: x,
		Y: y,
		Direction: direction,
		Landscape: land,
	}
}