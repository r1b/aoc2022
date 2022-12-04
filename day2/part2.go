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
var YourMoveDeltas = map[string]int{
	"X": -1,
	"Y": 0,
	"Z": 1,
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

		theirMove := TheirMoves[moves[0]]
		// XXX: We want a euclidian mod..
		yourMove := ((theirMove+YourMoveDeltas[moves[1]])%3 + 3) % 3

		score += roundPoints(theirMove, yourMove)
	}

	fmt.Println("Score: ", score)
}
