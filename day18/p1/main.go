package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

type pair struct {
	children []*pair
	parent   *pair
	value    int
}

func add(p, p2 *pair) *pair {
	newParent := &pair{}
	if p == nil {
		return p2
	}
	if p2 == nil {
		return p
	}
	newParent.children = append(newParent.children, p, p2)
	p.parent = newParent
	p2.parent = newParent
	return newParent
}

func (p *pair) depth() int {
	var depth int
	if len(p.children) > 0 {
		depth++
	}
	var childDepth int
	for _, c := range p.children {
		if childDepth < c.depth() {
			childDepth = c.depth()
		}
	}
	depth += childDepth
	return depth
}

func (p *pair) explode() {
	depth := p.depth()
	if depth <= 4 {
		return
	}
	var currentDepth int
	ps := p.children
	for {
		var nextChildren []*pair
		currentDepth++
		if currentDepth == depth-1 {
			// do explosion
			for _, p := range ps {
				if len(p.children) == 0 {
					continue
				}
				for idx, p2 := range p.children {
					switch idx {
					case 0:
						p2.mergeToLeft(p2.value)
					case 1:
						p2.mergeToRight(p2.value)
					}
					p2.parent = nil
				}
				p.children = nil
				return
			}
		}
		nextChildren = append(nextChildren, ps...)

		ps = getChildren(nextChildren)
	}
}

func getChildren(ps []*pair) []*pair {
	var children []*pair
	for _, p := range ps {
		children = append(children, p.children...)
	}
	return children
}

func (p *pair) mergeToLeft(v int) {
	for {
		if p.parent == nil {
			return
		}
		if p.parent.children[0] != p {
			p = p.parent.children[0]
			break
		}
		p = p.parent
	}

	for len(p.children) > 0 {
		p = p.children[1]
	}
	p.value += v
}
func (p *pair) mergeToRight(v int) {
	for {
		if p.parent == nil {
			return
		}
		if p.parent.children[1] != p {
			p = p.parent.children[1]
			break
		}
		p = p.parent
	}

	for len(p.children) > 0 {
		p = p.children[0]
	}
	p.value += v
}

func (p *pair) spawnChildren(v1, v2 int) {
	p.children = append(p.children, &pair{value: v1, parent: p}, &pair{value: v2, parent: p})
}

func split(v int) (int, int) {
	middle := v / 2
	if v%2 == 0 {
		return middle, middle
	}
	return middle, middle + 1
}

func (p *pair) split() bool {
	if p.value >= 10 {
		p.spawnChildren(split(p.value))
		p.value = 0
		return true
	}
	if p.children == nil || len(p.children) == 0 {
		return false
	}
	for _, pc := range p.children {
		if ok := pc.split(); ok {
			return ok
		}
	}
	return false
}

func (p *pair) isSplitable() bool {
	if p.value >= 10 {
		return true
	}
	if p.children == nil {
		return false
	}
	for _, pc := range p.children {
		if pc.isSplitable() {
			return true
		}
	}
	return false
}

func (p *pair) isBalance() bool {
	// need explosion
	if p.depth() > 4 {
		return false
	}
	// need split
	if p.isSplitable() {
		return false
	}
	return true
}

func (p *pair) sum() int {
	if len(p.children) == 0 {
		return p.value
	}
	var sum int
	for index, c := range p.children {
		if index == 0 {
			sum += 3 * c.sum()
		}
		if index == 1 {
			sum += 2 * c.sum()
		}
	}
	return sum
}

func (p *pair) rebalanceUntil() {

	for {
		if p.isBalance() {
			return
		}
		println()
		p.print()
		println()
		for p.depth() > 4 {
			print("explode:")
			p.explode()
			p.print()
			println()
		}
		print("split:")
		p.split()
		p.print()
		println()
	}
}

func (p *pair) print() {
	if len(p.children) == 0 {
		print(p.value)
		return
	}
	for index, pc := range p.children {
		if index == 0 {
			print("[")
		}
		pc.print()
		if index == 0 {
			print(",")
		}
		if index == 1 {
			print("]")
		}
	}
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var root *pair
	for scanner.Scan() {
		a := []byte(scanner.Text())
		for i := len(a)/2 - 1; i >= 0; i-- {
			opp := len(a) - 1 - i
			a[i], a[opp] = a[opp], a[i]
		}
		root = add(root, lineToPair(string(a)))
		root.rebalanceUntil()
		root.print()

		println()
	}
	fmt.Fprintf(os.Stdout, "answer: %d", root.sum())
}

func lineToPair(line string) *pair {
	s := &stack{List: list.List{}}
	for _, c := range line {
		switch c {
		case '[':
			ps := s.pop2()
			parent := &pair{}
			parent.children = append(parent.children, ps...)
			ps[0].parent = parent
			ps[1].parent = parent
			s.push(parent)
		default:
			if 0 <= c-'0' && c-'0' <= 9 {
				s.push(&pair{value: int(c - '0')})
			}
		}
	}

	return s.Remove(s.Back()).(*pair)
}

type stack struct {
	list.List
}

func (s *stack) push(p *pair) {
	s.PushBack(p)
}

func (s *stack) pop2() []*pair {
	if s.Len() < 2 {
		panic("stack is empty")
	}
	p1 := s.Remove(s.Back()).(*pair)
	p2 := s.Remove(s.Back()).(*pair)
	return []*pair{p1, p2}
}
