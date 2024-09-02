package main

import (
	"fmt"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Grid struct {
	Width  int
	Height int
	Cells  [][]int
}

func NewGrid(width, height int) *Grid {
	g := &Grid{Width: width, Height: height}
	g.Cells = make([][]int, width)
	for i := range g.Cells {
		g.Cells[i] = make([]int, height)
	}

	// Fill grid with empty cells
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// fmt.Printf("Setting cell %dx%d to empty\n", x, y)
			g.Cells[x][y] = CellTypeEmpty
		}
	}

	g.PlaceRandomRooms()

	return g
}
func (g *Grid) TranslateCoordinatesToScreen(point Point, screen *ebiten.Image) (int, int) {
	ratioX, ratioY := ScreenToGridRatio(screen, g.Width, g.Height)

	// Translate point to screen
	x := float32(point.X) * ratioX
	y := float32(point.Y) * ratioY

	return int(x), int(y)
}

func (g *Grid) CanPlaceRoom(at Point, width, height int) bool {
	// Check if room can be placed at given position
	if at.X+width >= g.Width || at.Y+height >= g.Height {
		return false
	}

	for x := at.X; x < at.X+width; x++ {
		for y := at.Y; y < at.Y+height; y++ {
			if g.Cells[x][y] != CellTypeEmpty {
				return false
			}
		}
	}

	return true
}

func (g *Grid) PlaceRandomRooms() {
	// Place random rooms
	numRooms := rand.Intn(4) + 3

	for i := 0; i < numRooms; i++ {
		roomWidth := rand.Intn(5) + 3
		roomHeight := rand.Intn(5) + 3
		for {
			roomX := rand.Intn(g.Width - roomWidth)
			roomY := rand.Intn(g.Height - roomHeight)
			roomPos := Point{roomX, roomY}
			if g.CanPlaceRoom(roomPos, roomWidth, roomHeight) {
				for x := roomX; x < roomX+roomWidth; x++ {
					for y := roomY; y < roomY+roomHeight; y++ {
						g.Cells[x][y] = CellTypeFilled
					}
				}
				break
			}
		}
	}
}

func (g *Grid) Dump() {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			fmt.Printf("%d", g.Cells[x][y])
		}
		fmt.Println()
	}
}
