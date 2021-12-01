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

func main() {
	input, err := readInput("input/1/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	var sums []int
	var lastThree [3]int

	for idx, n := range input {
		if idx >= 2 {
			lastThree[2] = n

			sums = append(sums, sum(lastThree[:]))

			copy(lastThree[:], lastThree[1:])
		} else {
			lastThree[idx%2] = n
		}
	}

	var count = 0
	var last int = sums[0]

	for _, n := range sums {
		if n > last {
			count++
		}

		last = n
	}

	fmt.Printf("Count: %d\n", count)
}
