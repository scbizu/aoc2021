package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var c [12]int

	for scanner.Scan() {
		for index, b := range scanner.Bytes() {
			switch b {
			case '0':
				c[index] -= 1
			case '1':
				c[index] += 1
			}
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	var gm, es strings.Builder
	for _, n := range c {
		if n > 0 {
			gm.WriteByte('1')
			es.WriteByte('0')
			continue
		}
		gm.WriteByte('0')
		es.WriteByte('1')
	}
	n, _ := strconv.ParseInt(gm.String(), 2, 64)
	m, _ := strconv.ParseInt(es.String(), 2, 64)
	fmt.Fprintf(os.Stdout, "answer[%d x %d]: %d", n, m, n*m)
}
