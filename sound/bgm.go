package sound

import (
	"bytes"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/go-mp3"
)

type BgmPlayer struct {
	player *oto.Player
}

func NewBgmPlayer(ac *AudioContext) *BgmPlayer {
	decodedMp3, err := mp3.NewDecoder(bytes.NewReader(bgmByte))
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	loo := audio.NewInfiniteLoop(decodedMp3, decodedMp3.Length())

	player := ac.Ctx.NewPlayer(loo)

	return &BgmPlayer{
		player: player,
	}
}

func (b *BgmPlayer) Play() {
	b.player.Play()
	defer b.player.Close()

	// We can wait for the sound to finish playing using something like this
	for b.player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}
