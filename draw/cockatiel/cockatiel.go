package cockatiel

import (
	"cockatiel-fright-crow/update"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

var (

	// オカメ画像
	okameImage *ebiten.Image

	okameImageSizeX = 0.15 // 元の大きさ 745
	okameImageSizeY = 0.15 // 元の大きさ 784

	endPositionX float64 = 0
	endPositionY float64 = 0
)

func ResetEndPosition() {
	endPositionX = 0
	endPositionY = 0
}

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

func ImageDraw(screen *ebiten.Image, state *update.GameState) {

	var positionX float64 = 0
	var positionY float64 = 0
	// カーソルの位置
	cursorX, cursorY := ebiten.CursorPosition()

	if state.State == "end" && endPositionX == 0 && endPositionY == 0 {
		endPositionX = float64(cursorX)
		endPositionY = float64(cursorY)
	}

	if endPositionX != 0 && endPositionY != 0 {
		positionX = endPositionX
		positionY = endPositionY
	} else {
		positionX = float64(cursorX)
		positionY = float64(cursorY)
	}

	oop := &ebiten.DrawImageOptions{}

	// オカメ画像の大きさ
	oop.GeoM.Scale(okameImageSizeX, okameImageSizeY)

	// オカメ画像の位置
	oop.GeoM.Translate(positionX, positionY)

	// オカメ画像描画
	screen.DrawImage(okameImage, oop)
}
