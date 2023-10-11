package main

import (
	_ "embed"
	"log"

	"cockatiel-fright-crow/draw/cockatiel"
	"cockatiel-fright-crow/draw/crow"
	"cockatiel-fright-crow/draw/jungle"
	"cockatiel-fright-crow/draw/start"
	"cockatiel-fright-crow/game"
	"cockatiel-fright-crow/sound"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

var (
	//go:embed assets/img/start.jpg
	startImageByte []byte

	//go:embed assets/img/cockatiel.png
	cockatielImageByte []byte

	//go:embed assets/img/crow.png
	crowImageByte []byte

	//go:embed assets/img/rare_cockatiel.png
	rareCockatielImageByte []byte

	//go:embed assets/img/jungle.png
	jungleImageByte []byte

	//go:embed assets/bgm/bgm.mp3
	bgmByte []byte
)

func main() {
	ac := sound.NewAudioContext()

	// bgm := sound.NewBgmPlayer(bgmByte, ac)

	// go bgm.Play()

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Cockatiel Fright Crow")
	if err := ebiten.RunGame(&game.Game{
		Start:          start.NewStart(startImageByte),
		JungleImage:    jungle.NewJungle(jungleImageByte),
		CockatielImage: cockatiel.NewCockatiel(cockatielImageByte),
		Crows:          crow.NewCrows(crowImageByte, rareCockatielImageByte),
		State:          game.NewGameState(),
		HitPlayer:      sound.NewHitPlayer(ac),
	}); err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
}
