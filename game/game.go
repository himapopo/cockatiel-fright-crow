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
)

type Game struct {
	State *update.GameState
}

func (g *Game) Update() error {
	elapsedTime++

	// スタート画面からのみスペースキーでゲーム開始可能
	if g.State.State == "start" {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.State.GameRun()
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		// ゲームオーバー表示
		g.State.GameReStart()
		// カラスの位置初期化
		crow.ResetCrows()
		// オカメの位置初期化
		cockatiel.ResetEndPosition()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// 背景画像描画
	jungle.ImageDraw(screen)

	// オカメ画像描画
	cockatiel.ImageDraw(screen, g.State)

	if g.State.State != "start" {
		// カラス画像描画
		crow.ImageDraw(screen, elapsedTime, g.State)

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(4, 12)
		op.GeoM.Scale(2.5, 2.5)

		text.DrawWithOptions(screen, "Score: "+strconv.Itoa(g.State.Score)+" | Level: "+strconv.Itoa(g.State.Level), bitmapfont.Face, op)
	}

	// スタート画面
	if g.State.State == "start" {
		start.ImageDraw(screen)
	}

	// ゲームオーバー
	if g.State.State == "end" {

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(160, 200)
		op.GeoM.Scale(2.5, 2.5)
		text.DrawWithOptions(screen, "GAME OVER", bitmapfont.Face, op)

		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(130, 230)
		op.GeoM.Scale(2.5, 2.5)
		text.DrawWithOptions(screen, "Score: "+strconv.Itoa(g.State.Score)+" | Level: "+strconv.Itoa(g.State.Level), bitmapfont.Face, op)

		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(110, 260)
		op.GeoM.Scale(2.5, 2.5)
		text.DrawWithOptions(screen, "Press R Key Return Start Page", bitmapfont.Face, op)
	}

	// ebitenutil.DebugPrint(screen, strconv.Itoa(elapsedTime/60))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth * 3, ScreenHeight * 3
}
