package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD          = "+"
	MULTIPLY     = "*"
	SYMBOL_INDEX = 4
)

func main() {

	raw, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var matrix [][]string

	rows := strings.Split(string(raw), "\n")
	for _, rawData := range rows {

		rawData = strings.ReplaceAll(rawData, "\r", "")
		parsedData := strings.Split(rawData, "")
		// length := len(parsedData)

		var actualData []string
		current := ""
		for i, item := range parsedData {

			if item != " " {
				// fmt.Println("Current", current, "item", item)

				current += item

				// fmt.Println("Current after", current)
			}

			if (item == " " || i == len(parsedData)-1) && current != "" {
				actualData = append(actualData, current)
				current = ""
				continue
			}

			// if i == len(parsedData)-1 && current != "" {
			// 	actualData = append(actualData, current)
			// 	current = ""
			// 	break
			// }

		}

		matrix = append(matrix, actualData)
	}

	fmt.Println("data", matrix)

	fmt.Println(len(matrix),
		"row 0", len(matrix[0]),
		"row 1", len(matrix[1]),
		"row 2", len(matrix[2]),
		"row 3", len(matrix[3]),
		"row 4", len(matrix[4]),
	)

	total := 0
	for i := 0; i < len(matrix[0]); i++ {

		sum := 0

		first, err := strconv.Atoi(matrix[0][i])
		if err != nil {
			fmt.Println(err)
		}
		second, err := strconv.Atoi(matrix[1][i])
		if err != nil {
			fmt.Println(err)
		}
		third, err := strconv.Atoi(matrix[2][i])
		if err != nil {
			fmt.Println(err)
		}
		fourth, err := strconv.Atoi(matrix[3][i])
		if err != nil {
			fmt.Println(err)
		}

		operand := matrix[4][i]

		if operand == ADD {
			sum = first + second + third + fourth
		} else {
			sum = first * second * third * fourth
		}
		// fmt.Println("sum", sum)
		total += sum
	}

	fmt.Println("Result:", total)
}
