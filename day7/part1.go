package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --------------------------------------------------------------------------------

const MAX_SIZE = 100000

// --------------------------------------------------------------------------------

type Directory struct {
	name     string
	size     int
	children []*Directory
}

func NewDirectory(name string) *Directory {
	dir := new(Directory)
	dir.name = name
	return dir
}

func (d *Directory) countSmall(total *int) int {
	totalSize := d.size

	for _, child := range d.children {
		totalSize += child.countSmall(total)
	}

	if totalSize < MAX_SIZE {
		*total += totalSize
	}

	return totalSize
}

// --------------------------------------------------------------------------------

func parseTree(stack *list.List) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cur := stack.Back().Value.(*Directory)
		line := scanner.Text()

		if line[:4] == "$ ls" {
			continue
		}

		if line[:4] == "$ cd" {
			name := line[5:]
			if name == ".." {
				stack.Remove(stack.Back())
			} else {
				dir := NewDirectory(name)
				cur.children = append(cur.children, dir)
				stack.PushBack(dir)
			}
		} else {
			reader := strings.NewReader(line)
			wordScanner := bufio.NewScanner(reader)
			wordScanner.Split(bufio.ScanWords)

			wordScanner.Scan()
			word := wordScanner.Text()
			if word != "dir" {
				size, err := strconv.Atoi(word)
				if err != nil {
					panic(err)
				}
				cur.size += size
			}
		}
	}
}

func parseRoot() *Directory {
	dummy := NewDirectory("dummy")

	stack := list.New()
	stack.PushBack(dummy)

	parseTree(stack)

	return dummy.children[0]
}

// --------------------------------------------------------------------------------

func main() {
	root := parseRoot()
	total := 0
	root.countSmall(&total)
	fmt.Println("Total size:", total)
}
