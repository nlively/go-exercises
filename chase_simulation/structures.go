package main

import "fmt"

type GridObjectType int

const (
	GridObjectEmpty GridObjectType = iota
	GridObjectRoom
	GridObjectHallway
	GridObjectObstacle
	GridObjectExit
)

type Point struct {
	X int
	Y int
}

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

type GridObject interface {
	Type() GridObjectType
	Position() Point
	Width() int
	Height() int
}

type Cell struct {
	object GridObject
}

type Grid struct {
	Cells [][]Cell
	exit  *Exit
}

func (g *Grid) PlaceObject(obj GridObject) {
	if obj.Type() == GridObjectExit {
		if g.exit != nil {
			panic("Exit already exists")
		}
		g.exit = obj.(*Exit)
	}
	at := obj.Position()

	for x := at.X; x < obj.Width()+at.X; x++ {
		for y := at.Y; y < obj.Height()+at.Y; y++ {
			g.Cells[x][y].object = obj
		}
	}
}

func (g *Grid) CanSafelyPlaceObject(at Point, width int, height int) bool {
	for x := at.X; x < width+at.X; x++ {
		for y := at.Y; y < height+at.Y; y++ {
			if g.Cells[x][y].object != nil {
				return false
			}
		}
	}
	return true
}

func (g *Grid) Render() {
	// Top border
	fmt.Printf("+")
	for i := 0; i < GridWidth; i++ {
		if exit != nil && exit.position.x == i && exit.position.y == 0 {
			fmt.Printf(" ") // Gap for the exit
		} else {
			fmt.Printf("-")
		}
	}
	fmt.Printf("+\n")

	// Middle with grid contents
	for y := 0; y < GridHeight; y++ {
		if exit != nil && exit.position.x == 0 && exit.position.y == y {
			fmt.Printf(" ") // Gap for the exit
		} else {
			fmt.Printf("|") // Left border
		}

		for x := 0; x < GridWidth; x++ {
			if grid[x][y].object != nil {
				fmt.Printf("R") // Represent rooms with 'R'
			} else {
				fmt.Printf(" ") // Blank space for empty cells
			}
		}

		if exit != nil && exit.position.x == GridWidth-1 && exit.position.y == y {
			fmt.Printf(" ") // Gap for the exit
		} else {
			fmt.Printf("|\n") // Right border
		}
	}

	// Bottom border
	fmt.Printf("+")
	for i := 0; i < GridWidth; i++ {
		if exit != nil && exit.position.x == i && exit.position.y == GridHeight-1 {
			fmt.Printf(" ") // Gap for the exit
		} else {
			fmt.Printf("-")
		}
	}
	fmt.Printf("+\n")
}

type Player struct {
	position Point
}
