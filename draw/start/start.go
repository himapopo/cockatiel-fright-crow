package start

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

var (
	startImage *ebiten.Image
)

func Init() {
	f, err := os.Open("./assets/img/スタート.jpg")
	defer f.Close()
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	i, err := jpeg.Decode(f)
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	startImage = ebiten.NewImageFromImage(i)
}

func ImageDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// トップ画像の大きさ
	op.GeoM.Scale(0.95, 0.88)

	screen.DrawImage(startImage, op)
}
