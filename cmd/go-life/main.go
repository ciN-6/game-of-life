package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"game-of-life/pkg/life"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth   = 1280
	screenHeight  = 1200
	characterSize = 10
)

type GridTemplate struct {
	Name   string
	Width  int
	Height int
}

var gridTemplates = []GridTemplate{
	{"Small (32x20)", 32, 20},
	{"Medium (64x40)", 64, 40},
	{"Large (80x50)", 80, 50},
}

type GameState int

const (
	StateGridSizeSelection GameState = iota
	StateSeedSelection
	StateRunning
	StatePaused
)

type Popup struct {
	x, y int
	text string
}

type Button struct {
	x, y, w, h int
	label      string
	action     func()
}

func (b *Button) Clicked(x, y int) bool {
	return x >= b.x && x <= b.x+b.w && y >= b.y && y <= b.y+b.h
}

type Game struct {
	grid       *life.Board
	state      GameState
	seedCount  int
	speed      int // delay in ticks (higher = slower)
	tickCount  int
	buttons    []*Button
	gridWidth  int
	gridHeight int
	uiY2       int
	popup      *Popup
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		state: StateGridSizeSelection,
		speed: 10,
	}
}

func (g *Game) initButtons() {
	gridBottom := g.gridHeight*characterSize + 20
	uiY1 := gridBottom
	g.uiY2 = gridBottom + 60
	uiY2 := g.uiY2

	g.buttons = []*Button{
		{10, uiY1, 80, 40, "Start", func() {
			if g.state == StateSeedSelection && g.seedCount > 0 {
				g.state = StateRunning
			} else if g.state == StatePaused {
				g.state = StateRunning
			}
		}},
		{100, uiY1, 80, 40, "Pause", func() {
			if g.state == StateRunning {
				g.state = StatePaused
			}
		}},
		{190, uiY1, 80, 40, "Faster", func() {
			if g.speed > 1 {
				g.speed--
			}
		}},
		{280, uiY1, 80, 40, "Slower", func() {
			g.speed++
		}},
		{370, uiY1, 80, 40, "Reset", func() {
			g.grid.Clear()
			g.state = StateSeedSelection
			g.seedCount = 0
		}},
		{550, uiY1, 80, 40, "Step", func() {
			g.grid.Step()
		}},
		{460, uiY1, 80, 40, "Clear", func() {
			g.grid.Clear()
			g.seedCount = 0
		}},
		{10, uiY2, 80, 40, "Back", func() {
			g.grid.Undo()
		}},
		{460, uiY2, 110, 40, "Change Size", func() {
			g.state = StateGridSizeSelection
		}},
		{340, uiY2, 80, 40, "Random", func() {
			g.grid.Clear()
			g.seedCount = 0
			for y := 0; y < g.gridHeight; y++ {
				for x := 0; x < g.gridWidth; x++ {
					density := 0.1 + rand.Float32()*(0.25-0.1)
					if rand.Float32() < density { // 10% to 25% chance of being alive
						g.grid.SetAlive(x, y, true)
						g.seedCount++
					}
				}
			}
		}},
	}
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.popup = nil // Close popup on left click
		if g.state == StateGridSizeSelection {
			for i, t := range gridTemplates {
				// Display buttons for templates: x=100, y=100 + i*60
				if mx >= 100 && mx <= 300 && my >= 100+i*60 && my <= 140+i*60 {
					g.gridWidth = t.Width
					g.gridHeight = t.Height
					g.grid = life.NewGrid(t.Width, t.Height)
					g.state = StateSeedSelection
					g.initButtons()
					return nil
				}
			}
			return nil
		}

		// Check buttons first
		for _, b := range g.buttons {
			if b.Clicked(mx, my) {
				b.action()
				return nil
			}
		}

		// Handle grid interaction
		gx, gy := mx/characterSize, my/characterSize
		if gx >= 0 && gx < g.gridWidth && gy >= 0 && gy < g.gridHeight {
			if !g.grid.GetAlive(gx, gy) {
				g.grid.SetAlive(gx, gy, true)
				if g.state == StateSeedSelection {
					g.seedCount++
				}
			} else {
				g.grid.SetAlive(gx, gy, false)
				if g.state == StateSeedSelection && g.seedCount > 0 {
					g.seedCount--
				}
			}
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.popup = nil
		gx, gy := mx/characterSize, my/characterSize
		if gx >= 0 && gx < g.gridWidth && gy >= 0 && gy < g.gridHeight {
			cell := g.grid.GetCell(gx, gy)
			if cell != nil {
				if cell.Character == nil {
					g.popup = &Popup{
						x:    mx,
						y:    my,
						text: fmt.Sprintf("Death cell\nDeathCount: %d", cell.DeathCount),
					}
				} else {
					g.popup = &Popup{
						x: mx,
						y: my,
						text: fmt.Sprintf("ID: %d\nType: %T\nDeathCount: %d",
							cell.Character.GetID(), cell.Character, cell.DeathCount),
					}
				}
			}
		}
	}

	if g.state == StateRunning {
		g.tickCount++
		if g.tickCount >= g.speed {
			g.grid.Step()
			g.tickCount = 0
		}
	}

	return nil
}

func (g *Game) drawPopup(screen *ebiten.Image) {
	if g.popup != nil {
		// Calculate size based on lines (approximate)
		lines := 0
		for _, char := range g.popup.text {
			if char == '\n' {
				lines++
			}
		}
		lines++ // add one for the first line

		width, height := 150, lines*15+10
		vector.DrawFilledRect(screen, float32(g.popup.x), float32(g.popup.y), float32(width), float32(height), color.RGBA{50, 50, 50, 200}, false)
		ebitenutil.DebugPrintAt(screen, g.popup.text, g.popup.x+5, g.popup.y+5)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})

	if g.state == StateGridSizeSelection {
		ebitenutil.DebugPrintAt(screen, "Select Grid Size:", 100, 50)
		for i, t := range gridTemplates {
			// Draw selection "combobox" elements
			vector.DrawFilledRect(screen, 100, float32(100+i*60), 200, 40, color.RGBA{80, 80, 80, 255}, false)
			ebitenutil.DebugPrintAt(screen, t.Name, 120, 110+i*60)
		}
		return
	}

	// Draw Grid (Ends at y=400)
	for y := 0; y < g.gridHeight; y++ {
		for x := 0; x < g.gridWidth; x++ {
			cell := g.grid.Cells[y*g.gridWidth+x]

			var c color.RGBA
			if cell.Character == nil {
				c = color.RGBA{0, 0, 0, 255} // Black for empty
			} else {
				col := cell.Character.GetColor()
				c = color.RGBA{col.R, col.G, col.B, col.A}
			}

			vector.DrawFilledRect(screen, float32(x*characterSize), float32(y*characterSize), float32(characterSize-1), float32(characterSize-1), c, false)
		}
	}

	// Draw Buttons
	for _, b := range g.buttons {
		vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.w), float32(b.h), color.RGBA{100, 100, 100, 255}, false)
		ebitenutil.DebugPrintAt(screen, b.label, b.x+10, b.y+10)
	}

	// Draw Info
	status := ""
	switch g.state {
	case StateSeedSelection:
		status = fmt.Sprintf("Seed Selection: %d", g.seedCount)
	case StateRunning:
		status = fmt.Sprintf("Running (Speed: %d)", g.speed)
	case StatePaused:
		status = "Paused"
	}
	// Align horizontally between Clear/Random button area (approx x=120-450)
	ebitenutil.DebugPrintAt(screen, status, 10, g.uiY2+40)

	livecharacters := fmt.Sprintf("Alive: %d", g.grid.CountAlive())
	ebitenutil.DebugPrintAt(screen, livecharacters, 150, g.uiY2+40)

	g.drawPopup(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Calculate required height based on grid height (each character is 10px, plus 200px for buttons)
	reqWidth := g.gridWidth * characterSize
	if reqWidth < 640 {
		reqWidth = 640
	}
	reqHeight := (g.gridHeight * characterSize) + 200
	if reqHeight < 600 {
		reqHeight = 600
	}
	return reqWidth, reqHeight
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go-Life Simulation")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
