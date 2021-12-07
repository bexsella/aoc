package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(path string) (input []int, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	strs := strings.Split(scanner.Text(), ",")

	for _, str := range strs {
		i, err := strconv.Atoi(str)

		if err == nil {
			input = append(input, i)
		}
	}

	return input, err
}

func optimalStepsWithCrabTech(input []int) int {
	var sum, avg int

	for _, num := range input {
		avg += num
	}

	avg /= len(input)

	for _, num := range input {
		diff := int(math.Abs(float64(num - avg)))
		step := (diff * (diff + 1)) / 2
		sum += step
	}

	return sum
}

func optimalSteps(input []int) int {
	var sorted []int = make([]int, len(input))
	var sum int

	copy(sorted, input)
	sort.Ints(sorted)

	median := sorted[len(sorted)/2]

	for _, num := range sorted {
		sum += int(math.Abs(float64(num - median)))
	}

	return sum
}

func Execute() {
	input, err := readInput("input/7/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Optimal human index: %d\n", optimalSteps(input))
	fmt.Printf("Optimal crab index: %d\n\n", optimalStepsWithCrabTech(input))
}
