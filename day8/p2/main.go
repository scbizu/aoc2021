package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var answer int
	for scanner.Scan() {
		pattern := make(map[string]string)
		parts := strings.Split(scanner.Text(), " | ")
		dgs := strings.Split(parts[0], " ")
		dgs = append(dgs, strings.Split(parts[1], " ")...)
		var oneP, fourP, sevenP string
		for _, dg := range dgs {
			sdg := sortDigit(dg)
			if _, ok := pattern[sdg]; ok {
				continue
			}
			switch len(dg) {
			case 2:
				pattern[sortDigit(dg)] = "1"
				oneP = sortDigit(dg)
			case 4:
				pattern[sortDigit(dg)] = "4"
				fourP = sortDigit(dg)
			case 3:
				pattern[sortDigit(dg)] = "7"
				sevenP = sortDigit(dg)
			case 7:
				pattern[sortDigit(dg)] = "8"
			}
		}

		for _, dg := range dgs {
			sdg := sortDigit(dg)
			switch len(dg) {
			case 5:
				co := strIntersectCount(sdg, oneP)
				cf := strIntersectCount(sdg, fourP)
				cs := strIntersectCount(sdg, sevenP)
				switch {
				case co == 2 && cf == 3 && cs == 3:
					pattern[sortDigit(dg)] = "3"
				case co == 1 && cf == 3 && cs == 2:
					pattern[sortDigit(dg)] = "5"
				case co == 1 && cf == 2 && cs == 2:
					pattern[sortDigit(dg)] = "2"
				}
			case 6:
				co := strIntersectCount(sdg, oneP)
				cf := strIntersectCount(sdg, fourP)
				cs := strIntersectCount(sdg, sevenP)
				switch {
				case co == 1 && cf == 3 && cs == 2:
					pattern[sortDigit(dg)] = "6"
				case co == 2 && cf == 4 && cs == 3:
					pattern[sortDigit(dg)] = "9"
				case co == 2 && cf == 3 && cs == 3:
					pattern[sortDigit(dg)] = "0"
				}
			}
		}

		log.Printf("pattern: %v\n", pattern)
		rdgs := strings.Split(parts[1], " ")
		var str strings.Builder
		for _, rdg := range rdgs {
			str.WriteString(pattern[sortDigit(rdg)])
		}
		answer += atoi(str.String())
	}

	fmt.Fprintf(os.Stdout, "answer: %d\n", answer)
}

func atoi(r string) int {
	i, _ := strconv.Atoi(r)
	return i
}

func sortDigit(dg string) string {
	bs := []byte(dg)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	return string(bs)
}

func strIntersectCount(a, b string) int {
	var c int
	for i := 0; i < len(a); i++ {
		if strings.Contains(b, string(a[i])) {
			c++
		}
	}
	return c
}
