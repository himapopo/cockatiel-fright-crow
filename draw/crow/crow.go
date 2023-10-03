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
	crowImageSizeX = 0.2
	crowImageSizeY = 0.2

	// カラスの初期位置
	defaultcrowPositionX = 960

	crowA *crow
	crowB *crow
	crowC *crow
	crowD *crow
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
}

func ImageDraw(screen *ebiten.Image, elapsedTime int, score *update.GameScore) {

	if crowA.running {
		crowA.imageDraw(screen, score)
	}

	if crowB.running {
		crowB.imageDraw(screen, score)
	}

	if crowC.running {
		crowC.imageDraw(screen, score)
	}

	if crowD.running {
		crowD.imageDraw(screen, score)
	}

	if elapsedTime%60 == 0 && elapsedTime/60%3 == 0 {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		// Y軸のマックス値
		py := rand.Intn(500)

		// 移動スピードは 8 ~ 2の間
		cs := rand.Intn(8-1) + 1

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

		if crowD.running {
			return
		}
	}
}

func (c *crow) imageDraw(screen *ebiten.Image, score *update.GameScore) {
	// 左に移動
	c.incrementCrowPositionCount()

	// 画面から消えたか判定
	if float64(defaultcrowPositionX)-c.moveCountNum < -160 {
		c.moveCountNum = 1
		c.running = false

		// スコアアップ
		score.ScoreUp()
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
