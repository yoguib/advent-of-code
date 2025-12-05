package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	data := string(content)
	parsed := strings.Fields(data)

	var matrix [][]string

	roll := "@"
	empty := "."
	rowLength := 0
	columnLength := len(parsed)

	result := 0
	for _, rawRow := range parsed {
		parsedRow := strings.Split(rawRow, "")
		rowLength = len(parsedRow)
		matrix = append(matrix, parsedRow)
	}
	isSet := true
	for isSet {
		removed := false

		for i, row := range matrix {

			for j, column := range row {
				count := 0
				if column == empty {
					continue
				}

				// check top
				if i != 0 && matrix[i-1][j] == roll {
					count++
				}

				// check diagonal
				if (i != 0 && j != 0) && matrix[i-1][j-1] == roll {
					count++
				}

				// check left
				if j != 0 && matrix[i][j-1] == roll {
					count++
				}

				// check right
				if j+1 < rowLength && matrix[i][j+1] == roll {
					count++
				}

				// check diagonal
				if (j+1 < rowLength && i != 0) && matrix[i-1][j+1] == roll {
					count++
				}

				// check bottom
				if (i+1 < columnLength) && matrix[i+1][j] == roll {
					count++
				}

				// check bottom
				if (i+1 < columnLength && j != 0) && matrix[i+1][j-1] == roll {
					count++
				}

				// check bottom
				if (i+1 < columnLength && j+1 < rowLength) && matrix[i+1][j+1] == roll {
					count++
				}

				if count < 4 {
					removed = true
					matrix[i][j] = empty
					result++
				}
			}
		}

		if !removed {
			isSet = false
		}
	}

	if result != 13 {
		fmt.Println("You got this wrong")
	}

	fmt.Println("Result", result)
}
