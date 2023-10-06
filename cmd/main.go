package main

import (
	"bytes"
	"log"
	"os"
	"time"

	"cockatiel-fright-crow/draw/cockatiel"
	"cockatiel-fright-crow/draw/crow"
	"cockatiel-fright-crow/draw/jungle"
	"cockatiel-fright-crow/draw/start"
	"cockatiel-fright-crow/game"
	"cockatiel-fright-crow/update"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/go-mp3"
	"golang.org/x/xerrors"
)

func main() {

	cockatiel.Init()

	jungle.Init()

	crow.Init()

	start.Init()

	// BGM再生 TODO: 別の箇所に処理切り出す
	fileBytes, err := os.ReadFile("./bgm/追いかけっこキャッハー.mp3")
	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	}

	decodedMp3, err := mp3.NewDecoder(bytes.NewReader(fileBytes))
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	loo := audio.NewInfiniteLoop(decodedMp3, decodedMp3.Length())

	op := &oto.NewContextOptions{}

	op.SampleRate = 44100

	op.ChannelCount = 2

	op.Format = oto.FormatSignedInt16LE

	otoCtx, readyChan, err := oto.NewContext(op)
	update.AudioContext = otoCtx
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	player := otoCtx.NewPlayer(loo)

	fu := func() {
		player.Play()
		defer player.Close()

		// We can wait for the sound to finish playing using something like this
		for player.IsPlaying() {
			time.Sleep(time.Millisecond)
		}
	}

	go fu()

	ebiten.SetWindowSize(game.ScreenWidth*3, game.ScreenHeight*3)
	ebiten.SetWindowTitle("Cockatiel Fright Crow")
	if err := ebiten.RunGame(&game.Game{
		State: update.NewGameState(),
	}); err != nil {
		e := xerrors.Errorf("error: %w", err)
		log.Fatalf("%+v\n", e)
	}
}
