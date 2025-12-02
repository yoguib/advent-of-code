package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	currentPosition := 50
	zeroCounts := 0

	content, err := os.ReadFile("./data.txt")
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	data := string(content)
	parsed := strings.Fields(data)

	for _, intruction := range parsed {
		finalPosition, counts := getFinalPosition(currentPosition, intruction)
		fmt.Println("result", finalPosition)
		zeroCounts += counts
		currentPosition = finalPosition
	}

	fmt.Println("answer", zeroCounts)
}

func getFinalPosition(currentPosition int, instruction string) (finalPosition int, zeroCounts int) {

	finalPosition = currentPosition
	direction, numberOfSteps := getDirections(instruction)
	fmt.Println("Direction", direction, "steps", numberOfSteps)

	if numberOfSteps == 0 {
		return finalPosition, zeroCounts
	}

	for numberOfSteps > 0 {
		if direction == "R" {
			finalPosition = moveForward(finalPosition)
		} else {
			finalPosition = moveBackwards(finalPosition)
		}

		if finalPosition == 0 {
			zeroCounts++
		}
		numberOfSteps--
	}
	return finalPosition, zeroCounts
}

func getDirections(instruction string) (direction string, numberOfSteps int) {

	var rawSteps string
	for i, data := range instruction {

		if i == 0 {
			direction = string(data)
			continue
		}
		rawSteps = rawSteps + string(data)
	}

	numberOfSteps, _ = strconv.Atoi(rawSteps)

	return direction, numberOfSteps
}

func moveForward(currentPosition int) int {
	if currentPosition == 99 {
		return 0
	}

	return currentPosition + 1
}

func moveBackwards(currentPosition int) int {
	if currentPosition == 0 {
		return 99
	}

	return currentPosition - 1
}
