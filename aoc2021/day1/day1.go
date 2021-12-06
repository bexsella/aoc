package day1

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

func sum(data []int) (sum int) {
	for _, n := range data {
		sum += n
	}

	return sum
}

func countIncreases(input []int) (count int) {
	var last int = input[0]

	for _, n := range input {
		if n > last {
			count++
		}

		last = n
	}

	return count
}

func sumRange(input []int, windowRange int) (output []int) {
	var lastInRange = make([]int, windowRange)

	for idx, n := range input {
		if idx >= windowRange-1 {
			lastInRange[windowRange-1] = n

			output = append(output, sum(lastInRange[:]))

			copy(lastInRange[:], lastInRange[1:])
		} else {
			lastInRange[idx] = n
		}
	}

	return
}

func Execute() {
	input, err := readInput("input/1/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Part 1: %d\n", countIncreases(input))
	fmt.Printf("Part 2: %d\n\n", countIncreases(sumRange(input, 3)))
}
