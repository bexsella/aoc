package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Movement struct {
	mtype int
	value int
}

type MoveInput struct {
	move string
	unit int
}

const (
	HORIZ_TYPE = iota
	DEPTH_TYPE = iota
)

var (
	Movements = map[string]Movement{
		"forward": {HORIZ_TYPE, 1},
		"down":    {DEPTH_TYPE, 1},
		"up":      {DEPTH_TYPE, -1},
	}
)

func ReadInput(path string) (input []MoveInput, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var text string
		var value int

		fmt.Sscanf(scanner.Text(), "%s %d", &text, &value)
		input = append(input, MoveInput{text, value})
	}

	return input, nil
}

func ProcessMovements(moves []MoveInput, useAim bool) int {
	var horiz, depth, aim int

	for _, m := range moves {
		switch Movements[m.move].mtype {
		case HORIZ_TYPE:
			horiz += m.unit

			if useAim {
				depth += aim * m.unit
			}
		case DEPTH_TYPE:
			if useAim {
				if Movements[m.move].value > 0 {
					aim += m.unit
				} else {
					aim -= m.unit
				}
			} else {
				depth += m.unit * Movements[m.move].value
			}
		}
	}

	return horiz * depth
}

func Execute() {
	input, err := ReadInput("input/2/input")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Dive Position: %d\n", ProcessMovements(input, false))
	fmt.Printf("Final Position %d\n\n", ProcessMovements(input, true))
}
