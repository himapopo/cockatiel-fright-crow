package cockatiel

import (
	"bytes"
	"cockatiel-fright-crow/game"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

const (
	DefaultX    = 784
	DefautlY    = 745
	ImageWidth  = 0.08
	ImageHeight = 0.08
)

type Cockatiel struct {
	image       *ebiten.Image
	imageWidth  float64
	imageHeight float64

	// endPosition
	epx float64
	epy float64

	// currentPosition
	cpx float64
	cpy float64
}

func NewCockatiel() *Cockatiel {

	// オカメ画像ファイル
	// b, err := os.ReadFile("./assets/img/えんぴつオカメ.png")
	// if err != nil {
	// 	e := xerrors.Errorf("error: %w", err)
	// 	log.Fatalf("%+v\n", e)
	// }

	i, err := png.Decode(bytes.NewReader(imageByte))
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	cockatielImage := ebiten.NewImageFromImage(i)

	return &Cockatiel{
		image:       cockatielImage,
		imageWidth:  0.08,
		imageHeight: 0.08,
		epx:         0,
		epy:         0,
		cpx:         0,
		cpy:         0,
	}
}

func (c *Cockatiel) Reset() {
	c.epx = 0
	c.epy = 0
}

func (c *Cockatiel) ImageDraw(screen *ebiten.Image, g *game.Game) {
	c.cpx, c.cpy = CurrentPosition()

	// ゲームオーバーした場所を保存
	if g.State.State == "end" && c.epx == 0 && c.epy == 0 {
		c.epx = c.cpx
		c.epy = c.cpy
	}

	// ゲームオーバーした場所が登録されている場合は固定
	if c.epx != 0 && c.epy != 0 {
		c.cpx = c.epx
		c.cpy = c.epy
	}

	oop := &ebiten.DrawImageOptions{}

	// オカメ画像の大きさ
	oop.GeoM.Scale(c.imageWidth, c.imageHeight)

	// オカメ画像の位置
	oop.GeoM.Translate(c.cpx, c.cpy)

	// オカメ画像描画
	screen.DrawImage(c.image, oop)
}

// オカメの現在地
func CurrentPosition() (float64, float64) {

	cursorX, cursorY := ebiten.CursorPosition()

	var cpx float64 = 0
	var cpy float64 = 0

	cpx = float64(cursorX)
	cpy = float64(cursorY)

	// --------画面からはみ出ない設定---------

	if cpx <= 0 {
		cpx = 0
	}

	if cpx >= (game.ScreenWidth - (DefaultX * ImageWidth)) {
		cpx = game.ScreenWidth - (DefaultX * ImageWidth)
	}

	if cpy <= 0 {
		cpy = 0
	}

	if cpy >= (game.ScreenHeight - (DefautlY * ImageHeight)) {
		cpy = game.ScreenHeight - (DefautlY * ImageHeight)
	}

	// --------画面からはみ出ない設定---------

	return cpx, cpy
}
