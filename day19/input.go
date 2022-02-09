package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

var scannerReader []string

func readInput() [][]point {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	buffer := bytes.Buffer{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, "---") {
			buffer.WriteString(line + "\n")
		} else {
			scannerReader = append(scannerReader, buffer.String())
			buffer.Reset()
		}
	}

	if buffer.Len() > 0 {
		scannerReader = append(scannerReader, buffer.String())
		buffer.Reset()
	}

	var input [][]point
	for _, pstr := range scannerReader {
		if pstr == "" {
			continue
		}
		var ps []point
		for _, line := range strings.Split(pstr, "\n") {
			if line == "" {
				continue
			}
			parts := strings.Split(line, ",")
			ps = append(ps, point{
				x: atoi(parts[0]),
				y: atoi(parts[1]),
				z: atoi(parts[2]),
			})
		}
		input = append(input, ps)
	}
	return input
}
