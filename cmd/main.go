package main

import (
	"log"

	"cockatiel-fright-crow/draw/cockatiel"
	"cockatiel-fright-crow/draw/crow"
	"cockatiel-fright-crow/draw/jungle"
	"cockatiel-fright-crow/game"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

func main() {

	cockatiel.Init()

	jungle.Init()

	crow.Init()

	ebiten.SetWindowSize(game.ScreenWidth*3, game.ScreenHeight*3)
	ebiten.SetWindowTitle("Cockatiel Fright Crow")
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
}
