package structures

type GridObjectType string

const (
	GridObjectEmpty    GridObjectType = ""
	GridObjectRoom     GridObjectType = "room"
	GridObjectHallway  GridObjectType = "hallway"
	GridObjectObstacle GridObjectType = "obstacle"
	GridObjectExit     GridObjectType = "exit"
)

type GridObject interface {
	Type() GridObjectType
	Position() Point
	Width() int
	Height() int
}
