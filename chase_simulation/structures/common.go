package structures

import "math"

type Point struct {
	X int
	Y int
}

func (p *Point) DistanceTo(other Point) int {
	return int(math.Abs(float64(p.X-other.X)) + math.Abs(float64(p.Y-other.Y)))
}

type Cell struct {
	object GridObject
}
