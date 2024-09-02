package main

import (
	"math"
)

type Player struct {
	CellMap         [][]int
	Position        PrecisePoint
	Orientation     float64
	Speed           int // cells per turn
	VisibilityRange int

	grid *Grid
}

func NewPlayer(grid *Grid) *Player {
	p := &Player{grid: grid}
	p.CellMap = make([][]int, grid.Width)
	for i := range p.CellMap {
		p.CellMap[i] = make([]int, grid.Height)
	}

	// Mark cells as unknown
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			p.CellMap[x][y] = CellTypeUnknown
		}
	}

	return p
}

func (p *Player) UpdateDiscoveredCells() {
	// Update cells that player can see
	for x := 0; x < p.grid.Width; x++ {
		for y := 0; y < p.grid.Height; y++ {
			// Check if cell is within visibility range
			if p.Position.DistanceTo(PrecisePoint{float64(x), float64(y)}) <= float64(p.VisibilityRange) {
				p.CellMap[x][y] = p.grid.Cells[x][y]
			}
		}
	}
}

func (p *Player) SetOrientation(degrees float64) {
	p.Orientation = degrees
}

func (p *Player) Rotate(degrees float64) {
	p.Orientation = math.Mod(p.Orientation+degrees, 360)
}

func (p *Player) Move() {
	for {

		// Move player in direction
		newX := p.Position.X + math.Cos(p.Orientation)*float64(p.Speed)
		newY := p.Position.Y + math.Sin(p.Orientation)*float64(p.Speed)

		// If new position is out of bounds, rotate player and try again
		if newX < 0 || newX >= float64(p.grid.Width) || newY < 0 || newY >= float64(p.grid.Height) {
			p.Rotate(10)
			continue
		}

		// If new position collides with a filled cell, rotate player and try again
		if p.grid.Cells[int(newX)][int(newY)] == CellTypeFilled {
			p.Rotate(10)
			continue
		}

		p.Position.X = newX
		p.Position.Y = newY

		break
	}

	// If all cells in player's field of view are known, change orientation
	unknownCells := 0
	for x := 0; x < p.grid.Width; x++ {
		for y := 0; y < p.grid.Height; y++ {
			if p.Position.DistanceTo(PrecisePoint{float64(x), float64(y)}) <= float64(p.VisibilityRange) && p.CellMap[x][y] == CellTypeUnknown {
				unknownCells++
			}
		}
	}

	if unknownCells == 0 {
		// rotate player 45 degrees
		p.Rotate(45)
	}

	// Update discovered cells
	p.UpdateDiscoveredCells()
}
