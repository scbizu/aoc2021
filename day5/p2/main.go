package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

var result map[point]int = make(map[point]int)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), " -> ")
		n1 := strings.Split(pairs[0], ",")
		n2 := strings.Split(pairs[1], ",")
		if n1[0] != n2[0] && n1[1] != n2[1] {
			log.Printf("%v -> %v", n1, n2)
			switch (atoi(n2[0]) - atoi(n1[0])) / (atoi(n2[1]) - atoi(n1[1])) {
			case 1:
				if small(n1[0], n2[0]) == atoi(n1[0]) {
					j := atoi(n1[1])
					for i := atoi(n1[0]); i <= atoi(n2[0]); i++ {
						log.Printf("point: [%d,%d]", i, j)
						result[point{i, j}]++
						j++
					}
					break
				}
				j := atoi(n1[1])
				for i := atoi(n1[0]); i >= atoi(n2[0]); i-- {
					log.Printf("point: [%d,%d]", i, j)
					result[point{i, j}]++
					j--
				}

			case -1:
				if small(n1[0], n2[0]) == atoi(n1[0]) {
					j := atoi(n1[1])
					for i := atoi(n1[0]); i <= atoi(n2[0]); i++ {
						log.Printf("point: [%d,%d]", i, j)
						result[point{i, j}]++
						j--
					}
					break
				}
				j := atoi(n1[1])
				for i := atoi(n1[0]); i >= atoi(n2[0]); i-- {
					log.Printf("point: [%d,%d]", i, j)
					result[point{i, j}]++
					j++
				}
			}
			continue

		}
		if n1[0] == n2[0] {
			s := small(n1[1], n2[1])
			b := big(n1[1], n2[1])
			for i := s; i <= b; i++ {
				result[point{atoi(n1[0]), i}]++
			}
		}

		if n1[1] == n2[1] {
			s := small(n1[0], n2[0])
			b := big(n1[0], n2[0])
			for i := s; i <= b; i++ {
				result[point{i, atoi(n1[1])}]++
			}

		}
	}

	var answer int

	log.Println("cal")

	for _, count := range result {
		if count <= 1 {
			continue
		}
		// log.Printf("point: %v\n", p)
		answer += 1
	}

	fmt.Fprintf(os.Stdout, "answer: %d\n", answer)
}

func big(a, b string) int {
	if atoi(a) > atoi(b) {
		return atoi(a)
	}
	return atoi(b)
}

func small(a, b string) int {
	if atoi(a) < atoi(b) {
		return atoi(a)
	}
	return atoi(b)
}

func atoi(r string) int {
	i, _ := strconv.Atoi(r)
	return i
}

func smallInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bigInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
