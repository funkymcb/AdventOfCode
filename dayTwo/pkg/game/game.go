package game

import (
	"fmt"
	"strings"
)

const (
	Rock    = 1
	Paper   = 2
	Scissor = 3
)

const (
	Loss = 0
	Draw = 3
	Win  = 6
)

type Tournament struct {
	Games      []Game
	TotalScore int
}

type Game struct {
	Opponent int
	Response int
	Win      bool
	Loss     bool
	Score    int
}

func decipherStrategy(cipher, call string) int {
	var result int

	switch call {
	case string(cipher[0]):
		result = Rock
		break
	case string(cipher[1]):
		result = Paper
		break
	case string(cipher[2]):
		result = Scissor
		break
	}

	return result
}

func New(str string) Game {
	calls := strings.Split(str, " ")

	op := decipherStrategy("ABC", calls[0])
	resp := decipherStrategy("XYZ", calls[1])

	return Game{
		Opponent: op,
		Response: resp,
	}
}

func (g *Game) SetWin() {
	g.Win = true
}

func (g *Game) SetLoss() {
	g.Loss = true
}

func (g *Game) SetScore() {
	var score int
	score = score + g.Response
	if g.Win {
		score = score + Win
	} else if g.Loss {
		score = score + Loss
	} else {
		score = score + Draw
	}
	g.Score = score
}

func (g *Game) Play() {
	switch g.Opponent {
	case Rock:
		switch g.Response {
		case Rock:
			break
		case Paper:
			g.SetWin()
			break
		case Scissor:
			g.SetLoss()
			break
		}
	case Paper:
		switch g.Response {
		case Rock:
			g.SetLoss()
			break
		case Paper:
			break
		case Scissor:
			g.SetWin()
			break
		}
	case Scissor:
		switch g.Response {
		case Rock:
			g.SetWin()
			break
		case Paper:
			g.SetLoss()
			break
		case Scissor:
			break
		}
	}
	g.SetScore()
}

func (t *Tournament) CalculateTotalScore() {
	var total int
	for _, game := range t.Games {
		total = total + game.Score
	}
	t.TotalScore = total
}

func (t *Tournament) ToString() {
	for i, game := range t.Games {

		fmt.Printf("Game %d:\n\t%d vs %d\n\tWinner: %v\n\tScore: %d\n",
			i,
			game.Opponent,
			game.Response,
			game.Win,
			game.Score,
		)
	}
	fmt.Println("Total Score:", t.TotalScore)
}
