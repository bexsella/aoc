package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInput(path string) (input []int, err error) {
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

func Sum(data []int) (sum int) {
	for _, n := range data {
		sum += n
	}

	return sum
}

func CountIncreases(input []int) (count int) {
	var last int = input[0]

	for _, n := range input {
		if n > last {
			count++
		}

		last = n
	}

	return count
}

func SumRange(input []int, windowRange int) (output []int) {
	var lastInRange = make([]int, windowRange)

	for idx, n := range input {
		if idx >= windowRange-1 {
			lastInRange[windowRange-1] = n

			output = append(output, Sum(lastInRange[:]))

			copy(lastInRange[:], lastInRange[1:])
		} else {
			lastInRange[idx] = n
		}
	}

	return
}

func Execute() {
	input, err := ReadInput("input/1/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Part 1: %d\n", CountIncreases(input))
	fmt.Printf("Part 2: %d\n\n", CountIncreases(SumRange(input, 3)))
}
