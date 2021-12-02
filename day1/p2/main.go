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

	var window []int

	var isIncreased int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		window = append(window, i)
		if len(window) == 4 {
			if sum(window[1:]) > sum(window[:3]) {
				isIncreased++
			}
			window = window[1:]
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Fprintf(os.Stdout, "answer: %d\n", isIncreased)
}

func sum(w []int) int {
	var s int
	fmt.Fprintf(os.Stdout, "window numbers: %v\n", w)
	for _, n := range w {
		s += n
	}
	fmt.Fprintf(os.Stdout, "window sum: %d\n", s)
	return s
}
