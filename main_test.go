package main

import (
	"slices"
	"testing"
)

func Test_ParseRanges(t *testing.T) {
	t.Run("valid inputs", func(t *testing.T) {
		input := "11-22,95-115,998-1012"
		expected := []string{"11-22", "95-115", "998-1012"}

		response := parseRanges(input)

		if slices.Compare(expected, response) != 0 {
			t.Error("Something went wrong",
				"Expected", expected,
				"Got", response)
		}
	})
}

func Test_ParseStartAndEnd(t *testing.T) {
	t.Run("valid range", func(t *testing.T) {
		input := "11-22"
		expectedStart := 11
		expectedEnd := 22

		start, end := parseStartAndEnd(input)

		if start != expectedStart {
			t.Error("Something went wrong",
				"expected", expectedStart,
				"got", start,
			)
		}

		if end != expectedEnd {
			t.Error("Something went wrong",
				"expected", expectedEnd,
				"got", expectedEnd,
			)
		}

	})
}

func Test_GetInvalidIdsByRange(t *testing.T) {

	t.Run("simple one", func(t *testing.T) {

		response := getInvalidIdsByRange(11, 22)
		expected := []int{11, 22}

		if !slices.Equal(response, expected) {
			t.Error("something went wrong",
				"expected", expected,
				"got", response)
		}
		t.Log("response", response)
	})

	t.Run("simple dois", func(t *testing.T) {

		response := getInvalidIdsByRange(99, 112)
		expected := []int{99, 111}

		if !slices.Equal(response, expected) {
			t.Error("something went wrong",
				"expected", expected,
				"got", response)
		}
		t.Log("response", response)
	})

	t.Run("Part two three equals", func(t *testing.T) {

		response := getInvalidIdsByRange(998, 1012)
		expected := []int{999, 1010}

		if !slices.Equal(response, expected) {
			t.Error("something went wrong",
				"expected", expected,
				"got", response,
			)
		}
	})

	t.Run("multiple doubles", func(t *testing.T) {

		response := getInvalidIdsByRange(2121212118, 2121212124)
		expected := []int{2121212121}

		if !slices.Equal(response, expected) {
			t.Error("something went wrong",
				"expected", expected,
				"got", response,
			)
		}
	})

	t.Run("multiple triples", func(t *testing.T) {

		response := getInvalidIdsByRange(824824821, 824824827)
		expected := []int{824824824}

		if !slices.Equal(response, expected) {
			t.Error("something went wrong",
				"expected", expected,
				"got", response,
			)
		}
	})
}
