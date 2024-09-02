package main

import (
	"math"
)

type Player struct {
	CellMap         [][]int
	Position        PrecisePoint
	Orientation     float64
	Speed           int
	VisibilityRange int

	grid            *Grid
	recentLocations []PrecisePoint
	maxHistory      int

	momentumFactor float64 // Higher value means more persistence in current direction
}

func NewPlayer(grid *Grid) *Player {
	p := &Player{grid: grid, maxHistory: 5, momentumFactor: 0.8} // Set momentum factor
	p.CellMap = make([][]int, grid.Width)
	for i := range p.CellMap {
		p.CellMap[i] = make([]int, grid.Height)
	}

	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			p.CellMap[x][y] = CellTypeUnknown
		}
	}

	return p
}

func (p *Player) UpdateDiscoveredCells() {
	for x := 0; x < p.grid.Width; x++ {
		for y := 0; y < p.grid.Height; y++ {
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
		newX := p.Position.X + math.Cos(p.Orientation)*float64(p.Speed)
		newY := p.Position.Y + math.Sin(p.Orientation)*float64(p.Speed)

		if newX < 0 || newX >= float64(p.grid.Width) || newY < 0 || newY >= float64(p.grid.Height) {
			p.Rotate(45) // Rotate more significantly to avoid getting stuck
			continue
		}

		if p.grid.Cells[int(newX)][int(newY)] == CellTypeFilled || p.isRecentlyVisited(newX, newY) {
			p.Rotate(45)
			continue
		}

		p.Position.X = newX
		p.Position.Y = newY
		p.recordVisit()
		break
	}

	p.UpdateDiscoveredCells()

	unknownCells := p.CountUnknownCellsInRange()

	if unknownCells == 0 || p.shouldChangeDirection() {
		p.AdjustOrientation()
	}
}

func (p *Player) AdjustOrientation() {
	// Calculate direction with the most unknown cells within visibility range, weighted by momentum
	maxUnknownCells := 0
	bestOrientation := p.Orientation

	for angle := 0.0; angle < 360; angle += 45 {
		unknownCells := 0
		tempOrientation := math.Mod(p.Orientation+angle, 360)

		for x := 0; x < p.grid.Width; x++ {
			for y := 0; y < p.grid.Height; y++ {
				if p.IsCellInDirection(x, y, tempOrientation) && p.CellMap[x][y] == CellTypeUnknown {
					unknownCells++
				}
			}
		}

		// Apply momentum factor to prefer the current direction
		if tempOrientation == p.Orientation {
			unknownCells = int(float64(unknownCells) * p.momentumFactor)
		}

		if unknownCells > maxUnknownCells {
			maxUnknownCells = unknownCells
			bestOrientation = tempOrientation
		}
	}

	p.Orientation = bestOrientation
}

func (p *Player) IsCellInDirection(x, y int, orientation float64) bool {
	dx := float64(x) - p.Position.X
	dy := float64(y) - p.Position.Y
	angleToCell := math.Atan2(dy, dx) * 180 / math.Pi

	return math.Abs(angleToCell-orientation) <= 45
}

func (p *Player) CountUnknownCellsInRange() int {
	count := 0
	for x := 0; x < p.grid.Width; x++ {
		for y := 0; y < p.grid.Height; y++ {
			if p.Position.DistanceTo(PrecisePoint{float64(x), float64(y)}) <= float64(p.VisibilityRange) && p.CellMap[x][y] == CellTypeUnknown {
				count++
			}
		}
	}
	return count
}

func (p *Player) recordVisit() {
	p.recentLocations = append(p.recentLocations, p.Position)
	if len(p.recentLocations) > p.maxHistory {
		p.recentLocations = p.recentLocations[1:]
	}
}

func (p *Player) isRecentlyVisited(x, y float64) bool {
	for _, loc := range p.recentLocations {
		if loc.X == x && loc.Y == y {
			return true
		}
	}
	return false
}

func (p *Player) shouldChangeDirection() bool {
	// Introduce a threshold for direction change based on how long the player has been moving in one direction
	return len(p.recentLocations) > p.maxHistory/2 && p.isRecentlyVisited(p.Position.X, p.Position.Y)
}
