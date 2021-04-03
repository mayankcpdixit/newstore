package landscape

type Landscape struct {
	Obstacles []Coordinate // list of coordinates
}

func NewLandscape(obstacles []Coordinate) *Landscape {
	return &Landscape{Obstacles: obstacles}
}

func (l *Landscape) HasObstacleAt(coordinate Coordinate) bool {
	for _,obs := range l.Obstacles {
		if obs.X == coordinate.X && obs.Y == coordinate.Y {
			return true
		}
	}
	return false
}