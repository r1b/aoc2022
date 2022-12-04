package main

import (
	"bufio"
	"fmt"
	"os"
)

// --------------------------------------------------------------------------------

// char -> count
type Compartment map[rune]int

// --------------------------------------------------------------------------------

func findCommonItem(group [3]Compartment) rune {
	for item := range group[0] {
		if group[1][item] != 0 && group[2][item] != 0 {
			return item
		}
	}
	return 0
}

func fillCompartment(items string) Compartment {
	compartment := Compartment{}
	for _, item := range items {
		compartment[item] += 1
	}
	return compartment
}

func main() {
	prioritiesSum := 0
	scanner := bufio.NewScanner(os.Stdin)

	var group [3]Compartment
	i := 0
	for scanner.Scan() {
		items := scanner.Text()
		group[i] = fillCompartment(items)
		i = (i + 1) % 3
		if i == 0 {
			commonItem := findCommonItem(group)
			if commonItem <= 'Z' {
				prioritiesSum += int(commonItem - '&')
			} else {
				prioritiesSum += int(commonItem - '`')
			}
		}
	}

	fmt.Println("Sum: ", prioritiesSum)
}
