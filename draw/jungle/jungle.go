package jungle

import (
	"bytes"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

type Jungle struct {
	image       *ebiten.Image
	imageWidth  float64
	imageHeight float64
	tx          float64
	ty          float64
}

func NewJungle() *Jungle {
	i, err := png.Decode(bytes.NewReader(imageByte))
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	jungleImage := ebiten.NewImageFromImage(i)

	return &Jungle{
		image:       jungleImage,
		imageWidth:  1.5,
		imageHeight: 1.5,
		tx:          -40,
		ty:          -40,
	}
}

func (j *Jungle) ImageDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// 背景画像の大きさ
	op.GeoM.Scale(j.imageWidth, j.imageHeight)

	// 背景画像の位置
	op.GeoM.Translate(j.tx, j.ty)

	screen.DrawImage(j.image, op)
}
