package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var flashCount int

var flashPoints = map[struct{ x, y int }]struct{}{}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]int

	for scanner.Scan() {
		line := scanner.Text()

		l := make([]int, len(line))

		for i := range line {
			l[i] = atoi(line[i])
		}

		grid = append(grid, l)
	}

	printGrid(grid)

	var syncStep int

	for {
		syncStep++
		flashPoints = make(map[struct{ x, y int }]struct{})
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[0]); y++ {
				grid = flash(grid, x, y)
			}
		}
		printGrid(grid)

		if len(flashPoints) == len(grid)*len(grid[0]) {
			break
		}
	}

	fmt.Fprintf(os.Stdout, "answer: %d\n", syncStep)
}

func atoi(b byte) int {
	return int(b) - 48
}

func printGrid(grid [][]int) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			print(grid[x][y])
		}
		println()
	}
	println("---")
}

func flash(grid [][]int, x, y int) [][]int {

	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return grid
	}

	if _, ok := flashPoints[struct{ x, y int }{x, y}]; ok {
		return grid
	}

	if grid[x][y] == 9 {
		grid[x][y] = 0
		flashPoints[struct{ x, y int }{x, y}] = struct{}{}
		flashCount++
		flash(grid, x-1, y-1)
		flash(grid, x-1, y)
		flash(grid, x-1, y+1)
		flash(grid, x, y+1)
		flash(grid, x, y-1)
		flash(grid, x+1, y-1)
		flash(grid, x+1, y)
		flash(grid, x+1, y+1)
	} else {
		grid[x][y]++
	}

	return grid
}
