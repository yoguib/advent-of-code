package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	data := "269194394-269335492,62371645-62509655,958929250-958994165,1336-3155,723925-849457,4416182-4470506,1775759815-1775887457,44422705-44477011,7612653647-7612728309,235784-396818,751-1236,20-36,4-14,9971242-10046246,8796089-8943190,34266-99164,2931385381-2931511480,277-640,894249-1083306,648255-713763,19167863-19202443,62-92,534463-598755,93-196,2276873-2559254,123712-212673,31261442-31408224,421375-503954,8383763979-8383947043,17194-32288,941928989-941964298,3416-9716"
	parsedData := parseRanges(data)
	var sum atomic.Int64
	var wg sync.WaitGroup

	for _, idRange := range parsedData {
		wg.Add(1)
		go func(ids string) {

			start, end := parseStartAndEnd(idRange)
			response := getInvalidIdsByRange(start, end)

			total := 0

			for _, i := range response {
				total += i
			}
			sum.Add(int64(total))
			wg.Done()
		}(idRange)
	}
	wg.Wait()

	fmt.Println("Result", sum.Load())
}

func parseRanges(input string) []string {
	response := strings.Split(input, ",")

	return response
}

func parseStartAndEnd(IdRange string) (start, end int) {
	data := strings.Split(IdRange, "-")

	start, _ = strconv.Atoi(data[0])
	end, _ = strconv.Atoi(data[1])

	return start, end
}

func getInvalidIdsByRange(start, end int) []int {
	var invalidIds []int
	for i := start; i <= end; i++ {
		asString := strconv.Itoa(i)

		length := len(asString)
		if length%2 == 0 {
			half := length / 2
			firstHalf := asString[:half]
			rest := asString[half:]

			if firstHalf == rest {
				invalidIds = append(invalidIds, i)
				continue
			}

		}

		skipAmount := 0

		for k, data := range asString {
			if k == 0 {
				continue
			}
			if string(data) == string(asString[0]) {
				skipAmount = k
				break
			}
		}
		if skipAmount == 0 {
			continue
		}

		previous := ""
		match := false
		for j := 0; j < length; j += skipAmount {

			if j+skipAmount > length {
				match = false
				break
			}

			if j == 0 {
				previous += asString[:skipAmount]
				continue
			}

			current := asString[j : j+skipAmount]
			if previous == current {
				match = true
			} else {
				match = false
				break
			}

			previous = current
		}

		if match {
			invalidIds = append(invalidIds, i)
			continue
		}

	}

	return invalidIds
}
