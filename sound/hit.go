package sound

import (
	"bytes"
	"log"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
)

type HitPlayer struct {
	player *oto.Player
}

func NewHitPlayer(ac *AudioContext) *HitPlayer {
	jabD, err := wav.DecodeWithoutResampling(bytes.NewReader(audio.Jab_wav))
	if err != nil {
		log.Fatal(err)
	}

	player := ac.Ctx.NewPlayer(jabD)

	return &HitPlayer{
		player: player,
	}
}

func (h *HitPlayer) Play() {
	h.player.Seek(0, 0)
	h.player.Play()
}
