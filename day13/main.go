package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	paper := make(map[point]struct{})

	for scanner.Scan() {

		parts := strings.Split(scanner.Text(), ",")

		paper[point{atoi(parts[0]), atoi(parts[1])}] = struct{}{}

	}

	paper = foldByX(paper, 655)
	paper = foldByY(paper, 447)
	paper = foldByX(paper, 327)
	paper = foldByY(paper, 223)
	paper = foldByX(paper, 163)
	paper = foldByY(paper, 111)
	paper = foldByX(paper, 81)
	paper = foldByY(paper, 55)
	paper = foldByX(paper, 40)
	paper = foldByY(paper, 27)
	paper = foldByY(paper, 13)
	paper = foldByY(paper, 6)

	printPaper(paper)

}

func printPaper(paper map[point]struct{}) {
	minX, minY := 0, 0
	maxX, maxY := 0, 0
	for p := range paper {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := paper[point{x, y}]; ok {
				fmt.Fprint(os.Stdout, "#")
			} else {
				fmt.Fprint(os.Stdout, ".")
			}
		}
		fmt.Fprintln(os.Stdout)
	}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func foldByX(paper map[point]struct{}, x int) map[point]struct{} {
	for p := range paper {
		if p.x == x {
			delete(paper, p)
		}
		if p.x > x {
			delete(paper, p)
			paper[point{x: x - (p.x - x), y: p.y}] = struct{}{}
		}
	}
	return paper
}

func foldByY(paper map[point]struct{}, y int) map[point]struct{} {
	for p := range paper {
		if p.y == y {
			delete(paper, p)
		}
		if p.y > y {
			delete(paper, p)
			paper[point{x: p.x, y: y - (p.y - y)}] = struct{}{}
		}
	}
	return paper
}
