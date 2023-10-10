package start

import (
	"bytes"
	"image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

type Start struct {
	image       *ebiten.Image
	imageWidth  float64
	imageHeight float64
}

func NewStart(imageByte []byte) *Start {
	i, err := jpeg.Decode(bytes.NewReader(imageByte))
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	startImage := ebiten.NewImageFromImage(i)

	return &Start{
		image:       startImage,
		imageWidth:  0.95,
		imageHeight: 0.88,
	}
}

func (s *Start) ImageDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// トップ画像の大きさ
	op.GeoM.Scale(s.imageWidth, s.imageHeight)

	screen.DrawImage(s.image, op)
}
