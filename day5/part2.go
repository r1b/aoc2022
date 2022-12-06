package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// --------------------------------------------------------------------------------

const CRATE_BUF_SIZE = 4

type Instruction struct {
	count    int
	srcStack int
	dstStack int
}

type Ship map[int]*list.List

// --------------------------------------------------------------------------------

func parseShip(scanner *bufio.Scanner) (Ship, int) {
	var stack int
	ship := make(Ship)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		reader := strings.NewReader(line)

		for stack = 0; ; stack++ {
			buf := make([]byte, CRATE_BUF_SIZE)
			// Crates and axes labels are 4 bytes for column 1..n-1, 3 bytes for column n
			bytesRead, err := io.ReadAtLeast(reader, buf, CRATE_BUF_SIZE-1)
			if err == io.EOF && bytesRead == 0 {
				break
			}
			// This skips both empty stack positions and axes labels
			if buf[0] != '[' {
				continue
			}
			crate := buf[1]
			if _, present := ship[stack]; !present {
				ship[stack] = list.New()
			}
			ship[stack].PushBack(crate)
		}
	}

	return ship, stack
}

func parseInstructions(scanner *bufio.Scanner) []Instruction {
	var instructions []Instruction

	for scanner.Scan() {
		line := scanner.Text()
		reader := strings.NewReader(line)
		wordScanner := bufio.NewScanner(reader)
		wordScanner.Split(bufio.ScanWords)

		var words []string

		for wordScanner.Scan() {
			words = append(words, wordScanner.Text())
		}

		count, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}

		srcStack, err := strconv.Atoi(words[3])
		if err != nil {
			panic(err)
		}

		dstStack, err := strconv.Atoi(words[5])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, Instruction{count: count, srcStack: srcStack - 1, dstStack: dstStack - 1})
	}

	return instructions
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ship, stackCount := parseShip(scanner)
	instructions := parseInstructions(scanner)

	for _, instruction := range instructions {
		srcStack := ship[instruction.srcStack]
		dstStack := ship[instruction.dstStack]

		tmp := list.New()
		for i := 0; i < instruction.count; i++ {
			tmp.PushBack(srcStack.Remove(srcStack.Front()))
		}
		for i := 0; i < instruction.count; i++ {
			dstStack.PushFront(tmp.Remove(tmp.Back()))
		}
	}

	var topCrates []byte

	for stack := 0; stack < stackCount; stack++ {
		topCrateElement := ship[stack].Front()
		topCrates = append(topCrates, topCrateElement.Value.(byte))
	}

	fmt.Println("Top: ", string(topCrates))
}
