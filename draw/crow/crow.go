package crow

import (
	"image"
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

	// カラスの移動値
	CrowPositionCountMap = map[string]map[string]float64{
		"A": {
			"count": 1, "exec": 0, "positionY": 0, "countSpeed": 0,
		},
		"B": {
			"count": 1, "exec": 0, "positionY": 0, "countSpeed": 0,
		},
		"C": {
			"count": 1, "exec": 0, "positionY": 0, "countSpeed": 0,
		},
		"D": {
			"count": 1, "exec": 0, "positionY": 0, "countSpeed": 0,
		},
	}
)

func Init() {
	kf, err := os.Open("./assets/img/カラス.png")
	defer kf.Close()
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	ki, _, err := image.Decode(kf)
	if err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
	crowImage = ebiten.NewImageFromImage(ki)
}

func ImageDraw(screen *ebiten.Image, elapsedTime int) {

	if CrowPositionCountMap["A"]["exec"] == 1 {
		crowImageDraw("A", screen)
	}

	if CrowPositionCountMap["B"]["exec"] == 1 {
		crowImageDraw("B", screen)
	}

	if CrowPositionCountMap["C"]["exec"] == 1 {
		crowImageDraw("C", screen)
	}

	if CrowPositionCountMap["D"]["exec"] == 1 {
		crowImageDraw("D", screen)
	}

	if elapsedTime%60 == 0 && elapsedTime/60%3 == 0 {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		// Y軸のマックス値
		py := rand.Intn(500)

		// 移動スピードは 8 ~ 2の間
		cs := rand.Intn(8-1) + 1

		if CrowPositionCountMap["A"]["exec"] == 0 {
			CrowPositionCountMap["A"]["exec"] = 1
			CrowPositionCountMap["A"]["positionY"] = float64(py)
			CrowPositionCountMap["A"]["countSpeed"] = float64(cs)
			return
		}

		if CrowPositionCountMap["A"]["exec"] == 1 &&
			CrowPositionCountMap["B"]["exec"] == 0 {
			CrowPositionCountMap["B"]["exec"] = 1
			CrowPositionCountMap["B"]["positionY"] = float64(py)
			CrowPositionCountMap["B"]["countSpeed"] = float64(cs)
			return
		}

		if CrowPositionCountMap["B"]["exec"] == 1 &&
			CrowPositionCountMap["C"]["exec"] == 0 {
			CrowPositionCountMap["C"]["exec"] = 1
			CrowPositionCountMap["C"]["positionY"] = float64(py)
			CrowPositionCountMap["C"]["countSpeed"] = float64(cs)
			return
		}

		if CrowPositionCountMap["C"]["exec"] == 1 &&
			CrowPositionCountMap["D"]["exec"] == 0 {
			CrowPositionCountMap["D"]["exec"] = 1
			CrowPositionCountMap["D"]["positionY"] = float64(py)
			CrowPositionCountMap["D"]["countSpeed"] = float64(cs)
			return
		}

		if CrowPositionCountMap["D"]["exec"] == 1 {
			return
		}
	}

}

func crowImageDraw(key string, screen *ebiten.Image) bool {

	IncrementCrowPositionCount(key)

	if float64(defaultcrowPositionX)-float64(CrowPositionCountMap[key]["count"]) < -160 {
		CrowPositionCountMap[key]["count"] = 1
		CrowPositionCountMap[key]["exec"] = 0
		return false
	}

	op := &ebiten.DrawImageOptions{}

	// カラス画像の大きさ
	op.GeoM.Scale(crowImageSizeX, crowImageSizeY)

	// カラス画像の位置
	op.GeoM.Translate(float64(defaultcrowPositionX)-float64(CrowPositionCountMap[key]["count"]), CrowPositionCountMap[key]["positionY"])

	// カラス画像描画
	screen.DrawImage(crowImage, op)

	return true
}

func IncrementCrowPositionCount(key string) {
	CrowPositionCountMap[key]["count"] += CrowPositionCountMap[key]["countSpeed"]
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
