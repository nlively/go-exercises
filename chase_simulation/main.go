package main

import (
	"log"
	"math/rand"

	"github.com/nlively/go-exercises/chase_simulation/structures"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GRID_WIDTH  = 60
	GRID_HEIGHT = 40
)

type Game struct {
	grid *structures.Grid
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")

	g.grid.RenderImage(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	grid := structures.NewGrid(GRID_WIDTH, GRID_HEIGHT)

	grid.AddRandomExit()
	for i := 0; i < rand.Intn(5)+3; i++ {
		grid.AddRandomRoom()
	}
	grid.PlacePlayers()

	game := &Game{grid: grid}

	ebiten.SetTPS(2) // game loop advances twice a second
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
