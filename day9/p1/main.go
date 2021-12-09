package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var heightmap [][]int
	placeHolder := make([]int, 102)
	for index := range placeHolder {
		placeHolder[index] = 9
	}
	heightmap = append(heightmap, placeHolder)
	for scanner.Scan() {
		str := scanner.Text()
		var row []int
		row = append(row, 9)
		for index := range str {
			row = append(row, atoi(string(str[index])))
		}
		row = append(row, 9)
		heightmap = append(heightmap, row)
	}
	heightmap = append(heightmap, placeHolder)

	var lowerPoints int

	for lineIndex := 1; lineIndex < len(heightmap)-1; lineIndex++ {
		line := heightmap[lineIndex]
		for colIndex := 1; colIndex < len(line)-1; colIndex++ {
			ele := heightmap[lineIndex][colIndex]
			if ele == 0 {
				lowerPoints++
				continue
			}
			if ele < heightmap[lineIndex-1][colIndex] && ele < heightmap[lineIndex][colIndex+1] && ele < heightmap[lineIndex+1][colIndex] && ele < heightmap[lineIndex][colIndex-1] {
				lowerPoints += ele + 1
			}
		}
	}
	fmt.Fprintf(os.Stdout, "answer: %d\n", lowerPoints)
}

func atoi(r string) int {
	i, _ := strconv.Atoi(r)
	return i
}
