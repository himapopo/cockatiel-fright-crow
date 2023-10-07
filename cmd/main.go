package main

import (
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

func main() {

	ac := sound.NewAudioContext()

	bgm := sound.NewBgmPlayer(ac)

	go bgm.Play()

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Cockatiel Fright Crow")
	if err := ebiten.RunGame(&game.Game{
		Start:          start.NewStart(),
		JungleImage:    jungle.NewJungle(),
		CockatielImage: cockatiel.NewCockatiel(),
		Crows:          crow.NewCrows(),
		State:          game.NewGameState(),
		HitPlayer:      sound.NewHitPlayer(ac),
	}); err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
}
