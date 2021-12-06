package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const matchedToken = "M"

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var numbers []string
	var dataset [][]string
	scanner := bufio.NewScanner(f)

	var line int
	for scanner.Scan() {
		line++
		if len(scanner.Text()) == 0 {
			continue
		}
		if line == 1 {
			numbers = strings.Split(scanner.Text(), ",")
			continue
		}
		dataset = append(dataset, strings.Split(scanner.Text(), " "))
	}

	rawDs := make([][]string, len(dataset))
	for i := range dataset {
		rawDs[i] = make([]string, len(dataset[i]))
		copy(rawDs[i], dataset[i])
	}

	for _, number := range numbers {
		lok, ok, lIdx, _ := resolveDatasetMatch(dataset, number)
		if lok {
			n := sumUnmarked(dataset[(lIdx/5)*5:(lIdx/5+1)*5], rawDs[(lIdx/5)*5:(lIdx/5+1)*5])
			fmt.Fprintf(os.Stdout, "answer: %d x %d = %d", atoi(number), n, atoi(number)*n)
			break
		}
		if ok {
			n := sumUnmarked(dataset[(lIdx/5)*5:(lIdx/5+1)*5], rawDs[(lIdx/5)*5:(lIdx/5+1)*5])
			fmt.Fprintf(os.Stdout, "answer: %d x %d = %d", atoi(number), n, atoi(number)*n)
			break
		}
	}
}

func sumUnmarked(d [][]string, raw [][]string) int {
	var sn int
	for lIdx := range d {
		for idx := range d[lIdx] {
			if d[lIdx][idx] != matchedToken {
				sn += atoi(raw[lIdx][idx])
			}
		}
	}
	return sn
}

func resolveDatasetMatch(d [][]string, k string) (bool, bool, int, int) {
	for lIdx, line := range d {
		for idx, n := range line {
			if n == matchedToken {
				continue
			}
			if atoi(n) == atoi(k) {
				d[lIdx][idx] = matchedToken
			}
			lok, ok := amIWin(d, lIdx, idx)
			if lok || ok {
				return lok, ok, lIdx, idx
			}
		}
	}
	return false, false, 0, 0
}

func atoi(r string) int {
	i, _ := strconv.Atoi(r)
	return i
}

func amIWin(d [][]string, lIdx, idx int) (bool, bool) {
	win := true
	// line
	// fix height pos(lIdx)
	// will return is unMatched
	for _, ele := range d[lIdx] {
		if ele != matchedToken {
			return false, false
		}
	}
	if win {
		return true, false
	}
	lineScopeStart, lineScopeEnd := lIdx/5*5, (lIdx/5+1)*5
	// col
	for _, l := range d[lineScopeStart:lineScopeEnd] {
		if l[idx] != matchedToken {
			return false, false
		}
	}
	return false, true
}
