package main

import (
	"image/color"
	"log"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	GRID_WIDTH  = 80
	GRID_HEIGHT = 60
)

type Game struct {
	Grid   *Grid
	Player *Player
}

func (g *Game) Update() error {
	g.Player.Move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// New image for grid
	gridImage := ebiten.NewImage(320, 240)
	playerImage := ebiten.NewImage(320, 240)

	// Draw grid
	g.DrawGrid(gridImage)
	g.DrawPlayerView(playerImage)

	gridPosition := &ebiten.GeoM{}
	gridPosition.Translate(240, 40)

	playerPosition := &ebiten.GeoM{}
	playerPosition.Translate(240, 320)

	screen.DrawImage(gridImage, &ebiten.DrawImageOptions{
		GeoM: *gridPosition,
	})
	screen.DrawImage(playerImage, &ebiten.DrawImageOptions{
		GeoM: *playerPosition,
	})
}

func (g *Game) DrawGrid(screen *ebiten.Image) {
	grid := g.Grid
	// Draw grid
	screen.Fill(color.RGBA{0, 0, 0xff, 0xff})

	roomColor := color.RGBA{0, 0xff, 0, 0xff}

	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if grid.Cells[x][y] == CellTypeFilled {
				tx, ty := grid.TranslateCoordinatesToScreen(Point{x, y}, screen)
				tw, th := grid.TranslateCoordinatesToScreen(Point{1, 1}, screen)
				vector.DrawFilledRect(screen, float32(tx), float32(ty), float32(tw), float32(th), roomColor, true)
			}
		}
	}

	// Draw player as triangle, pointed in direction of orientation
	player := g.Player
	tx, ty := grid.TranslateCoordinatesToScreen(player.Position.ToPoint(), screen)
	tw, th := grid.TranslateCoordinatesToScreen(Point{1, 1}, screen)

	vector.DrawFilledRect(screen, float32(tx), float32(ty), float32(tw), float32(th), color.RGBA{0xff, 0, 0, 0xff}, true)

	// vertices := []ebiten.Vertex{
	// 	{
	// 		DstX:   float32(tx),
	// 		DstY:   float32(ty),
	// 		ColorR: 1, ColorG: 0, ColorB: 0, ColorA: 1, // Red
	// 	},
	// 	{
	// 		DstX:   float32(tx + tw),
	// 		DstY:   float32(ty),
	// 		ColorR: 1, ColorG: 0, ColorB: 0, ColorA: 1, // Red
	// 	},
	// 	{
	// 		DstX:   float32(tx + tw/2),
	// 		DstY:   float32(ty + th),
	// 		ColorR: 1, ColorG: 0, ColorB: 0, ColorA: 1, // Red
	// 	},
	// }

	// indices := []uint16{0, 1, 2}

	// // Create a GeoM (geometry matrix) for transformations
	// geoM := ebiten.GeoM{}
	// // geoM.Translate(-25, -25) // Move the triangle to the origin for rotation
	// geoM.Rotate(player.Orientation) // Apply rotation
	// // geoM.Translate(160, 120)  // Move the triangle back to its original position

	// // Apply the transformation matrix to the vertices
	// for i := range vertices {
	// 	x, y := geoM.Apply(float64(vertices[i].DstX), float64(vertices[i].DstY))
	// 	vertices[i].DstX = float32(x)
	// 	vertices[i].DstY = float32(y)
	// }

	// screen.DrawTriangles(vertices, indices, ebiten.NewImage(tw, th), nil)

}

func (g *Game) DrawPlayerView(screen *ebiten.Image) {
	p := g.Player
	unknownColor := color.RGBA{0x80, 0x80, 0x80, 0xff}
	// Draw player
	screen.Fill(unknownColor)

	emptyColor := color.RGBA{0xcc, 0xcc, 0xcc, 0xff}
	filledColor := color.RGBA{0x00, 0x33, 0x00, 0xff}

	// Draw grid, revealing what is known to player
	for x := 0; x < len(p.CellMap); x++ {
		for y := 0; y < len(p.CellMap[x]); y++ {
			if p.CellMap[x][y] == CellTypeUnknown {
				continue
			}

			tx, ty := p.grid.TranslateCoordinatesToScreen(Point{x, y}, screen)
			tw, th := p.grid.TranslateCoordinatesToScreen(Point{1, 1}, screen)

			var color color.Color
			if p.CellMap[x][y] == CellTypeEmpty {
				color = emptyColor
			} else {
				color = filledColor
			}
			vector.DrawFilledRect(screen, float32(tx), float32(ty), float32(tw), float32(th), color, true)

		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	grid := NewGrid(GRID_WIDTH, GRID_HEIGHT)
	player := NewPlayer(grid)

	// Place player in random room
	for {
		x := rand.Intn(GRID_WIDTH)
		y := rand.Intn(GRID_HEIGHT)
		pos := Point{x, y}
		if grid.CanPlaceRoom(pos, 1, 1) {
			player.Position = pos.ToPrecisePoint()
			break
		}
	}

	// Randomize player's orientation
	player.Orientation = rand.Float64() * 2 * 3.14159
	player.Speed = 1
	player.VisibilityRange = 5

	game := &Game{
		Grid:   grid,
		Player: player,
	}

	ebiten.SetTPS(2) // game loop advances twice a second
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
