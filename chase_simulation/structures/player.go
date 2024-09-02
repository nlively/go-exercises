package structures

type Player struct {
	position        Point
	visibilityRange int
	speed           int // cells per turn

	cellMap [][]GridObjectType
}

func (p *Player) Position() Point {
	return p.position
}

func (p *Player) Speed() int {
	return p.speed
}

func (p *Player) Width() int {
	return 1
}

func (p *Player) Height() int {
	return 1
}

func (p *Player) CanSeePlayer(other *Player) bool {
	return p.position.DistanceTo(other.position) <= p.visibilityRange
}
