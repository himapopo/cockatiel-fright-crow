package update

type GameScore struct {
	Score int
}

func NewGameScore() *GameScore {
	return &GameScore{
		Score: 0,
	}
}

func (gs *GameScore) CountUp() {
	gs.Score++
}
