package main

import (
	"bufio"
	"fmt"
	"os"
)

// --------------------------------------------------------------------------------

const WINDOW_SIZE = 4

// --------------------------------------------------------------------------------

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	buf := scanner.Text()

	counts := make(map[byte]int)
	start := 0
	end := WINDOW_SIZE

	for i := start; i < end; i++ {
		counts[buf[i]] += 1
	}

	for ; end < len(buf); start, end = start+1, end+1 {
		if len(counts) == WINDOW_SIZE {
			fmt.Println("Found marker", end)
			return
		}

		startChar := buf[start]
		endChar := buf[end]

		counts[startChar] -= 1
		if counts[startChar] == 0 {
			delete(counts, startChar)
		}

		counts[endChar] += 1
	}
}
