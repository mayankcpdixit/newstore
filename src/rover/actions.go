package rover

import (
	"newstore/src/landscape"
)

type Movement int

const (
	MOVE_FORWARD Movement = iota
	MOVE_BACKWARD
)

type Rotation bool

const (
	CLOCKWISE 		Rotation = true
	ANTICLOCKWISE 	Rotation = false
)

type State int

const (
	STOPPED State = iota
	MOBILE
)

func (rover *Rover) Listen(command string) (State){

	for _, char := range command {
		switch string(char) {
		case "F": {
			coordinate := rover.GetNextCoordinate(MOVE_FORWARD)
			if rover.Landscape.HasObstacleAt(coordinate) {
				return STOPPED
			}
			rover.Move(coordinate)
		}
		case "B": {
			coordinate := rover.GetNextCoordinate(MOVE_BACKWARD)
			if rover.Landscape.HasObstacleAt(coordinate) {
				return STOPPED
			}
			rover.Move(coordinate)
		}
		case "L": {
			rover.Rotate(ANTICLOCKWISE)
		}
		case "R": {
			rover.Rotate(CLOCKWISE)
		}
		// TODO: catch invalid command throw err
		}
	}
	return MOBILE
}

func (rover *Rover) Move(coordinate landscape.Coordinate) {
	rover.Coordinate = coordinate
}

func (rover *Rover) GetNextCoordinate(movement Movement) landscape.Coordinate {
	coordinate := rover.Coordinate
	delta := 1
	if movement == MOVE_BACKWARD {
		delta = -1
	}
	switch rover.Direction {
	case "east": {
		coordinate.X += delta
	}
	case "west": {
		coordinate.X += (delta * -1)
	}
	case "north": {
		coordinate.Y += delta
	}
	case "south": {
		coordinate.Y += (delta * -1)
	}
	// TODO: catch invalid direction throw err
	}
	return coordinate
}

func (rover *Rover) Rotate (rotation Rotation) {
	directions := []string{"east", "south", "west", "north"} // "clockwise ordered list" of direction. Do not change the order.
	delta := 1
	if rotation == ANTICLOCKWISE {
		delta = -1
	}
	for idx, dir := range directions {
		if dir == rover.Direction {
			rover.Direction = directions[(idx+4+delta)%4]
			break;
		}
	}
}

func (rover *Rover) PathFinder(destination landscape.Coordinate) string {
	dict := map[landscape.Coordinate]string{} // coordinate -> shortest_path_to_that_coordinate_from_rover_location
	return findPath(rover.Coordinate, destination, rover.Landscape.Obstacles, dict)
}


func findPath(source, destination landscape.Coordinate, obstacles []landscape.Coordinate, dict map[landscape.Coordinate]string) string {
	// TODO: Part 3
	return ""
}