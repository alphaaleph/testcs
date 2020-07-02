/*
Package internal manages the server logic
*/
package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	fnameTitle = "fname"
	lnameTitle = "lname"
	ageTitle = "age"
)

//validateFile checks the file header to see if it is a valid file
func validateFile(record []string) bool {
	if strings.ToLower(strings.TrimSpace(record[0])) == fnameTitle &&
		strings.ToLower(strings.TrimSpace(record[1])) == lnameTitle &&
		strings.ToLower(strings.TrimSpace(record[2])) == ageTitle  {
		return true
	}
	return false
}

func IngestCsvData(records [][]string) {


	//not a valid file
	if (!validateFile(records[0])) {
		return
	}

	//file validated, continue reading
	//loop thru the records
	for _, record := range records[1:] {
		age, _ := strconv.Atoi(strings.TrimSpace(record[2]))
		person := Person{
			FName: strings.TrimSpace(record[0]),
			LName: strings.TrimSpace(record[1]),
			Age:   Age(age),
		}

		Handle().add(person)
	}
}

func Report() {
	fmt.Fprintf(os.Stdout, "\n\nCSV REPORT\n")

	//report average info
	avgAge := Handle().Ages.getAverageAge()
	fmt.Fprintln(os.Stdout, "Average age:", avgAge)
	if avgPerson, avgErr := avgAge.getPerson(Handle().AgeMap); avgErr == nil {
		fmt.Fprintln(os.Stdout, "Average person:", avgPerson.FName, avgPerson.LName)
	} else {
		fmt.Fprintln(os.Stdout, "Average person: none")
	}
	fmt.Fprintf(os.Stdout, "\n")

	//report median info
	medAge := Handle().Ages.getMedianAge()
	fmt.Fprintln(os.Stdout, "Median age:", medAge)
	if medPerson, medErr := medAge.getPerson(Handle().AgeMap); medErr == nil {
		fmt.Fprintln(os.Stdout, "Median person:", medPerson.FName, medPerson.LName)
	} else {
		fmt.Fprintln(os.Stdout, "Median person: none")
	}
	fmt.Fprintf(os.Stdout, "\n")

	//report elapsed time
	fmt.Fprintln(os.Stdout, "Elapsed time:", time.Since(Handle().Start).Seconds(), "seconds")
	fmt.Fprintln(os.Stdout, "Number of routines:", Handle().Routines)
}