package update

type GameState struct {
	Score int
	State string
}

func NewGameState() *GameState {
	return &GameState{
		Score: 0,
		State: "start",
	}
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
}