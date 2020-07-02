/*
Package internal manages the server logic
*/
package internal

//Person holds the information for one person in the data
type Person struct {
	FName 	string 	`json:"fname"`
	LName	string	`json:"lname"`
	Age				`json:"age"`
}

//Persons is a slice of person that the server will use to group people of the same age
type Persons []Person

func (p Person) add(age chan<- Age, person chan<- Person) {
	//add a person's age to Ages slice to be used for calculations
	age <- p.Age
	//add a person to the person tracker map
	person <- p
}

