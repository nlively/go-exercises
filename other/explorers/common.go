package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CellTypeUnknown = 0
	CellTypeEmpty   = 1
	CellTypeFilled  = 2
)

type PrecisePoint struct {
	X float64
	Y float64
}

func (p *PrecisePoint) DistanceTo(other PrecisePoint) float64 {
	return math.Abs(p.X-other.X) + math.Abs(p.Y-other.Y)
}

func (p *PrecisePoint) ToPoint() Point {
	return Point{int(p.X), int(p.Y)}
}

type Point struct {
	X int
	Y int
}

func (p *Point) DistanceTo(other Point) int {
	return int(math.Abs(float64(p.X-other.X)) + math.Abs(float64(p.Y-other.Y)))
}

func (p *Point) ToPrecisePoint() PrecisePoint {
	return PrecisePoint{float64(p.X), float64(p.Y)}
}

func ScreenToGridRatio(screen *ebiten.Image, gridWidth int, gridHeight int) (float32, float32) {
	// Figure out ratio of screen to grid
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	ratioX := float32(screenWidth) / float32(gridWidth)
	ratioY := float32(screenHeight) / float32(gridHeight)
	return ratioX, ratioY
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
