package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("./data.txt")
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	data := string(content)
	inputs := strings.Fields(data)
	// inputs := []string{
	// 	"987654321111111",
	// 	"811111111111119",
	// 	"234234234234278",
	// 	"818181911112111",
	// }

	sum := 0
	for _, input := range inputs {

		data := formatInput(input)
		for len(data) > 12 {

			res := removeSmallest(data)
			data = res
			// fmt.Println("Data result", data)
		}

		finalS := ""
		for _, number := range data {
			finalS += strconv.Itoa(number)
		}
		fmt.Println("Final", finalS)
		final, _ := strconv.Atoi(finalS)
		sum += final

	}

	fmt.Println("Result", sum)

	if sum != 171990312704598 {
		fmt.Println("She is broken")
	} else {
		fmt.Println("Good stuff")
	}

}

func removeSmallest(data []int) []int {
	index := 0
	for i, number := range data {
		if i+2 > len(data) {
			index = i
			break
		}

		if number < data[i+1] {
			index = i
			break
		}
	}

	return slices.Delete(data, index, index+1)
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
