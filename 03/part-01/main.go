package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// inputs := []string{
	// 	"987654321111111",
	// 	"811111111111119",
	// 	"234234234234278",
	// 	"818181911112111",
	// }

	content, err := os.ReadFile("./data.txt")
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	data := string(content)
	inputs := strings.Fields(data)

	sum := 0
	for _, input := range inputs {

		data := formatInput(input)
		res := findGreatestJoltage(data)
		fmt.Println("Response", res)
		sum += res
	}

	fmt.Println("Result", sum)

}

func findGreatestJoltage(data []int) int {
	greatest, index := findGreatestNumber(data[:len(data)-1])

	second, _ := findGreatestNumber(data[index+1:])

	number := fmt.Sprintf("%d%d", greatest, second)

	parsed, _ := strconv.Atoi(number)
	return parsed
}

func formatInput(input string) []int {
	data := strings.Split(input, "")

	var numbers []int

	for _, text := range data {
		parsed, _ := strconv.Atoi(text)
		numbers = append(numbers, parsed)
	}
	return numbers
}

func findGreatestNumber(numbers []int) (greatest, index int) {

	for i, number := range numbers {
		if number > greatest {
			greatest = number
			index = i
		}

	}
	return greatest, index

}
