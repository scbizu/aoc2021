package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	input = "SHHBNFBCKNHCNOSHHVFF"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	polymerMap := make(map[string]string)

	for scanner.Scan() {

		parts := strings.Split(scanner.Text(), " -> ")

		polymerMap[parts[0]] = parts[1]
	}

	growMap := make(map[string]int)

	for index := range input {
		if index > 0 {
			growMap[input[index-1:index+1]]++
		}
	}

	fmt.Fprintf(os.Stdout, "initial: growMap: %v\n", growMap)
	for i := 1; i < 41; i++ {
		growMap = growInMap(polymerMap, growMap)
		fmt.Fprintf(os.Stdout, "after %d grow: growMap: %v\n", i, growMap)
	}

	max, min := countGrowMapMaxMinWords(growMap)

	fmt.Fprintf(os.Stdout, "answer: %d\n", max-min)
}

func grow(input string, ploymerMap map[string]string) string {

	newInput := fmt.Sprintf("%c%s%c", input[0], ploymerMap[input], input[1])
	return newInput
}

func growInMap(polymerMap map[string]string, growMap map[string]int) map[string]int {
	newGroupMap := make(map[string]int)
	for k, v := range growMap {
		newInput := grow(k, polymerMap)
		newGroupMap[newInput[:2]] += v
		newGroupMap[newInput[1:]] += v
	}
	return newGroupMap
}

func countGrowMapMaxMinWords(growMap map[string]int) (int, int) {
	wordsMap := make(map[string]int)
	for k, v := range growMap {
		wordsMap[string(k[1])] += v
	}
	wordsMap[string(input[0])]++
	fmt.Fprintf(os.Stdout, "wordsMap: %v\n", wordsMap)
	var max, min int
	for _, v := range wordsMap {
		if v > max {
			max = v
		}
		if v < min || min == 0 {
			min = v
		}
	}
	return max, min

}
