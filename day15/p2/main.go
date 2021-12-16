package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

const (
	matrixSize = 100
)

type point struct {
	x, y  int
	value int
}

type xyPoint struct {
	x, y int
}

var visitedMap = make(map[xyPoint]int)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var chiton = make([][]int, matrixSize)

	var j int

	for scanner.Scan() {
		line := scanner.Text()
		for i, word := range line {
			chiton[i] = append(chiton[i], atoi(word))
		}
		j++
	}
	chiton = chitonExpansion(chiton)
	ps := &points{}
	newPQ(ps)
	var x, y, v int
	for {
		fmt.Fprintf(os.Stdout, "x: %d,y: %d\n", x, y)
		if x == len(chiton)-1 && y == len(chiton[x])-1 {
			break
		}
		around := getAroundPoints(chiton, point{x, y, v})
		for _, p := range around {
			fmt.Fprintf(os.Stdout, "around: x: %d,y: %d\n", p.x, p.y)
			v2, ok := visitedMap[xyPoint{p.x, p.y}]
			if !ok || p.value < v2 {
				heap.Push(ps, p)
				visitedMap[xyPoint{p.x, p.y}] = p.value
			}
		}
		p := heap.Pop(ps).(point)

		fmt.Fprintf(os.Stdout, "pop: %v\n", p)
		x, y, v = p.x, p.y, p.value
	}
	fmt.Fprintf(os.Stdout, "answer: %d\n", v)
}

func atoi(s rune) int {
	return int(s - '0')
}

func getAroundPoints(chiton [][]int, p point) []point {
	var ps []point
	v := p.value

	if p.x > 0 {
		ps = append(ps, point{p.x - 1, p.y, v + chiton[p.x-1][p.y]})
	}

	if p.x < len(chiton)-1 {
		ps = append(ps, point{p.x + 1, p.y, chiton[p.x+1][p.y] + v})
	}
	if p.y > 0 {
		ps = append(ps, point{p.x, p.y - 1, chiton[p.x][p.y-1] + v})
	}
	if p.y < len(chiton[p.x])-1 {
		ps = append(ps, point{p.x, p.y + 1, chiton[p.x][p.y+1] + v})
	}
	return ps
}

func chitonExpansion(raw [][]int) [][]int {
	newChiton := make([][]int, len(raw)*5)

	for index := range newChiton {
		newChiton[index] = make([]int, len(raw[0])*5)
	}

	for x := range raw {
		for y := range raw[x] {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					newChiton[x+(i*len(raw))][y+(len(raw[x])*j)] = up(raw[y][x] + i + j)
				}
			}
		}
	}

	return newChiton
}

func up(v int) int {
	if v > 9 {
		return v - 9
	}
	return v
}

// points impls heap.Interface
type points []point

func (p points) Len() int { return len(p) }
func (p points) Less(i, j int) bool {
	return p[i].value < p[j].value
}
func (p points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p *points) Push(x interface{}) {
	*p = append(*p, x.(point))
}
func (p *points) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

type priorityQueue struct {
	heap.Interface
}

func newPQ(p *points) priorityQueue {
	heap.Init(p)
	return priorityQueue{p}
}
