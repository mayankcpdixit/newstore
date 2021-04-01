package rover
/*
	1. improve code quality for Move
	2. more tests
*/
/*
	east 	F -> X++
			B -> X--
			L -> Direction = north
			R -> Direction = south
	west 	F -> X--
			B -> X++
			L -> Direction = south
			R -> Direction = north
	north 	F -> Y++
			B -> Y--
			L -> Direction = west
			R -> Direction = east
	south 	F -> Y--
			B -> Y++
			L -> Direction = east
			R -> Direction = west	
*/ 
func (rover *Rover) Move(command string) { // FLFFFRFLB
	for _, char := range command {
		// fmt.Println(r)
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
			if rover.Direction == "east" {
				rover.Direction = "north"
			} else if rover.Direction == "west" {
				rover.Direction = "south"
			} else if rover.Direction == "north" {
				rover.Direction = "west"
			} else if rover.Direction == "south" {
				rover.Direction = "east"
			}
		}
		case "R": {
			if rover.Direction == "east" {
				rover.Direction = "south"
			} else if rover.Direction == "west" {
				rover.Direction = "north"
			} else if rover.Direction == "north" {
				rover.Direction = "east"
			} else if rover.Direction == "south" {
				rover.Direction = "west"
			}
		}

		} // switch ends
	}
}