/*
Package internal manages the server logic
*/
package internal

import (
	"errors"
	"fmt"
	"os"
)

// Age is a uint type that holds a person age
type Age uint

//AgeMap will keep track of the ages and groups of persons belonging to that age
type AgeMap map[Age]*Persons

// toUint8 converts Age to uint
func (age Age) toUint() uint {
	return uint(age)
}

// toAge converts uint to Age
func (age *Age) toAge(value uint) {
	*age = Age(value)
}

//getPerson retrieves a person from the list that matches the age, if it can't find a person, it will throw an error
func (age Age) getPerson(ageMap AgeMap) (Person, error) {

	//check if the age map was init
	if ageMap == nil {
		errText := fmt.Sprintf("age map was not initialized")
		fmt.Fprintln(os.Stdout, errText)
		return Person{}, errors.New(errText)
	}

	//check if age map is empty
	if len(ageMap) <= 0 {
		errText := fmt.Sprintf("age map is empty")
		fmt.Fprintln(os.Stdout, errText)
		return Person{}, errors.New(errText)
	}

	//if we find persons matching the age, return the first one in the list
	if persons, ok := ageMap[age]; ok {
		return (*persons)[0], nil
	}

	//there are no persons matching the age
	errText := fmt.Sprintf("no matching person")
	fmt.Fprintln(os.Stdout, errText)
	return Person{}, errors.New(errText)
}