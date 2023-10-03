package game

import (
	"cockatiel-fright-crow/draw/cockatiel"
	"cockatiel-fright-crow/draw/crow"
	"cockatiel-fright-crow/draw/jungle"
	"cockatiel-fright-crow/draw/start"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

type Game struct{}

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
		crow.ImageDraw(screen, elapsedTime)
	}

	if !RunGame {
		// スタート画面画像描画
		start.ImageDraw(screen)
	}

	ebitenutil.DebugPrint(screen, strconv.Itoa(elapsedTime/60))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth * 3, ScreenHeight * 3
}
