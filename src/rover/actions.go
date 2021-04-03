package rover

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
			x,y := rover.GetNextCoordinate(MOVE_FORWARD)
			if rover.Landscape.HasObstacleAt(x,y) {
				return STOPPED
			}
			rover.Move(x,y)
		}
		case "B": {
			x,y := rover.GetNextCoordinate(MOVE_BACKWARD)
			if rover.Landscape.HasObstacleAt(x,y) {
				return STOPPED
			}
			rover.Move(x,y)
		}
		case "L": {
			rover.Rotate(ANTICLOCKWISE)
		}
		case "R": {
			rover.Rotate(CLOCKWISE)
		}
		}
	}
	return MOBILE
}

func (rover *Rover) Move(x,y int) {
	rover.X, rover.Y = x, y
}

func (rover *Rover) GetNextCoordinate(movement Movement) (int, int){
	x,y := rover.X, rover.Y
	delta := 1
	if movement == MOVE_BACKWARD {
		delta = -1
	}
	switch rover.Direction {
	case "east": {
		x += delta
	}
	case "west": {
		x += (delta * -1)
	}
	case "north": {
		y += delta
	}
	case "south": {
		y += (delta * -1)
	}
	}
	return x,y
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