package structures

import (
	"fmt"

	"math/rand"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	LEFT_EDGE   = 0
	TOP_EDGE    = 1
	RIGHT_EDGE  = 2
	BOTTOM_EDGE = 3
)

type Grid struct {
	Cells   [][]Cell
	exit    *Exit
	objects []GridObject
	width   int
	height  int

	attacker *Player
	defender *Player
}

func NewGrid(width int, height int) *Grid {
	g := &Grid{width: width, height: height}
	g.Cells = make([][]Cell, width)
	for i := 0; i < width; i++ {
		g.Cells[i] = make([]Cell, height)
	}
	return g
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

	g.objects = append(g.objects, obj)
}

func (g *Grid) PlacePlayers() {
	var attackerPosition, defenderPosition Point

	// Randomly place attacker
	for {
		attackerPosition = Point{rand.Intn(g.width), rand.Intn(g.height)}
		if g.CanSafelyPlaceObject(attackerPosition, 1, 1) {
			break
		}
	}

	// Randomly place defender, ensuring that defender is not within 5 cells of attacker
	for {
		defenderPosition = Point{rand.Intn(g.width), rand.Intn(g.height)}
		if g.CanSafelyPlaceObject(defenderPosition, 1, 1) &&
			attackerPosition.DistanceTo(defenderPosition) > 5 {
			break
		}
	}

	g.attacker = &Player{position: attackerPosition}
	g.defender = &Player{position: defenderPosition}
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
	for i := 0; i < g.width; i++ {
		if g.exit != nil && g.exit.position.X == i && g.exit.position.Y == 0 {
			fmt.Printf(" ") // Gap for the exit
		} else {
			fmt.Printf("—")
		}
	}
	fmt.Printf("+\n")

	// Middle with grid contents
	for y := 0; y < g.height; y++ {
		if g.exit != nil && g.exit.position.X == 0 && g.exit.position.Y == y {
			fmt.Printf(" ") // Gap for the exit
		} else {
			fmt.Printf("|") // Left border
		}

		for x := 0; x < g.width; x++ {
			if g.Cells[x][y].object != nil {
				switch g.Cells[x][y].object.Type() {
				case GridObjectRoom:
					fmt.Printf("R") // Represent rooms with 'R'
				case GridObjectExit:
					fmt.Printf("X") // Represent exits with 'E'
				default:
					fmt.Printf("?") // Represent unknown objects with '?'
				}
			} else {
				fmt.Printf(" ") // Blank space for empty cells
			}
		}

		if g.exit != nil && g.exit.position.X == g.width-1 && g.exit.position.Y == y {
			fmt.Printf(" ") // Gap for the exit
		} else {
			fmt.Printf("|\n") // Right border
		}
	}

	// Bottom border
	fmt.Printf("+")
	for i := 0; i < g.width; i++ {
		if g.exit != nil && g.exit.position.X == i && g.exit.position.Y == g.height-1 {
			fmt.Printf(" ") // Gap for the exit
		} else {
			fmt.Printf("—")
		}
	}
	fmt.Printf("+\n")
}

func (g *Grid) AddRandomExit() {
	var x, y int
	// Randomize edge
	edge := rand.Intn(4)
	switch edge {
	case LEFT_EDGE:
		x = 0
		y = rand.Intn(g.height)
	case TOP_EDGE:
		x = rand.Intn(g.width)
		y = 0
	case RIGHT_EDGE:
		x = g.width - 1
		y = rand.Intn(g.height)
	case BOTTOM_EDGE:
		x = rand.Intn(g.width)
		y = g.height - 1
	}

	exit := Exit{
		position: Point{x, y},
	}
	g.PlaceObject(&exit)
}

func (g *Grid) AddRandomRoom() {
	var position Point
	var width, height int

	for {
		// Randomize width and height
		width = rand.Intn(4) + 4
		height = rand.Intn(4) + 4

		// Randomize position
		x := rand.Intn(g.width-width-2) + 2
		y := rand.Intn(g.height-height-2) + 2
		position = Point{x, y}

		if g.CanSafelyPlaceObject(Point{x, y}, width, height) {
			break
		}
	}

	room := Room{
		position: position,
		width:    width,
		height:   height,
	}

	g.PlaceObject(&room)
}

func (g *Grid) screenToGridRatio(screen *ebiten.Image) (float32, float32) {
	// Figure out ratio of screen to grid
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	ratioX := float32(screenWidth) / float32(g.width)
	ratioY := float32(screenHeight) / float32(g.height)
	return ratioX, ratioY
}

func (g *Grid) translatePointToScreen(point Point, screen *ebiten.Image) (int, int) {
	ratioX, ratioY := g.screenToGridRatio(screen)

	// Translate point to screen
	x := float32(point.X) * ratioX
	y := float32(point.Y) * ratioY

	return int(x), int(y)
}

func (g *Grid) translateDimensionsToScreen(width int, height int, screen *ebiten.Image) (int, int) {
	ratioX, ratioY := g.screenToGridRatio(screen)

	// Translate dimensions to screen
	x := float32(width) * ratioX
	y := float32(height) * ratioY

	return int(x), int(y)
}

func (g *Grid) RenderImage(screen *ebiten.Image) {
	var x, y int
	var gridPath vector.Path
	width, height := g.translateDimensionsToScreen(g.width, 1, screen)

	// Traverse borders and draw lines for all contiguous points, leaving a gap for the exit

	// Top border
	x, y = g.translatePointToScreen(Point{0, 0}, screen)
	gridPath.MoveTo(float32(x), float32(y))
	gridPath.LineTo(float32(x+width), float32(y))

	// Right border
	x, y = g.translatePointToScreen(Point{g.width - 1, 0}, screen)
	gridPath.LineTo(float32(x+width), float32(y+height))

	// Bottom border
	x, y = g.translatePointToScreen(Point{0, g.height - 1}, screen)
	gridPath.LineTo(float32(x), float32(y+height))

	// Left border
	x, y = g.translatePointToScreen(Point{0, 0}, screen)
	gridPath.LineTo(float32(x), float32(y))

	gridPath.Close()

	roomColor := color.RGBA{0x00, 0x00, 0xff, 0xff}
	exitColor := color.RGBA{0x00, 0xff, 0xff, 0xff}
	unknownColor := color.RGBA{0x99, 0x99, 0x99, 0xff}

	for _, object := range g.objects {
		rx, ry := g.translatePointToScreen(object.Position(), screen)
		rw, rh := g.translateDimensionsToScreen(object.Width(), object.Height(), screen)
		var color color.Color

		switch object.Type() {
		case GridObjectRoom:
			color = roomColor
		case GridObjectExit:
			color = exitColor
		default:
			color = unknownColor
		}

		vector.DrawFilledRect(screen, float32(rx), float32(ry), float32(rw), float32(rh), color, true)
	}

	// Place attacker and defender
	attackerColor := color.RGBA{0xff, 0x00, 0x00, 0xff}
	defenderColor := color.RGBA{0x00, 0xff, 0x00, 0xff}

	g.RenderPlayer(screen, g.attacker, attackerColor)
	g.RenderPlayer(screen, g.defender, defenderColor)
}

func (g *Grid) RenderPlayer(screen *ebiten.Image, player *Player, color color.Color) {
	rx, ry := g.translatePointToScreen(player.Position(), screen)
	rw, rh := g.translateDimensionsToScreen(1, 1, screen)

	vector.DrawFilledRect(screen, float32(rx), float32(ry), float32(rw), float32(rh), color, true)
}

/*

	Attacker and Defender algorithm

	- if attacker can see defender
		- set move course = defender position
		- re-evaluate move course each turn while defender is visible
	- else
		- set move course = random position
		- hold move course until defender is visible or collision is detected

	- if defender can see attacker
		- if defender can see exit
			- if defender is closer to exit than attacker is
				- set move course = exit position
				- re-evaluate move course each turn while attacker is visible
			- else
				- set move course = away from attacker
				- re-evaluate move course each turn while attacker is visible
		- else
			- set move course = away from attacker
			- re-evaluate move course each turn while attacker is visible

	- else
		- set move course = random position
		- hold move course until attacker is visible or collision is detected

  Area navigation algorithm

	- if move course is set
		- move towards move course at speed
		- if collision is detected
			- adjust path as needed, keeping as close to course as possible

*/

func (g *Grid) AdvancePlayers() {
	// Can attacker see defender?
	if g.attacker.CanSeePlayer(g.defender) {
		// determine target move location in the direction of defender, factoring in obstacles and edges

	}

	// Can defender see attacker?
	if g.defender.CanSeePlayer(g.attacker) {
		// Move defender away from attacker
	}
}
