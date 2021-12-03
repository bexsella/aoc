package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInput(path string) (input []int, bitSize int, err error) {
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

func BitCounter(input []int, bitSize int) (output [][]int) {
	// Clean this mess up:
	output = make([][]int, 2)
	output[0] = make([]int, bitSize)
	output[1] = make([]int, bitSize)

	// Count occurances of 1 or 0
	for _, n := range input {
		for b := 0; b < bitSize; b++ {
			output[n>>b&1][b]++
		}
	}

	return
}

func GetByMax(input []int, bit int, bitSize int) (output []int) {
	counter := BitCounter(input, bitSize)
	var max int

	if counter[1][bit] >= counter[0][bit] {
		max = 1
	} else {
		max = 0
	}

	for _, n := range input {
		if n>>bit&1 == max {
			output = append(output, n)
		}
	}

	if len(output) > 1 {
		return GetByMax(output, bit-1, bitSize)
	}

	return
}

func GetByMin(input []int, bit int, bitSize int) (output []int) {
	counter := BitCounter(input, bitSize)
	var min int

	if counter[0][bit] <= counter[1][bit] {
		min = 0
	} else {
		min = 1
	}

	for _, n := range input {
		if n>>bit&1 == min {
			output = append(output, n)
		}
	}

	if len(output) > 1 {
		return GetByMin(output, bit-1, bitSize)
	}

	return
}

func CalcLifeSupportRating(input []int, bitSize int) int {
	oxy := GetByMax(input, bitSize-1, bitSize)
	co2 := GetByMin(input, bitSize-1, bitSize)
	return oxy[0] * co2[0]
}

func CalcPowerConsumption(input []int, bitSize int) int {
	var power int
	var mask = ((1 << bitSize) - 1)

	counter := BitCounter(input, bitSize)

	// Build the resulting number
	for b := 0; b < bitSize; b++ {
		if counter[1][b] > counter[0][b] {
			power |= 1 << b
		}
	}

	return power * (power ^ mask)
}

func Execute() {
	input, bitSize, err := ReadInput("input/3/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Power Consumption: %d\n", CalcPowerConsumption(input, bitSize))
	fmt.Printf("Life Support Rating: %d\n\n", CalcLifeSupportRating(input, bitSize))
}
