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
	var last, isIncreased, index int
	for scanner.Scan() {
		index++
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(os.Stdout, "read in: %d, last: %d\n", i, last)
		if index == 1 {
			last = i
			continue
		}
		if i > last {
			isIncreased++
		}
		last = i
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Fprintf(os.Stdout, "answer: %d\n", isIncreased)
}
