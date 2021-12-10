package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
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

	cs := make(map[byte]int)
SCAN:
	for scanner.Scan() {
		stack := stack{list.List{}}
		line := scanner.Text()
		for i := range line {
			switch line[i] {
			case '<', '(', '[', '{':
				stack.PushFront(byte(line[i]))
			case '>', ')', ']', '}':
				if stack.Len() == 0 {
					cs[byte(line[i])]++
					goto SCAN
				}
				if ok := pair(stack.Front().Value.(byte), line[i]); !ok {
					cs[byte(line[i])]++
					log.Printf("got unexpected %c", line[i])
					goto SCAN
				}
				// pop
				stack.Remove(stack.Front())
			}
		}
	}

	var answer int

	for c, n := range cs {
		switch c {
		case '>':
			answer += n * 25137
		case ')':
			answer += n * 3
		case ']':
			answer += n * 57
		case '}':
			answer += n * 1197
		}
	}

	fmt.Fprintf(os.Stdout, "%d\n", answer)
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
