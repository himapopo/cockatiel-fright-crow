package game

import (
	"strconv"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	// スクリーンサイズ
	ScreenWidth  = 960
	ScreenHeight = 720
)

var (
	CallCount = 1
)

type HitPlayer interface {
	Play()
}

type Start interface {
	ImageDraw(creen *ebiten.Image)
}

type Jungle interface {
	ImageDraw(creen *ebiten.Image)
}

type Cockatiel interface {
	ImageDraw(creen *ebiten.Image, state *Game)
	Reset()
}

type Crows interface {
	ImageDraw(creen *ebiten.Image, game *Game)
	Reset()
}

type Game struct {
	Start          Start
	JungleImage    Jungle
	CockatielImage Cockatiel
	Crows          Crows

	HitPlayer HitPlayer
	State     *GameState
}

func (g *Game) Update() error {
	CallCount++
	// CallCountが永遠増えないように
	if CallCount%60 == 0 && CallCount == 101 {
		CallCount = 1
	}

	// スタート画面からのみスペースキーでゲーム開始可能
	if g.State.State == "start" {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.State.GameRun()
		}
	}

	if g.State.State == "end" {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.State.GameReStart()
			// カラスの位置初期化
			g.Crows.Reset()
			// オカメの位置初期化
			g.CockatielImage.Reset()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// 背景画像描画
	g.JungleImage.ImageDraw(screen)

	// オカメ画像描画
	g.CockatielImage.ImageDraw(screen, g)

	if g.State.State != "start" {
		// カラス画像描画
		g.Crows.ImageDraw(screen, g)

		g.textDraw(screen, 4, 12, 2.5, 2.5, "Score: "+strconv.Itoa(g.State.Score))
	}

	// スタート画面
	if g.State.State == "start" {
		g.Start.ImageDraw(screen)
	}

	// ゲームオーバー
	if g.State.State == "end" {

		g.textDraw(screen, 160, 200, 2.5, 2.5, "GAME OVER")

		g.textDraw(screen, 160, 230, 2.5, 2.5, "Score: "+strconv.Itoa(g.State.Score))

		g.textDraw(screen, 110, 260, 2.5, 2.5, "Press R Key Return Start Page")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) textDraw(screen *ebiten.Image, tx, ty, sx, sy float64, msg string) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(tx, ty)
	op.GeoM.Scale(sx, sy)
	text.DrawWithOptions(screen, msg, bitmapfont.Face, op)
}
