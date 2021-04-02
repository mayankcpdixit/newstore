package rover


func (rover *Rover) Move(command string) {
	for _, char := range command {
		switch string(char) {
			case "F": {
				if rover.Direction == "east" {
					rover.X++
				} else if rover.Direction == "west" {
					rover.X--
				} else if rover.Direction == "north" {
					rover.Y++
				} else if rover.Direction == "south" {
					rover.Y--
				}
			}
			case "B": {
				if rover.Direction == "east" {
					rover.X--
				} else if rover.Direction == "west" {
					rover.X++
				} else if rover.Direction == "north" {
					rover.Y--
				} else if rover.Direction == "south" {
					rover.Y++
				}
			}
			case "L": {
				rover.Rotate(false)
			}
			case "R": {
				rover.Rotate(true)
			}

		} // switch ends
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