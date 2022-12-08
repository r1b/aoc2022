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

const DESIRED_FREE_SPACE = 30000000
const TOTAL_DISK_SPACE = 70000000

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

func (d *Directory) totalSize() int {
	totalSize := d.size

	for _, child := range d.children {
		totalSize += child.totalSize()
	}

	return totalSize
}

// TODO: DRY
func (d *Directory) findSmallest(smallestSize *int, minimumSizeToDelete int) int {
	totalSize := d.size

	for _, child := range d.children {
		totalSize += child.findSmallest(smallestSize, minimumSizeToDelete)
	}

	if totalSize > minimumSizeToDelete && totalSize < *smallestSize {
		*smallestSize = totalSize
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

	usedSize := root.totalSize()
	minimumSizeToDelete := DESIRED_FREE_SPACE - (TOTAL_DISK_SPACE - usedSize)

	smallestSize := usedSize
	root.findSmallest(&smallestSize, minimumSizeToDelete)
	fmt.Println("Smallest size:", smallestSize)
}
