package crow

import (
	"cockatiel-fright-crow/update"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

var (
	// カラス画像
	crowImage *ebiten.Image

	// カラス画像サイズ
	crowImageSizeX = 0.2 // 元の大きさ 800
	crowImageSizeY = 0.2 // 元の大きさ 773

	// カラスの初期位置
	defaultcrowPositionX = 960

	crowA *crow
	crowB *crow
	crowC *crow
	crowD *crow
	crowE *crow
	crowF *crow
	crowG *crow
	crowH *crow
)

type crow struct {
	name         string
	moveCountNum float64
	moveSpeed    float64
	positionY    float64
	image        *ebiten.Image
	running      bool
}

func newCrow(name string) *crow {
	return &crow{
		name:  name,
		image: crowImage,
	}
}

func ResetCrows() {
	crowA = newCrow("A")
	crowB = newCrow("B")
	crowC = newCrow("C")
	crowD = newCrow("D")
	crowE = newCrow("E")
	crowF = newCrow("F")
	crowG = newCrow("G")
	crowH = newCrow("H")
}

func Init() {
	f, err := os.Open("./assets/img/カラス.png")
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
	crowImage = ebiten.NewImageFromImage(i)

	crowA = newCrow("A")
	crowB = newCrow("B")
	crowC = newCrow("C")
	crowD = newCrow("D")
	crowE = newCrow("E")
	crowF = newCrow("F")
	crowG = newCrow("G")
	crowH = newCrow("H")
}

func ImageDraw(screen *ebiten.Image, elapsedTime int, state *update.GameState) {

	if crowA.running {
		crowA.imageDraw(screen, state)
	}

	if crowB.running {
		crowB.imageDraw(screen, state)
	}

	if crowC.running {
		crowC.imageDraw(screen, state)
	}

	if crowD.running {
		crowD.imageDraw(screen, state)
	}

	if state.Level > 1 && crowE.running {
		crowE.imageDraw(screen, state)
	}

	if state.Level > 2 && crowF.running {
		crowF.imageDraw(screen, state)
	}

	if state.Level > 3 && crowG.running {
		crowG.imageDraw(screen, state)
	}

	if state.Level > 4 && crowH.running {
		crowH.imageDraw(screen, state)
	}

	// カラス出現頻度 最初は3秒ごと
	frequently := 3
	// カラス最高スピード
	maxSpeed := 5
	// カラス最低スピード
	minSpeed := 1

	frequentlySec := 60

	if state.Level > 1 {
		frequently = 2
	}

	if state.Level > 2 {
		frequently = 2
		maxSpeed = 7
	}

	if state.Level > 3 {
		frequently = 2
		maxSpeed = 9
	}

	if state.Level > 4 {
		frequently = 1
		maxSpeed = 11
		frequentlySec = 80
	}

	if state.Level > 5 {
		frequently = 1
		maxSpeed = 11
		frequentlySec = 60
	}

	if state.Level > 6 {
		frequently = 1
		maxSpeed = 11
		frequentlySec = 25
	}

	if state.Level > 7 {
		frequently = 1
		maxSpeed = 11
		frequentlySec = 10
	}

	if elapsedTime%frequentlySec == 0 && elapsedTime/frequentlySec%frequently == 0 {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		// Y軸のマックス値
		py := rand.Intn(550)

		// 移動スピードは 8 ~ 2の間
		cs := rand.Intn(maxSpeed-minSpeed) + minSpeed

		if !crowA.running {
			crowA.running = true
			crowA.positionY = float64(py)
			crowA.moveSpeed = float64(cs)
			return
		}

		if crowA.running &&
			!crowB.running {
			crowB.running = true
			crowB.positionY = float64(py)
			crowB.moveSpeed = float64(cs)
			return
		}

		if crowB.running &&
			!crowC.running {
			crowC.running = true
			crowC.positionY = float64(py)
			crowC.moveSpeed = float64(cs)
			return
		}

		if crowC.running &&
			!crowD.running {
			crowD.running = true
			crowD.positionY = float64(py)
			crowD.moveSpeed = float64(cs)
			return
		}

		if state.Level > 1 &&
			crowD.running &&
			!crowE.running {
			crowE.running = true
			crowE.positionY = float64(py)
			crowE.moveSpeed = float64(cs)
			return
		}

		if state.Level > 3 &&
			crowE.running &&
			!crowF.running {
			crowF.running = true
			crowF.positionY = float64(py)
			crowF.moveSpeed = float64(cs)
			return
		}

		if state.Level > 5 &&
			crowF.running &&
			!crowG.running {
			crowG.running = true
			crowG.positionY = float64(py)
			crowG.moveSpeed = float64(cs)
			return
		}

		if state.Level > 7 &&
			crowG.running &&
			!crowH.running {
			crowH.running = true
			crowH.positionY = float64(py)
			crowH.moveSpeed = float64(cs)
			return
		}
	}
}

func (c *crow) imageDraw(screen *ebiten.Image, state *update.GameState) {
	// 左に移動
	if state.State == "run" {
		c.incrementCrowPositionCount()
	}

	// オカメとカラスの衝突判定
	if state.State == "run" && c.conflictValid() {
		state.GameEnd()
	}

	if state.State == "run" {
		// 画面から消えたか判定
		if float64(defaultcrowPositionX)-c.moveCountNum < -160 {
			c.moveCountNum = 1
			c.running = false

			// スコアアップ
			state.ScoreUp()

			// 10匹ずつでレベルアップ
			if state.Score%5 == 0 {
				state.LevelUp()
			}
		}
	}

	op := &ebiten.DrawImageOptions{}

	// カラス画像の大きさ
	op.GeoM.Scale(crowImageSizeX, crowImageSizeY)

	// カラス画像の位置
	op.GeoM.Translate(float64(defaultcrowPositionX)-c.moveCountNum, c.positionY)

	// カラス画像描画
	screen.DrawImage(crowImage, op)
}

func (c *crow) incrementCrowPositionCount() {
	c.moveCountNum += c.moveSpeed
}

func (c *crow) conflictValid() bool {

	// --------オカメが画面からはみ出ない設定---------
	var cpl float64 = 0
	var cpt float64 = 0

	// オカメの一番左地点, オカメの一番上地点
	cursolx, cursoly := ebiten.CursorPosition()

	cpl = float64(cursolx)
	cpt = float64(cursoly)

	if cpl <= 0 {
		cpl = 0
	}

	if cpl >= (960 - (784 * 0.08)) {
		cpl = 960 - (784 * 0.08)
	}

	if cpt <= 0 {
		cpt = 0
	}

	if cpt >= (720 - (745 * 0.08)) {
		cpt = 720 - (745 * 0.08)
	}

	// 745, 784 はオカメ画像の元々大きさ
	cpr := cpl + (745 * 0.08) // オカメの一番右地点
	cpb := cpt + (784 * 0.08) // オカメの一番下地点

	// --------オカメが画面からはみ出ない設定---------

	// オカメの一番右地点が カラスの一番左地点を超えているか。 オカメの一番右地点がカラスの右地点より手前か。
	cockatielRightValid := cpr >= float64(defaultcrowPositionX)-c.moveCountNum && cpr <= float64(defaultcrowPositionX)-c.moveCountNum+(790*crowImageSizeX)

	// オカメの一番左地点がカラスの一番右地点より手前か。オカメの一番左地点がカラスの左地点を超えているか。
	cockatielLeftValid := float64(defaultcrowPositionX)-c.moveCountNum+(790*crowImageSizeX) >= cpl && cpl >= float64(defaultcrowPositionX)-c.moveCountNum

	// オカメの一番下地点がカラスの一番上地点を超えているか。オカメの下地点が、カラスの下地点より手前か。
	cockatielBottomValid := cpb >= c.positionY && cpb <= c.positionY+(763*crowImageSizeY)

	// オカメの一番上地点がカラスの一番下地点より手前か。オカメの一番上地点がカラスの一番上地点を超えているか。
	cockatielTopValid := c.positionY+(763*crowImageSizeY) >= cpt && cpt >= c.positionY

	// 当たり判定
	return (cockatielRightValid && (cockatielBottomValid || cockatielTopValid)) ||
		(cockatielLeftValid && (cockatielBottomValid || cockatielTopValid)) ||
		(cockatielTopValid && (cockatielRightValid || cockatielLeftValid)) ||
		(cockatielBottomValid && (cockatielRightValid || cockatielLeftValid))
}

// func execNextCrow(key string) bool {
// 	return CrowPositionCountMap[nextKey(key)]["exec"] == 1
// }

// func nextKey(key string) string {
// 	switch key {
// 	case "A":
// 		return "B"
// 	case "B":
// 		return "C"
// 	case "C":
// 		return "D"
// 	case "D":
// 		return "A"
// 	}
// 	return "A"
// }
