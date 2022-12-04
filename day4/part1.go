package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --------------------------------------------------------------------------------

type Range struct {
	start, end int
}

func (r1 Range) Contains(r2 Range) bool {
	return (r1.start <= r2.start) && (r1.end >= r2.end)
}

// --------------------------------------------------------------------------------

func main() {
	count := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		elves := strings.Split(line, ",")

		r1Spec := strings.Split(elves[0], "-")
		r2Spec := strings.Split(elves[1], "-")

		r1Start, err := strconv.Atoi(r1Spec[0])
		if err != nil {
			fmt.Println("Could not convert spec", r1Spec[0])
			return
		}
		r1End, err := strconv.Atoi(r1Spec[1])
		if err != nil {
			fmt.Println("Could not convert spec", r1Spec[1])
			return
		}
		r2Start, err := strconv.Atoi(r2Spec[0])
		if err != nil {
			fmt.Println("Could not convert spec", r2Spec[0])
			return
		}
		r2End, err := strconv.Atoi(r2Spec[1])
		if err != nil {
			fmt.Println("Could not convert spec", r2Spec[1])
			return
		}

		r1 := Range{start: r1Start, end: r1End}
		r2 := Range{start: r2Start, end: r2End}

		if r1.Contains(r2) || r2.Contains(r1) {
			count += 1
		}
	}

	fmt.Println("Count: ", count)
}
