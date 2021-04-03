package rover

type Movement int

const (
	FORWARD  Movement = iota
	BACKWARD
)

type Rotation bool

const (
	CLOCKWISE 		Rotation = true
	ANTICLOCKWISE 	Rotation = false
)

func (rover *Rover) Listen(command string) {

	for _, char := range command {
		switch string(char) {
		case "F": {
			rover.Move(FORWARD)
		}
		case "B": {
			rover.Move(BACKWARD)
		}
		case "L": {
			rover.Rotate(ANTICLOCKWISE)
		}
		case "R": {
			rover.Rotate(CLOCKWISE)
		}
		}
	}
}

func (rover *Rover) Move(movement Movement) {
	delta := 1
	if movement == BACKWARD {
		delta = -1
	}
	switch rover.Direction {
	case "east": {
		rover.X += delta
	}
	case "west": {
		rover.X += (delta * -1)
	}
	case "north": {
		rover.Y += delta
	}
	case "south": {
		rover.Y += (delta * -1)
	}
	}
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