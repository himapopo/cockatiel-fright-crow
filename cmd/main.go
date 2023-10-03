package main

import (
	"log"

	"cockatiel-fright-crow/draw/cockatiel"
	"cockatiel-fright-crow/draw/crow"
	"cockatiel-fright-crow/draw/jungle"
	"cockatiel-fright-crow/draw/start"
	"cockatiel-fright-crow/game"
	"cockatiel-fright-crow/update"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/xerrors"
)

func main() {

	cockatiel.Init()

	jungle.Init()

	crow.Init()

	start.Init()

	ebiten.SetWindowSize(game.ScreenWidth*3, game.ScreenHeight*3)
	ebiten.SetWindowTitle("Cockatiel Fright Crow")
	if err := ebiten.RunGame(&game.Game{
		State: update.NewGameState(),
	}); err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
}
