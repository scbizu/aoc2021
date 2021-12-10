package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
)

type stack struct {
	list.List
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var allPoints []int

SCAN:
	for scanner.Scan() {
		stack := stack{list.List{}}
		line := scanner.Text()
		log.Printf("%s", line)
		for i := range line {
			switch line[i] {
			case '<', '(', '[', '{':
				stack.PushFront(byte(line[i]))
			case '>', ')', ']', '}':
				if stack.Len() == 0 {
					goto SCAN
				}
				if ok := pair(stack.Front().Value.(byte), line[i]); !ok {
					goto SCAN
				}
				// pop
				stack.Remove(stack.Front())
			}
		}
		var points int
		// range stack from front
		for e := stack.Front(); e != nil; e = e.Next() {
			points *= 5

			switch e.Value.(byte) {
			case '<':
				points += 4
			case '(':
				points += 1
			case '[':
				points += 2
			case '{':
				points += 3
			}
		}

		allPoints = append(allPoints, points)
	}

	sort.Slice(allPoints, func(i, j int) bool {
		return allPoints[i] < allPoints[j]
	})

	fmt.Fprintf(os.Stdout, "answer: %d\n", allPoints[(len(allPoints)-1)/2])
}

func pair(c1, c2 byte) bool {
	switch c1 {
	case '<':
		return c2 == '>'
	case '(':
		return c2 == ')'
	case '[':
		return c2 == ']'
	case '{':
		return c2 == '}'
	}
	return false
}
