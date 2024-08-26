package main

import "math/rand"

const (
	GRID_WIDTH  = 100
	GRID_HEIGHT = 100

	LEFT_EDGE   = 0
	TOP_EDGE    = 1
	RIGHT_EDGE  = 2
	BOTTOM_EDGE = 3
)

func initializeGrid() Grid {
	grid := Grid{}
	grid.Cells = make([][]Cell, GRID_HEIGHT)
	for i := 0; i < GRID_HEIGHT; i++ {
		grid.Cells[i] = make([]Cell, GRID_WIDTH)
	}
	return grid
}

func addRandomExit(grid Grid) {
	var x, y int
	// Randomize edge
	edge := rand.Intn(4)
	switch edge {
	case LEFT_EDGE:
		x = 0
		y = rand.Intn(GRID_HEIGHT)
	case TOP_EDGE:
		x = rand.Intn(GRID_WIDTH)
		y = 0
	case RIGHT_EDGE:
		x = GRID_WIDTH - 1
		y = rand.Intn(GRID_HEIGHT)
	case BOTTOM_EDGE:
		x = rand.Intn(GRID_WIDTH)
		y = GRID_HEIGHT - 1
	}

	exit := Exit{
		position: Point{x, y},
	}

	grid.PlaceObject(&exit)
}

func addRandomRoom(grid Grid) {
	// Randomize width and height
	width := rand.Intn(8) + 1
	height := rand.Intn(8) + 1

	var position Point
	for {
		// Randomize position
		x := rand.Intn(GRID_WIDTH - width)
		y := rand.Intn(GRID_HEIGHT - height)
		position = Point{x, y}

		if grid.CanSafelyPlaceObject(Point{x, y}, width, height) {
			break
		}
	}

	room := Room{
		position: position,
		width:    width,
		height:   height,
	}

	grid.PlaceObject(&room)
	addRandomExit(grid)
	for i := 0; i < rand.Intn(5)+3; i++ {
		addRandomRoom(grid)
	}
}

func main() {
	grid := initializeGrid()

	addRandomExit(grid)
}
