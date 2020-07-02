package internal

import (
	"fmt"
	"os"
	"testing"
)

type testAges struct {
	ages Ages
	ageTest Age
}

func TestGetAverageAge(t *testing.T) {

	//set examples to run
	tests := []testAges {
		{Ages{}, 0},							// 0
		{Ages{89}, 89},							// 1
		{Ages{100, 18}, 59},					// 2
		{Ages{25,25,25,25}, 25},				// 3
		{Ages{12,55,71,34,8}, 36},				// 4
		{Ages{45,90,12,56,23,11}, 40},			// 5
		{Ages{45,23,41,41,18,3}, 29},			// 6
	}

	//check the table tests
	for index, test := range tests {
		result := test.ages.getAverageAge()
		//check if we got the same result
		if result != test.ageTest {
			fmt.Fprintf(os.Stdout, "FAILED: index=%v, expected=%v, got=%v\n", index, test.ageTest, result)
		} else {
			fmt.Fprintf(os.Stdout, "PASSED: index=%v\n", index)
		}
	}
}

func TestGetMedianAge(t *testing.T) {

	//set examples to run
	tests := []testAges {
		{Ages{}, 0},							// 0
		{Ages{89}, 89},							// 1
		{Ages{100, 18}, 59},					// 2
		{Ages{25,25,25,25}, 25},				// 3
		{Ages{12,55,71,34,8}, 34},				// 4
		{Ages{45,90,12,56,23,11}, 34},			// 5
		{Ages{45,23,41,41,18,3}, 32},			// 6
	}

	//check the table tests
	for index, test := range tests {
		result := test.ages.getMedianAge()
		//check if we got the same result
		if result != test.ageTest {
			fmt.Fprintf(os.Stdout, "FAILED: index=%v, expected=%v, got=%v\n", index, test.ageTest, result)
		} else {
			fmt.Fprintf(os.Stdout, "PASSED: index=%v\n", index)
		}
	}
}
