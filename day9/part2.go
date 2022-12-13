package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --------------------------------------------------------------------------------

const NUM_KNOTS = 10

// --------------------------------------------------------------------------------

type Visited map[string]bool

type Posn struct {
	i int
	j int
}

func (p Posn) key() string {
	return strconv.Itoa(p.i) + "," + strconv.Itoa(p.j)
}

type Move struct {
	delta Posn
	count int
}

func NewMove(i, j, count int) Move {
	return Move{delta: Posn{i: i, j: j}, count: count}
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

type Rope struct {
	knots [NUM_KNOTS]Posn
}

func NewRope() Rope {
	var knots [NUM_KNOTS]Posn
	for i := 0; i < NUM_KNOTS; i++ {
		knots[i] = Posn{i: 0, j: 0}
	}
	return Rope{knots: knots}
}

func (r Rope) apply(delta Posn) Rope {
	var nextKnots [NUM_KNOTS]Posn
	nextKnots[0] = Posn{i: r.knots[0].i + delta.i, j: r.knots[0].j + delta.j}
	for knot := 1; knot < NUM_KNOTS; knot++ {
		nextLeaderI := nextKnots[knot-1].i
		nextLeaderJ := nextKnots[knot-1].j

		curFollowerI := r.knots[knot].i
		curFollowerJ := r.knots[knot].j

		nextKnot := Posn{i: curFollowerI, j: curFollowerJ}

		if abs(nextLeaderI, curFollowerI) == 2 || abs(nextLeaderJ, curFollowerJ) == 2 {
			if abs(nextLeaderI, curFollowerI) > 0 {
				if nextLeaderI > curFollowerI {
					nextKnot.i = curFollowerI + 1
				} else {
					nextKnot.i = curFollowerI - 1
				}
			}
			if abs(nextLeaderJ, curFollowerJ) > 0 {
				if nextLeaderJ > curFollowerJ {
					nextKnot.j = curFollowerJ + 1
				} else {
					nextKnot.j = curFollowerJ - 1
				}
			}
		}

		nextKnots[knot] = nextKnot
	}
	return Rope{knots: nextKnots}
}

// --------------------------------------------------------------------------------

func main() {
	var moves []Move
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		count, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		if tokens[0] == "U" {
			moves = append(moves, NewMove(1, 0, count))
		} else if tokens[0] == "D" {
			moves = append(moves, NewMove(-1, 0, count))
		} else if tokens[0] == "L" {
			moves = append(moves, NewMove(0, -1, count))
		} else if tokens[0] == "R" {
			moves = append(moves, NewMove(0, 1, count))
		}
	}

	visited := make(Visited)
	rope := NewRope()

	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			rope = rope.apply(move.delta)
			visited[rope.knots[NUM_KNOTS-1].key()] = true
		}
	}

	fmt.Println("Visited:", len(visited))
}
