package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var all []string

	for scanner.Scan() {
		all = append(all, scanner.Text())
	}
	oxy := all
	co2 := make([]string, len(all))
	copy(co2, all)
	for i := 0; i < 12; i++ {
		oxy = reduce(oxy, i, true)
		co2 = reduce(co2, i, false)
	}

	n, _ := strconv.ParseInt(oxy[0], 2, 64)
	m, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Fprintf(os.Stdout, "answer[%d x %d]: %d", n, m, n*m)
}

func reduce(d []string, offsetIndex int, mode bool) []string {
	if len(d) == 1 {
		return d
	}
	var sub0, sub1 []string
	for _, str := range d {
		for index := range str {
			if index != offsetIndex {
				continue
			}
			switch str[offsetIndex] {
			case '0':
				sub0 = append(sub0, str)
			case '1':
				sub1 = append(sub1, str)
			}
		}
	}
	switch mode {
	case true:
		if len(sub0) > len(sub1) {
			return sub0
		}
		if len(sub0) == len(sub1) {
			return sub1
		}
		return sub1
	case false:
		if len(sub0) > len(sub1) {
			return sub1
		}
		if len(sub0) == len(sub1) {
			return sub0
		}
		return sub0
	}
	panic("do not reach")
}
