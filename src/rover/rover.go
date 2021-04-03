package rover

import "newstore/src/landscape"

type Rover struct {
	Coordinate landscape.Coordinate
	Direction  string
	Landscape  *landscape.Landscape
}


func New(coordinate landscape.Coordinate, direction string, land *landscape.Landscape) *Rover {
	return &Rover{
		Coordinate: coordinate,
		Direction:  direction,
		Landscape:  land,
	}
}