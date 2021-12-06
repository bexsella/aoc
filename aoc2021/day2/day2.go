package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Move struct {
	move  string
	value int
}

func readInput(path string) (input []Move, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var text string
		var value int

		fmt.Sscanf(scanner.Text(), "%s %d", &text, &value)
		input = append(input, Move{text, value})
	}

	return input, nil
}

func processMovements(moves []Move, useAim bool) int {
	var horiz, depth, aim int
	var depthAimValues = map[string]int{"up": -1, "down": 1}

	for _, m := range moves {
		switch m.move {
		case "forward":
			horiz += m.value

			if useAim {
				depth += aim * m.value
			}
		case "up", "down":
			if useAim {
				aim += m.value * depthAimValues[m.move]
			} else {
				depth += m.value * depthAimValues[m.move]
			}
		}
	}

	return horiz * depth
}

func Execute() {
	input, err := readInput("input/2/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Dive Position: %d\n", processMovements(input, false))
	fmt.Printf("Final Position %d\n\n", processMovements(input, true))
}
