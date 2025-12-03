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

	// 169172125774058
	// 171990312704578
	var sum uint64
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
		sum += uint64(final)
		// [9883332222222 2 2 2]
		// 234234234234278
		//   4 34234234278
		//  34 34 34234278
	}

	fmt.Println("Result", sum)

}

func removeSmallest(data []int) []int {
	index := 0
	for i, number := range data {
		if number == 1 {
			index = i
			break
		}

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
