package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cave struct {
	next []*cave
	name string
}

var caveMap map[string]*cave = make(map[string]*cave)

var answer []string

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		from, to := parts[0], parts[1]

		switch {
		case from == "start":
			addCaveToMap(from, to)
		case to == "start":
			addCaveToMap(to, from)
		default:
			addCaveToMap(from, to)
			addCaveToMap(to, from)
		}
	}
	printCaveMap()

	c := caveMap["start"]

	visitFrom(c, []string{})

	fmt.Fprintf(os.Stdout, "answer: %v\n", answer)
	fmt.Fprintf(os.Stdout, "answer: %v\n", len(answer))

}

func addCaveToMap(from, to string) {
	if fromCave, ok := caveMap[from]; ok {
		fromCave.next = append(fromCave.next, &cave{name: to})
	} else {
		fromCave = &cave{name: from}
		fromCave.next = append(fromCave.next, &cave{name: to})
		caveMap[from] = fromCave
	}
}

func visitFrom(c *cave, chain []string) []string {
	fmt.Fprintf(os.Stdout, "on cave: %v\n", c.name)
	if c.name == "end" {
		edchain := append(chain, c.name)
		answer = append(answer, strings.Join(edchain, "-"))
		return chain
	}
	if c.name[0] >= 'a' && c.name[0] <= 'z' && strings.Contains(strings.Join(chain, ","), c.name) {
		edchain := append(chain, c.name)
		fmt.Fprintf(os.Stdout, "end chain: %v\n", edchain)
		// end
		return chain
	}
	chain = append(chain, c.name)
	nextCave := caveMap[c.name]
	for _, cave := range nextCave.next {
		visitFrom(cave, chain)
	}
	return chain
}

func printCaveMap() {
	for _, c := range caveMap {
		fmt.Fprintf(os.Stdout, "%s: ", c.name)
		for _, n := range c.next {
			fmt.Fprintf(os.Stdout, "%s, ", n.name)
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
}
