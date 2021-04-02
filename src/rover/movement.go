package rover

func (rover *Rover) Listen(command string) {
	for _, char := range command {
		switch string(char) {
		case "F": {
			rover.Move(false)
		}
		case "B": {
			rover.Move(true)
		}
		case "L": {
			rover.Rotate(false)
		}
		case "R": {
			rover.Rotate(true)
		}
		}
	}
}

func (rover *Rover) Move(backward bool) {
	delta := 1
	if backward {
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

func (rover *Rover) Rotate (Clockwise bool) {
	directions := []string{"east", "south", "west", "north"} // "clockwise ordered list" of direction. Do not change the order.
	delta := 1
	if !Clockwise {
		delta = -1
	}
	for idx, dir := range directions {
		if dir == rover.Direction {
			rover.Direction = directions[(idx+4+delta)%4]
			break;
		}
	}
}