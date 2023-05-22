package main

import (
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/keyboard/keyboard"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("images/starship.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

var posX *float64
var posY *float64
var rot *int

func moveShip(rotation int, positionY float64, positionX float64) {
	*rot = rotation
	*posY = positionY
	*posX = positionX
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for _, p := range g.keys {
		_, ok := keyboard.KeyRect(p)

		if ok && p == ebiten.KeyArrowDown {
			if *rot == 0 {
				*posX += float64(img.Bounds().Max.X)
			}
			moveShip(180, *posY+5, *posX)
		}

		if ok && p == ebiten.KeyArrowRight {
			if *rot == -90 {
				*posY -= float64(img.Bounds().Max.Y)
			}
			moveShip(90, *posY, *posX+5)
		}

		if ok && p == ebiten.KeyArrowLeft {
			if *rot == 90 {
				*posY += float64(img.Bounds().Max.Y)
			}
			moveShip(-90, *posY, *posX-5)
		}

		if ok && p == ebiten.KeyArrowUp {
			if *rot == 180 {
				*posX -= float64(img.Bounds().Max.X)
			}
			moveShip(0, *posY-5, *posX)
		}
	}

	op.GeoM.Rotate(float64(*rot%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(*posX, *posY)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	posX = new(float64)
	posY = new(float64)
	rot = new(int)

	*rot = 0
	*posX = 10
	*posY = 10

	ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowSize(ebiten.ScreenSizeInFullscreen())

	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
