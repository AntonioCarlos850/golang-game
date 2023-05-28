package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math"
	"time"

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

var posX *float64 = new(float64)
var posY *float64 = new(float64)
var rot *int = new(int)
var moveSpeed float64 = 4
var scrWidth int = 640
var scrHeight int = 480
var boost int = 100
var boostTime time.Time

func moveShip(rotation int, positionY float64, positionX float64) {
	if positionY < float64(scrHeight) && positionY > 0 {
		*posY = positionY
	}

	if positionX < float64(scrWidth) && positionX > 0 {
		*posX = positionX
	}

	*rot = rotation
}

func verifyShipMovement(keys *[]ebiten.Key) {
	var space bool = false
	for _, p := range *keys {
		_, ok := keyboard.KeyRect(p)

		if !ok {
			continue
		}

		if p == ebiten.KeySpace {
			boostTime = time.Now()

			if boost < 2 {
				continue
			}

			space = true
			boost--
			moveSpeed = 6
			moveShip(*rot, *posY, *posX)
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

	if space == false {
		if boost < 100 && time.Now().Sub(boostTime) > 2*time.Second {
			boost++
		}

		moveSpeed = 4
	}
}

func (g *Game) Update() error {
	var keys []ebiten.Key
	keys = inpututil.AppendPressedKeys(keys[:0])
	verifyShipMovement(&keys)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var recharge float64 = 0
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(float64(*rot%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(*posX, *posY)
	screen.DrawImage(img, op)

	if time.Now().Sub(boostTime) < 2*time.Second {
		recharge = 2 - time.Now().Sub(boostTime).Seconds()
	}

	ebitenutil.DebugPrint(screen, fmt.Sprint("Boost qtd: ", boost))
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Boost recharge time: ", fmt.Sprintf("%.2f", recharge)), 0, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	*rot = 0
	*posX = 10
	*posY = 10

	ebiten.SetWindowSize(scrWidth, scrHeight)
	// ebiten.SetWindowSize(ebiten.ScreenSizeInFullscreen())

	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
