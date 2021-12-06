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
	dataset = fmtDataSet(dataset)
	var ok bool
	for _, number := range numbers {
		dataset, ok = resolveDatasetMatch(dataset, number)
		if ok {
			n := sumUnmarked(dataset)
			fmt.Fprintf(os.Stdout, "answer: %d x %d = %d", n, atoi(number), n*atoi(number))
			return
		}
	}
}

func fmtDataSet(d [][]string) [][]string {
	for i, l := range d {
		for index, ele := range l {
			if ele == "" {
				d[i] = append(d[i][:index], d[i][index+1:]...)
			}
		}
	}
	return d
}

func sumUnmarked(d [][]string) int {
	log.Printf("%v", d)
	var sn int
	for lIdx := range d {
		for idx := range d[lIdx] {
			if d[lIdx][idx] != matchedToken {
				log.Printf("unmarked: %d\n", atoi(d[lIdx][idx]))
				sn += atoi(d[lIdx][idx])
			}
		}
	}
	return sn
}

func resolveDatasetMatch(d [][]string, k string) ([][]string, bool) {
search:
	for lIdx, line := range d {
		for idx, n := range line {
			if n == matchedToken {
				continue
			}
			if atoi(n) == atoi(k) {
				d[lIdx][idx] = matchedToken
			}
			ok := amIWin(d, lIdx, idx)
			if ok {
				if len(d)/5 == 1 {
					return d, true
				}
				log.Printf("block %d complete with number %d\n", whichBlock(lIdx), atoi(k))
				lineScopeStart, lineScopeEnd := lIdx/5*5, (lIdx/5+1)*5
				log.Printf("block info: %v\n", d[lineScopeStart:lineScopeEnd])
				// leave dataset
				d = append(d[:lineScopeStart], d[lineScopeEnd:]...)
				log.Printf("left block: %d\n", len(d)/5)
				goto search
			}
		}
	}
	return d, false
}

func whichBlock(line int) int {
	return line / 5
}

func atoi(r string) int {
	i, _ := strconv.Atoi(r)
	return i
}

func amIWin(d [][]string, lIdx, idx int) bool {
	win := true
	// line
	// fix height pos(lIdx)
	// will return is unMatched
	for _, ele := range d[lIdx] {
		if ele != matchedToken {
			win = false
			break
			// fallthrough to col
		}
	}

	if win {
		return true
	}

	lineScopeStart, lineScopeEnd := lIdx/5*5, (lIdx/5+1)*5
	// col
	for _, l := range d[lineScopeStart:lineScopeEnd] {
		if l[idx] != matchedToken {
			return false
		}
	}
	return true
}
