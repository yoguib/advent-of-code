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

		matrix = append(matrix, parsedData)
	}

	fmt.Println(len(matrix),
		"row 0", len(matrix[0]),
		"row 1", len(matrix[1]),
		"row 2", len(matrix[2]),
		"row 3", len(matrix[3]),
		"row 4", len(matrix[4]),
	)

	var numbers []string

	symbolIndex := 0
	sum := 0

	for i := range matrix[0] {
		num := ""

		if matrix[0][i] != " " {
			num += matrix[0][i]
		}
		if matrix[1][i] != " " {
			num += matrix[1][i]
		}
		if matrix[2][i] != " " {
			num += matrix[2][i]
		}
		if matrix[3][i] != " " {
			num += matrix[3][i]
		}

		if num != "" {
			numbers = append(numbers, num)
			num = ""
		}

		if (matrix[0][i] == " " &&
			matrix[1][i] == " " &&
			matrix[2][i] == " " &&
			matrix[3][i] == " " &&
			matrix[4][i] == " ") || i == len(matrix[0])-1 {

			t := 0

			fmt.Println("numners", numbers)
			for i, s := range numbers {
				fmt.Println("number", s, "index", i)
				parsed, err := strconv.Atoi(s)
				if err != nil {
					fmt.Println(err)
				}

				if i == 0 {
					t = parsed
					continue
				}

				if matrix[4][symbolIndex] == ADD {
					t += parsed
				} else {
					t *= parsed
				}
			}

			numbers = []string{}
			sum += t
			symbolIndex = i + 1
			continue
		}

	}
	fmt.Println("Total:", sum)

	return
}
