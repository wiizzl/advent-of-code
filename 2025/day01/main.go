package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1(input string) int {
	currentPosition := 50
	zeroCounter := 0

	rotations := strings.SplitSeq(input, "\n")

	for rotation := range rotations {
		distance64, err := strconv.ParseInt(rotation[1:], 10, 0)
		distance := int(distance64)

		if err != nil {
			continue
		}

		if rotation[0] == 'L' {
			distance = -distance
		}

		currentPosition = (currentPosition + distance) % 100
		if currentPosition < 0 {
			currentPosition += 100
		}

		if currentPosition == 0 {
			zeroCounter++
		}
	}

	return zeroCounter
}

func part2(input string) int {
	return 0
}

func main() {
	fmt.Println("Part1:", part1(input))
	fmt.Println("Part2:", part2(input))
}
