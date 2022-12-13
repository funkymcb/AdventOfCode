package main

import (
	"bufio"
	"os"

	"github.com/funkymcb/AdventOfCode/dayTwo/pkg/game"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func countScore(game string) int {
	return 1
}

func readGames(f *os.File) *game.Tournament {
	var t game.Tournament

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// g := game.New(scanner.Text()) // day 2 part 1
		g := game.NewCorrect(scanner.Text()) // day 2 part 2
		g.Play()
		t.Games = append(t.Games, g)
	}

	err := scanner.Err()
	checkError(err)

	return &t
}

func main() {
	file, err := os.Open("./strategy_guide")
	checkError(err)
	defer file.Close()

	tournament := readGames(file)
	tournament.CalculateTotalScore()
	tournament.ToString()
}
