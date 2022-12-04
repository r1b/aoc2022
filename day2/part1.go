package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// --------------------------------------------------------------------------------

const (
	Rock = iota
	Paper
	Scissors
)
const DrawScore = 3
const WinScore = 6

var MovePoints = []int{1, 2, 3}
var TheirMoves = map[string]int{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}
var YourMoves = map[string]int{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

// --------------------------------------------------------------------------------

func roundPoints(theirMove int, yourMove int) int {
	movePoints := MovePoints[yourMove]
	// Win
	if yourMove == (theirMove+1)%3 {
		return movePoints + WinScore
	}
	// Draw
	if yourMove == theirMove {
		return movePoints + DrawScore
	}
	// Loss
	return movePoints
}

func main() {
	score := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(line, " ")
		score += roundPoints(TheirMoves[moves[0]], YourMoves[moves[1]])
	}

	fmt.Println("Score: ", score)
}
