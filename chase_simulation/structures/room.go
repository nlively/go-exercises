package structures

type Room struct {
	position Point
	width    int
	height   int
}

func (r *Room) Type() GridObjectType {
	return GridObjectRoom
}

func (r *Room) Position() Point {
	return r.position
}

func (r *Room) Width() int {
	return r.width
}

func (r *Room) Height() int {
	return r.height
}
