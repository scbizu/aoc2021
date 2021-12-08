package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var answer int
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		dgs := strings.Split(parts[1], " ")
		for _, dg := range dgs {
			switch len(dg) {
			case 2, 4, 3, 7:
				answer++
			}
		}
	}

	fmt.Fprintf(os.Stdout, "answer: %d\n", answer)
}
