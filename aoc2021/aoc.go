package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(path string) (input []int, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err == nil {
			input = append(input, num)
		}
	}

	return input, nil
}

func sum(data []int) int {
	var sum int

	for _, n := range data {
		sum += n
	}

	return sum
}

func countIncreases(input []int) int {
	var count = 0
	var last int = input[0]

	for _, n := range input {
		if n > last {
			count++
		}

		last = n
	}

	return count
}

func collateWindowIncreases(input []int, windowRange int) (output []int) {
	var sums []int
	var lastInRange = make([]int, windowRange)

	for idx, n := range input {
		if idx >= windowRange-1 {
			lastInRange[windowRange-1] = n

			sums = append(sums, sum(lastInRange[:]))

			copy(lastInRange[:], lastInRange[1:])
		} else {
			lastInRange[idx%(windowRange-1)] = n
		}
	}

	return sums
}

func main() {
	input, err := readInput("input/1/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Part 1 Count: %d\n", countIncreases(input))
	fmt.Printf("Part 2 Count: %d\n", countIncreases(collateWindowIncreases(input, 3)))
}
