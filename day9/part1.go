package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	head Posn
	tail Posn
}

func NewRope() Rope {
	return Rope{head: Posn{i: 0, j: 0}, tail: Posn{i: 0, j: 0}}
}

func (r Rope) apply(delta Posn) Rope {
	var nextTail Posn
	nextHead := Posn{i: r.head.i + delta.i, j: r.head.j + delta.j}
	if abs(nextHead.i, r.tail.i) > 1 || abs(nextHead.j, r.tail.j) > 1 {
		nextTail = r.head
	} else {
		nextTail = r.tail
	}
	return Rope{head: nextHead, tail: nextTail}
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
			visited[rope.tail.key()] = true
		}
	}

	fmt.Println("Visited:", len(visited))
}
