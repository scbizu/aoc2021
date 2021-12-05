package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var h, depth, aim int

	for scanner.Scan() {
		ps := strings.Split(scanner.Text(), " ")
		switch ps[0] {
		case "up":
			n, _ := strconv.Atoi(ps[1])
			aim -= n
			if aim < 0 {
				aim = 0
			}
		case "down":
			n, _ := strconv.Atoi(ps[1])
			aim += n
		case "forward":
			n, _ := strconv.Atoi(ps[1])
			h += n
			depth += aim * n
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Fprintf(os.Stdout, "answer: [%d*%d] = %d", h, depth, h*depth)
}
