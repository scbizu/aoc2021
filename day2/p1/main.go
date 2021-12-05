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

	var h, v int

	for scanner.Scan() {
		ps := strings.Split(scanner.Text(), " ")
		switch ps[0] {
		case "up":
			n, _ := strconv.Atoi(ps[1])
			v -= n
		case "down":
			n, _ := strconv.Atoi(ps[1])
			v += n
		case "forward":
			n, _ := strconv.Atoi(ps[1])
			h += n
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Fprintf(os.Stdout, "answer: [%d*%d] = %d", h, v, h*v)
}
