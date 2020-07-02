/*
Package internal manages the server logic
*/
package internal

import (
	"fmt"
	"os"
	"testing"
)

type testAgeConvert struct {
	age Age
	ageMap AgeMap
	person Person
}

func TestGetPerson(t *testing.T) {

	p1 := Person{FName: "Jane", LName: "Doe", Age: 44}
	p2 := Person{FName: "John", LName: "Doe", Age: 49}
	ps1 := Persons{p1}
	ps2 := Persons{p2}

	ageMapTest := make(AgeMap)
	ageMapTest[p1.Age] = &ps1
	ageMapTest[p2.Age] = &ps2


	//set examples to run
	tests := []testAgeConvert {
		{0, nil, Person{}},						// 0
		{0, AgeMap{}, Person{}},				// 1
		{44, ageMapTest, p1},					// 2
		{55, ageMapTest, Person{}},				// 4
	}

	//check the table tests
	for index, test := range tests {
		result, _ := test.age.getPerson(test.ageMap)
		//check if we got the same result
		if result != test.person{
			fmt.Fprintf(os.Stdout, "FAILED: index=%v, expected=%v, got=%v\n", index, test.person, result)
		} else {
			fmt.Fprintf(os.Stdout, "PASSED: index=%v\n", index)
		}
	}
}
