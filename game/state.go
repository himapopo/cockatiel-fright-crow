package game

type GameState struct {
	Score int
	State string
	Level int
}

func NewGameState() *GameState {
	return &GameState{
		Score: 0,
		State: "start",
		Level: 1,
	}
}

func (gs *GameState) LevelUp() {
	// レベルマックス10
	if gs.Level == 10 {
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
}

func (gs *GameState) GameStart() {
	gs.State = "start"
}

func (gs *GameState) GameReStart() {
	gs.State = "start"
	gs.Score = 0
	gs.Level = 1
}
