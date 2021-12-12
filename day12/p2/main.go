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

var answer map[string]struct{} = make(map[string]struct{})

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

	var litters []string

	for cave := range caveMap {
		if cave == "start" || cave == "end" {
			continue
		}
		if cave[0] >= 'a' && cave[0] <= 'z' {
			litters = append(litters, cave)
		}
	}

	for _, l := range litters {
		visitFrom(c, []string{}, l, 2)
	}

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

func visitFrom(c *cave, chain []string, twiceWord string, twiceCount int) []string {
	fmt.Fprintf(os.Stdout, "on cave: %v\n", c.name)
	if c.name == "end" {
		edchain := append(chain, c.name)
		answer[strings.Join(edchain, "-")] = struct{}{}
		return chain
	}
	if c.name[0] >= 'a' && c.name[0] <= 'z' {
		switch {
		case twiceWord == c.name:
			twiceCount--
			if twiceCount == -1 {
				edchain := append(chain, c.name)
				fmt.Fprintf(os.Stdout, "end chain: %v\n", edchain)
				// end
				return chain

			}
		case strings.Contains(strings.Join(chain, ","), c.name):
			edchain := append(chain, c.name)
			fmt.Fprintf(os.Stdout, "end chain: %v\n", edchain)
			// end
			return chain
		}
	}
	chain = append(chain, c.name)
	nextCave := caveMap[c.name]
	for _, cave := range nextCave.next {
		visitFrom(cave, chain, twiceWord, twiceCount)
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
