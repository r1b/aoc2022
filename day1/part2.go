package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var calorieTotals []int

	curTotal := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			calorieTotals = append(calorieTotals, curTotal)
			curTotal = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Could not convert line to int", line)
			return
		}
		curTotal += calories
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calorieTotals)))

	top3Sum := 0
	for _, calories := range calorieTotals[:3] {
		top3Sum += calories
	}
	fmt.Println("Top 3 calories: ", top3Sum)
}
