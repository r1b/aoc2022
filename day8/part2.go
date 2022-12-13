package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// --------------------------------------------------------------------------------

func key(i, j int) string {
	return strconv.Itoa(i) + "," + strconv.Itoa(j)
}

func viewingDistance(forest *[][]int, i, j, iDelta, jDelta int) int {
	distance := 0
	height := (*forest)[i][j]

	for {
		i += iDelta
		j += jDelta

		if i < 0 || i >= len(*forest) || j < 0 || j >= len((*forest)[i]) {
			return distance
		}

		distance += 1

		if (*forest)[i][j] >= height {
			return distance
		}
	}
}

func main() {
	var forest [][]int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for _, c := range line {
			tree, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	bestScore := 0

	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			right := viewingDistance(&forest, i, j, 0, 1)
			up := viewingDistance(&forest, i, j, 1, 0)
			left := viewingDistance(&forest, i, j, -1, 0)
			down := viewingDistance(&forest, i, j, 0, -1)

			score := right * up * left * down

			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Println("Best score:", bestScore)
}
