package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) (input [9]int, err error) {
	file, err := os.Open(path)

	if err != nil {
		return input, err
	}

	scanner := bufio.NewScanner(file)
	var strs []string

	if scanner.Scan() {
		strs = strings.Split(scanner.Text(), ",")
	}

	for _, str := range strs {
		num, err := strconv.Atoi(str)

		if err == nil {
			input[num]++
		}
	}

	return input, nil
}

func sumLife(input [9]int) (sum int) {
	for _, nums := range input {
		sum += nums
	}

	return
}

func simulateLifecycle(input [9]int, numDays int) int {
	var lifeTimer [9]int

	copy(lifeTimer[:], input[:])

	for day := 0; day < numDays; day++ {
		var newTimer [9]int

		for i := 0; i < 9; i++ {
			if i == 0 {
				newTimer[8] += lifeTimer[i]
				newTimer[6] += lifeTimer[i]
			} else {
				newTimer[i-1] += lifeTimer[i]
			}
		}

		copy(lifeTimer[:], newTimer[:])
	}

	return sumLife(lifeTimer)
}

func Execute() {
	input, err := readInput("input/6/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Sum of 80 day lifecycle: %d\n", simulateLifecycle(input, 80))
	fmt.Printf("Sum of 256 day lifecycle: %d\n", simulateLifecycle(input, 256))
}
