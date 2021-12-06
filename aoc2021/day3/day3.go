package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(path string) (input []int, bitSize int, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, 0, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bitSize = len(scanner.Text())
		num, _ := strconv.ParseInt(scanner.Text(), 2, len(scanner.Text())+1)
		input = append(input, int(num))
	}

	return
}

func bitCounter(input []int, bitSize int) (output [][]int) {
	// Clean this mess up:
	output = make([][]int, 2)
	output[0], output[1] = make([]int, bitSize), make([]int, bitSize)

	// Count occurances of 1 or 0
	for _, n := range input {
		for b := 0; b < bitSize; b++ {
			output[n>>b&1][b]++
		}
	}

	return
}

func filterByBit(input []int, bit int, bitSize int, filter func(a, b int) int) (filteredResult int, output []int) {
	counter := bitCounter(input, bitSize)
	var filteredBit int = filter(counter[0][bit], counter[1][bit])

	for _, n := range input {
		if n>>bit&1 == filteredBit {
			output = append(output, n)
		}
	}

	if len(output) > 1 {
		return filterByBit(output, bit-1, bitSize, filter)
	}

	return output[0], nil
}

func calcLifeSupportRating(input []int, bitSize int) int {
	oxy, _ := filterByBit(input, bitSize-1, bitSize, func(b0, b1 int) int {
		if b1 >= b0 {
			return 1
		}

		return 0
	})

	co2, _ := filterByBit(input, bitSize-1, bitSize, func(b0, b1 int) int {
		if b0 <= b1 {
			return 0
		}

		return 1
	})

	return oxy * co2
}

func calcPowerConsumption(input []int, bitSize int) int {
	var power int
	var mask = ((1 << bitSize) - 1)

	counter := bitCounter(input, bitSize)

	// Build the resulting number
	for b := 0; b < bitSize; b++ {
		if counter[1][b] > counter[0][b] {
			power |= 1 << b
		}
	}

	return power * (power ^ mask)
}

func Execute() {
	input, bitSize, err := readInput("input/3/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Power Consumption: %d\n", calcPowerConsumption(input, bitSize))
	fmt.Printf("Life Support Rating: %d\n\n", calcLifeSupportRating(input, bitSize))
}
