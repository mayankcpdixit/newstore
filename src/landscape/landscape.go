package landscape

type Landscape struct {
	Obstacles [][]int // list of coordinates
}

func Init(obstacles [][]int) *Landscape {
	l := Landscape{Obstacles: obstacles}
	return &l;
}

func (l *Landscape) HasObstacleAt(x,y int) bool {
	for _,obs := range l.Obstacles {
		if obs[0] == x && obs[1] == y {
			return true
		}
	}
	return false
}