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

func main() {
	var forest [][]int
	visible := make(map[string]bool)
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

	// left to right
	for i := 0; i < len(forest); i++ {
		largest := forest[i][0]
		visible[key(i, 0)] = true
		for j := 1; j < len(forest[i]); j++ {
			if forest[i][j] > largest {
				largest = forest[i][j]
				visible[key(i, j)] = true
			}
		}
	}

	// right to left
	for i := 0; i < len(forest); i++ {
		largest := forest[i][len(forest[i])-1]
		visible[key(i, len(forest[i])-1)] = true
		for j := len(forest[i]) - 2; j >= 0; j-- {
			if forest[i][j] > largest {
				largest = forest[i][j]
				visible[key(i, j)] = true
			}
		}
	}

	// top to bottom
	for j := 0; j < len(forest); j++ {
		largest := forest[0][j]
		visible[key(0, j)] = true
		for i := 1; i < len(forest); i++ {
			if forest[i][j] > largest {
				largest = forest[i][j]
				visible[key(i, j)] = true
			}
		}
	}

	// bottom to top
	for j := 0; j < len(forest); j++ {
		largest := forest[len(forest[j])-1][j]
		visible[key(len(forest[j])-1, j)] = true
		for i := len(forest[j]) - 2; i >= 0; i-- {
			if forest[i][j] > largest {
				largest = forest[i][j]
				visible[key(i, j)] = true
			}
		}
	}

	fmt.Println("Visible:", len(visible))
}
