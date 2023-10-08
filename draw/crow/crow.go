package crow

import (
	"bytes"
	"cockatiel-fright-crow/draw/cockatiel"
	"cockatiel-fright-crow/game"
	"image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

type Crows struct {
	crowA *crow
	crowB *crow
	crowC *crow
	crowD *crow
	crowE *crow
	crowF *crow
	crowG *crow
	crowH *crow
	crowI *crow
	rare  *crow // レアキャラ。実はオカメ。
}

var (
	// カラス最高スピード
	maxSpeed = 5
	// カラス最低スピード
	minSpeed = 1

	// カラス出現頻度 最初は3秒ごと
	frequently    = 3
	frequentlySec = 60
)

func NewCrows() *Crows {
	// カラス画像
	i, err := png.Decode(bytes.NewReader(imageByte))
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	crowImage := ebiten.NewImageFromImage(i)

	// レアキャラ画像
	i, err = png.Decode(bytes.NewReader(rareImageByte))
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	rareImage := ebiten.NewImageFromImage(i)

	return &Crows{
		crowA: newCrow("A", crowImage),
		crowB: newCrow("B", crowImage),
		crowC: newCrow("C", crowImage),
		crowD: newCrow("D", crowImage),
		crowE: newCrow("E", crowImage),
		crowF: newCrow("F", crowImage),
		crowG: newCrow("G", crowImage),
		crowH: newCrow("H", crowImage),
		crowI: newCrow("I", crowImage),
		rare:  newCrow("rare", rareImage),
	}
}

type crow struct {
	name        string
	image       *ebiten.Image
	imageWidth  float64
	imageHeight float64
	// default positionX
	dx float64
	// x軸移動値
	movex float64
	// 移動速度
	moveSpeed float64
	// current position y
	cpy float64
	// 移動中フラグ
	run bool
}

func newCrow(name string, image *ebiten.Image) *crow {
	return &crow{
		name:        name,
		image:       image,
		imageWidth:  0.2,
		imageHeight: 0.2,
		dx:          game.ScreenWidth,
	}
}

func (c *crow) resetCrow() {
	c.movex = 0
	c.moveSpeed = 0
	c.cpy = 0
	c.run = false
}

func (c *Crows) Reset() {
	c.crowA.resetCrow()
	c.crowB.resetCrow()
	c.crowC.resetCrow()
	c.crowD.resetCrow()
	c.crowE.resetCrow()
	c.crowF.resetCrow()
	c.crowG.resetCrow()
	c.crowH.resetCrow()

	maxSpeed = 9

	minSpeed = 2

	frequently = 1
	frequentlySec = 30
}

func (c *Crows) ImageDraw(screen *ebiten.Image, g *game.Game) {
	if c.runTiming() {
		// カラス投入
		c.runCrow(g)
	}

	if c.rare.run {
		c.rare.imageDraw(screen, g)
	}

	if c.crowA.run {
		c.crowA.imageDraw(screen, g)
	}

	if c.crowB.run {
		c.crowB.imageDraw(screen, g)
	}

	if c.crowC.run {
		c.crowC.imageDraw(screen, g)
	}

	if c.crowD.run {
		c.crowD.imageDraw(screen, g)
	}

	if c.crowE.run {
		c.crowE.imageDraw(screen, g)
	}

	if g.State.Level > 1 && c.crowF.run {
		c.crowF.imageDraw(screen, g)
	}

	if g.State.Level > 2 && c.crowG.run {
		c.crowG.imageDraw(screen, g)
	}

	if g.State.Level > 3 && c.crowH.run {
		c.crowH.imageDraw(screen, g)
	}

	if g.State.Level > 4 && c.crowH.run {
		c.crowI.imageDraw(screen, g)
	}
}

func (c *Crows) runTiming() bool {
	//
	return game.CallCount%frequentlySec == 0 && game.CallCount/frequentlySec%frequently == 0
}

func (c *Crows) changeCrowSpec(g *game.Game) {
	if g.State.Level > 1 {
		maxSpeed = 10
		frequentlySec = 25
	}

	if g.State.Level > 2 {
		frequently = 1
		maxSpeed = 11
		frequentlySec = 20
	}

	if g.State.Level > 3 {
		frequently = 1
		maxSpeed = 13
		frequentlySec = 10
	}

	if g.State.Level > 4 {
		frequently = 1
		maxSpeed = 14
		frequentlySec = 5
	}

	// if g.State.Level > 5 {
	// 	frequently = 1
	// 	maxSpeed = 11
	// 	frequentlySec = 40
	// }

	// if g.State.Level > 6 {
	// 	frequently = 1
	// 	maxSpeed = 11
	// 	frequentlySec = 30
	// }

	// if g.State.Level > 7 {

	// }

	// if g.State.Level > 8 {

	// }

	// if g.State.Level > 9 {

	// }
}

func (c *Crows) runCrow(g *game.Game) {

	c.changeCrowSpec(g)

	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Y軸のマックス値
	py := rand.Intn(550)

	// 移動スピードは 8 ~ 2の間
	cs := rand.Intn(maxSpeed-minSpeed) + minSpeed

	// レベル5以上からレアキャラ出現
	if g.State.Level > 4 {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		rareNum := rand.Intn(500)
		if rareNum == 250 && !c.rare.run {
			c.rare.run = true
			c.rare.cpy = float64(py)
			c.rare.moveSpeed = float64(cs)
			return
		}
	}

	if !c.crowA.run {
		c.crowA.run = true
		c.crowA.cpy = float64(py)
		c.crowA.moveSpeed = float64(cs)
		return
	}

	if c.crowA.run &&
		!c.crowB.run {
		c.crowB.run = true
		c.crowB.cpy = float64(py)
		c.crowB.moveSpeed = float64(cs)
		return
	}

	if c.crowB.run &&
		!c.crowC.run {
		c.crowC.run = true
		c.crowC.cpy = float64(py)
		c.crowC.moveSpeed = float64(cs)
		return
	}

	if c.crowC.run &&
		!c.crowD.run {
		c.crowD.run = true
		c.crowD.cpy = float64(py)
		c.crowD.moveSpeed = float64(cs)
		return
	}

	if c.crowD.run &&
		!c.crowE.run {
		c.crowE.run = true
		c.crowE.cpy = float64(py)
		c.crowE.moveSpeed = float64(cs)
		return
	}

	if g.State.Level > 1 &&
		c.crowE.run &&
		!c.crowF.run {
		c.crowF.run = true
		c.crowF.cpy = float64(py)
		c.crowF.moveSpeed = float64(cs)
		return
	}

	if g.State.Level > 2 &&
		c.crowF.run &&
		!c.crowG.run {
		c.crowG.run = true
		c.crowG.cpy = float64(py)
		c.crowG.moveSpeed = float64(cs)
		return
	}

	if g.State.Level > 3 &&
		c.crowG.run &&
		!c.crowH.run {
		c.crowH.run = true
		c.crowH.cpy = float64(py)
		c.crowH.moveSpeed = float64(cs)
		return
	}

	if g.State.Level > 4 &&
		c.crowH.run &&
		!c.crowI.run {
		c.crowI.run = true
		c.crowI.cpy = float64(py)
		c.crowI.moveSpeed = float64(cs)
		return
	}
}

func (c *crow) imageDraw(screen *ebiten.Image, g *game.Game) {
	// 左に移動
	if g.State.State == "run" {
		c.incrementCrowPositionCount()
	}

	// オカメとカラスの衝突判定
	if g.State.State == "run" && c.conflictValid() {
		// ゲーム終了ステータスに変化
		g.State.GameEnd()
		// 衝突音
		g.HitPlayer.Play()
	}

	// 画面から消えたか判定
	if g.State.State == "run" && float64(c.dx)-c.movex < -160 {
		c.resetCrow()

		// スコアアップ
		g.State.ScoreUp()

		// 10匹ずつでレベルアップ
		if g.State.Score%10 == 0 {
			g.State.LevelUp()
		}
	}

	op := &ebiten.DrawImageOptions{}

	// カラス画像の大きさ
	op.GeoM.Scale(c.imageWidth, c.imageHeight)

	// カラス画像の位置
	op.GeoM.Translate(c.dx-c.movex, c.cpy)

	// カラス画像描画
	screen.DrawImage(c.image, op)
}

func (c *crow) incrementCrowPositionCount() {
	c.movex += c.moveSpeed
}

func (c *crow) conflictValid() bool {

	cpl, cpt := cockatiel.CurrentPosition()

	// 745, 784 はオカメ画像の元々大きさ
	cpr := cpl + (745 * 0.08) // オカメの一番右地点
	cpb := cpt + (784 * 0.08) // オカメの一番下地点

	// オカメの一番右地点が カラスの一番左地点を超えているか。 オカメの一番右地点がカラスの右地点より手前か。
	cockatielRightValid := cpr >= float64(c.dx)-c.movex && cpr <= float64(c.dx)-c.movex+(790*c.imageWidth)

	// オカメの一番左地点がカラスの一番右地点より手前か。オカメの一番左地点がカラスの左地点を超えているか。
	cockatielLeftValid := float64(c.dx)-c.movex+(790*c.imageWidth) >= cpl && cpl >= float64(c.dx)-c.movex

	// オカメの一番下地点がカラスの一番上地点を超えているか。オカメの下地点が、カラスの下地点より手前か。
	cockatielBottomValid := cpb >= c.cpy && cpb <= c.cpy+(763*c.imageHeight)

	// オカメの一番上地点がカラスの一番下地点より手前か。オカメの一番上地点がカラスの一番上地点を超えているか。
	cockatielTopValid := c.cpy+(763*c.imageHeight) >= cpt && cpt >= c.cpy

	// 当たり判定
	return (cockatielRightValid && (cockatielBottomValid || cockatielTopValid)) ||
		(cockatielLeftValid && (cockatielBottomValid || cockatielTopValid)) ||
		(cockatielTopValid && (cockatielRightValid || cockatielLeftValid)) ||
		(cockatielBottomValid && (cockatielRightValid || cockatielLeftValid))
}
