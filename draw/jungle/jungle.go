package jungle

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

var (
	jungleImage      *ebiten.Image
	jungleImageSizeX = 1.5
	jungleImageSizeY = 1.5
)

func Init() {
	f, err := os.Open("./assets/img/ジャングル.png")
	defer f.Close()
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	i, err := png.Decode(f)
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	jungleImage = ebiten.NewImageFromImage(i)
}

func ImageDraw(screen *ebiten.Image) {
	jop := &ebiten.DrawImageOptions{}

	// 背景画像の大きさ
	jop.GeoM.Scale(jungleImageSizeX, jungleImageSizeY)

	// 背景画像の位置
	jop.GeoM.Translate(-40, -40)

	screen.DrawImage(jungleImage, jop)
}
