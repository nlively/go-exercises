package structures

type Exit struct {
	position Point
}

func (r *Exit) Type() GridObjectType {
	return GridObjectExit
}

func (r *Exit) Position() Point {
	return r.position
}

func (r *Exit) Width() int {
	return 1
}

func (r *Exit) Height() int {
	return 1
}
