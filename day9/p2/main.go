package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// X and . fullfill is cool and easy for debugging
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var heightmap [][]string

	placeHolder := make([]string, 102)
	for index := range placeHolder {
		placeHolder[index] = "X"
	}
	heightmap = append(heightmap, placeHolder)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		row = append(row, "X")
		for index := range line {
			if line[index] == '9' {
				row = append(row, "X")
			} else {
				row = append(row, ".")
			}
		}
		row = append(row, "X")
		heightmap = append(heightmap, row)
	}
	heightmap = append(heightmap, placeHolder)

	var basins []int
	for x := 1; x < len(heightmap)-1; x++ {
		for y := 1; y < len(heightmap[x])-1; y++ {
			var block int
			block = findUntil(block, heightmap, x, y)
			if block > 0 {
				basins = append(basins, block)
			}
		}
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	fmt.Printf("basins: %v", basins[:3])
	fmt.Fprintf(os.Stdout, "answer: %d\n", basins[0]*basins[1]*basins[2])

}

func findUntil(block int, m [][]string, x, y int) int {
	if m[x][y] == "X" {
		return block
	}
	block++
	m[x][y] = "X"
	block = findUntil(block, m, x+1, y)
	block = findUntil(block, m, x, y+1)
	block = findUntil(block, m, x-1, y)
	block = findUntil(block, m, x, y-1)
	return block
}

func atoi(r string) int {
	i, _ := strconv.Atoi(r)
	return i
}
