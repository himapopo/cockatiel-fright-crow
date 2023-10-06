package update

import (
	"bytes"
	"log"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	raudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
)

var (
	AudioContext *oto.Context
)

type GameState struct {
	Score int
	State string
	Level int

	hitPlayer *oto.Player
}

func NewGameState() *GameState {

	// 効果音初期化 TODO: 別の箇所に処理切り出す
	jabD, err := wav.DecodeWithoutResampling(bytes.NewReader(raudio.Jab_wav))
	if err != nil {
		log.Fatal(err)
	}

	player := AudioContext.NewPlayer(jabD)

	return &GameState{
		Score: 0,
		State: "start",
		Level: 1,

		hitPlayer: player,
	}
}

func (gs *GameState) LevelUp() {
	// レベルマックス8
	if gs.Level == 8 {
		return
	}
	gs.Level++
}

func (gs *GameState) ScoreUp() {
	gs.Score++
}

func (gs *GameState) GameRun() {
	gs.State = "run"
}

func (gs *GameState) GameEnd() {
	gs.State = "end"

	gs.hitPlayer.Seek(0, 0)
	gs.hitPlayer.Play()
}

func (gs *GameState) GameStart() {
	gs.State = "start"
}

func (gs *GameState) GameReStart() {
	gs.State = "start"
	gs.Score = 0
	gs.Level = 1
}
