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
}

var posX *float64
var posY *float64
var rot *int
var moveSpeed float64 = 2

func moveShip(rotation int, positionY float64, positionX float64) {
	*rot = rotation
	*posY = positionY
	*posX = positionX
}

func verifyShipMovement(keys *[]ebiten.Key) {
	for _, p := range *keys {
		_, ok := keyboard.KeyRect(p)

		if !ok {
			continue
		}

		if p == ebiten.KeyArrowDown {
			if *rot == 0 {
				*posX += float64(img.Bounds().Max.X)
			}
			moveShip(180, *posY+moveSpeed, *posX)
		}

		if p == ebiten.KeyArrowRight {
			if *rot == -90 {
				*posY -= float64(img.Bounds().Max.Y)
			}
			moveShip(90, *posY, *posX+moveSpeed)
		}

		if p == ebiten.KeyArrowLeft {
			if *rot == 90 {
				*posY += float64(img.Bounds().Max.Y)
			}
			moveShip(-90, *posY, *posX-moveSpeed)
		}

		if p == ebiten.KeyArrowUp {
			if *rot == 180 {
				*posX -= float64(img.Bounds().Max.X)
			}
			moveShip(0, *posY-moveSpeed, *posX)
		}
	}
}

func (g *Game) Update() error {
	var keys []ebiten.Key
	keys = inpututil.AppendPressedKeys(keys[:0])
	verifyShipMovement(&keys)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
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
