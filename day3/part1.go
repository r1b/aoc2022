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

func findCommonItem(a Compartment, b Compartment) rune {
	for item := range a {
		if b[item] != 0 {
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

func splitRucksack(rucksack string) (Compartment, Compartment) {
	m := len(rucksack) / 2
	return fillCompartment(rucksack[:m]), fillCompartment(rucksack[m:])
}

func main() {
	prioritiesSum := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		compartmentA, compartmentB := splitRucksack(line)
		commonItem := findCommonItem(compartmentA, compartmentB)
		if commonItem <= 'Z' {
			prioritiesSum += int(commonItem - '&')
		} else {
			prioritiesSum += int(commonItem - '`')
		}
	}

	fmt.Println("Sum: ", prioritiesSum)
}
