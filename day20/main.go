package main

import (
	"fmt"
	"os"
	"strings"
)

type pixelState uint8

const (
	stateLight pixelState = iota
	stateDark
)

func (ps pixelState) Byte() byte {
	switch ps {
	case stateLight:
		return '#'
	case stateDark:
		return '.'
	default:
		panic("invalid pixel state")
	}
}

var table algorithmString = initAlgorithmString()

type algorithmString string

func initAlgorithmString() algorithmString {
	return algorithmString(checksum)
}

func (a algorithmString) Index(i int) pixelState {
	switch a[i] {
	case '.':
		return stateDark
	case '#':
		return stateLight
	default:
		panic("invalid character")
	}
}

type image [][]byte

func buildImage(x, y int) image {
	var newImage [][]byte
	for i := 0; i < x; i++ {
		var line []byte
		for j := 0; j < y; j++ {
			line = append(line, '.')
		}
		newImage = append(newImage, line)
	}
	return newImage
}

func binaryToDecimal(bin []byte) int {
	var decimal int
	for _, b := range bin {
		b = toBin(b)
		decimal = decimal*2 + int(b)
	}
	return decimal
}

func getBackground(n int) byte {
	if []byte(table)[0] == '.' {
		return '.'
	}
	// '#'
	switch n % 2 {
	case 1:
		return []byte(table)[0]
	case 0:
		return []byte(table)[len(table)-1]
	}
	panic("can not reach")
}

func (i image) expand(n int) (image, image) {
	newImage := buildImage(len(i)+2, len(i[0])+2)
	shadow := buildImage(len(i)+2, len(i[0])+2)
	for x := range newImage {
		for y := range newImage[x] {
			if x == 0 || y == 0 || x == len(newImage)-1 || y == len(newImage[x])-1 {
				newImage[x][y] = getBackground(n)
				shadow[x][y] = getBackground(n)
			}
		}
	}
	for x := range i {
		for y := range i[x] {
			newImage[x+1][y+1] = i[x][y]
			shadow[x+1][y+1] = i[x][y]
		}
	}
	return newImage, shadow
}

func toBin(b byte) byte {
	switch b {
	case '.':
		return 0
	case '#':
		return 1
	}
	panic("invalid byte")
}

func (i image) countLit() int32 {
	var count int32
	for _, line := range i {
		for _, p := range line {
			if p == '#' {
				count++
			}
		}
	}
	return count
}

func (i image) relight(n int) image {
	newImage, shadow := i.expand(n)
	placeholder := getBackground(n)
	println("new image:")
	newImage.print()
	println()

	for y := range newImage {
		for x := range newImage[y] {
			var binary []byte
			// #1
			if x-1 >= 0 && y-1 >= 0 {
				binary = append(binary, newImage[y-1][x-1])
			} else {
				binary = append(binary, placeholder)
			}
			// #2
			if y-1 >= 0 {
				binary = append(binary, newImage[y-1][x])
			} else {
				binary = append(binary, placeholder)
			}
			// #3
			if x+1 < len(newImage[y]) && y-1 >= 0 {
				binary = append(binary, newImage[y-1][x+1])
			} else {
				binary = append(binary, placeholder)
			}
			// #4
			if x-1 >= 0 {
				binary = append(binary, newImage[y][x-1])
			} else {
				binary = append(binary, placeholder)
			}
			// #5
			binary = append(binary, newImage[y][x])
			// #6
			if x+1 < len(newImage[y]) {
				binary = append(binary, newImage[y][x+1])
			} else {
				binary = append(binary, placeholder)
			}
			// #7
			if x-1 >= 0 && y+1 < len(newImage) {
				binary = append(binary, newImage[y+1][x-1])
			} else {
				binary = append(binary, placeholder)
			}
			// #8
			if y+1 < len(newImage) {
				binary = append(binary, newImage[y+1][x])
			} else {
				binary = append(binary, placeholder)
			}
			// #9
			if x+1 < len(newImage[y]) && y+1 < len(newImage) {
				binary = append(binary, newImage[y+1][x+1])
			} else {
				binary = append(binary, placeholder)
			}
			shadow[y][x] = table.Index(binaryToDecimal(binary)).Byte()
		}
	}
	return shadow
}

func (i image) print() {
	for _, line := range i {
		for _, b := range line {
			print(string(b))
		}
		println()
	}
}

func main() {
	var readIn image
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var row []byte
		for _, b := range line {
			switch string(b) {
			case ".":
				row = append(row, '.')
			case "#":
				row = append(row, '#')
			}
		}
		readIn = append(readIn, row)
	}
	fmt.Fprintln(os.Stdout, "initial image:")
	readIn.print()
	println()
	for i := 0; i < 50; i++ {
		readIn = readIn.relight(i)
		println("result image:")
		readIn.print()
		println()
	}

	fmt.Fprintf(os.Stdout, "lit: %d\n", readIn.countLit())
}
