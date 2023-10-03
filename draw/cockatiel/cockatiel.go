package cockatiel

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

var (

	// オカメ画像
	okameImage *ebiten.Image

	okameImageSizeX = 0.15
	okameImageSizeY = 0.15
)

func Init() {
	// オカメ画像ファイル
	f, err := os.Open("./assets/img/えんぴつオカメ.png")
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
	okameImage = ebiten.NewImageFromImage(i)
}

func ImageDraw(screen *ebiten.Image) {
	// カーソルの位置
	cursorX, cursorY := ebiten.CursorPosition()

	oop := &ebiten.DrawImageOptions{}

	// オカメ画像の大きさ
	oop.GeoM.Scale(okameImageSizeX, okameImageSizeY)

	// オカメ画像の位置
	oop.GeoM.Translate(float64(cursorX), float64(cursorY))

	// オカメ画像描画
	screen.DrawImage(okameImage, oop)
}
