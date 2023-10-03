package game

import (
	"cockatiel-fright-crow/draw/cockatiel"
	"cockatiel-fright-crow/draw/crow"
	"cockatiel-fright-crow/draw/jungle"
	"cockatiel-fright-crow/draw/start"
	"cockatiel-fright-crow/update"
	"strconv"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	// スクリーンサイズ
	ScreenWidth  = 320 // 960
	ScreenHeight = 240 // 720
)

var (
	elapsedTime = 1

	RunGame = false
)

type Game struct {
	Score *update.GameScore
}

func (g *Game) Update() error {
	elapsedTime++

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		RunGame = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		RunGame = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// 背景画像描画
	jungle.ImageDraw(screen)

	// オカメ画像描画
	cockatiel.ImageDraw(screen)

	if RunGame {
		// カラス画像描画
		crow.ImageDraw(screen, elapsedTime, g.Score)

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(4, 12)
		op.GeoM.Scale(2.5, 2.5)

		text.DrawWithOptions(screen, "Score: "+strconv.Itoa(g.Score.Score), bitmapfont.Face, op)
	}

	if !RunGame {
		// スタート画面画像描画
		start.ImageDraw(screen)
	}

	// ebitenutil.DebugPrint(screen, strconv.Itoa(elapsedTime/60))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth * 3, ScreenHeight * 3
}
