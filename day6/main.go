package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	p1Days = 81
	p2Days = 257
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var dataSet []string
	for scanner.Scan() {
		dataSet = strings.Split(scanner.Text(), ",")
	}
	state := make(map[int]int)
	for _, d := range dataSet {
		state[atoi(d)] += 1
	}
	log.Printf("initial state: %v", state)
	for i := 1; i < p2Days; i++ {
		state = spawn(state)
		log.Printf("after %d days, state: %v", i, state)
	}
	var answer int
	for _, s := range state {
		answer += s
	}
	fmt.Fprintf(os.Stdout, "answer: %d", answer)
}

func atoi(r string) int {
	i, _ := strconv.Atoi(r)
	return i
}

func spawn(state map[int]int) map[int]int {
	newState := make(map[int]int)
	for i := 8; i >= 0; i-- {
		if _, ok := state[i]; !ok {
			continue
		}
		j := i - 1
		if j < 0 {
			newState[6] += state[i]
			newState[8] += state[i]
			continue
		}
		newState[j] = state[i]
	}
	return newState
}
