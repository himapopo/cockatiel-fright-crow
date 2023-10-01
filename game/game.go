package game

import (
	"cockatiel-fright-crow/draw/cockatiel"
	"cockatiel-fright-crow/draw/crow"
	"cockatiel-fright-crow/draw/jungle"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	// スクリーンサイズ
	ScreenWidth  = 320 // 960
	ScreenHeight = 240 // 720
)

var (
	elapsedTime = 1
)

type Game struct{}

func (g *Game) Update() error {
	elapsedTime++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// 背景画像描画
	jungle.ImageDraw(screen)

	// オカメ画像描画
	cockatiel.ImageDraw(screen)

	// カラス画像描画
	crow.ImageDraw(screen, elapsedTime)

	ebitenutil.DebugPrint(screen, strconv.Itoa(elapsedTime/60))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth * 3, ScreenHeight * 3
}
