/*
Package internal manages the server logic
*/
package internal

import (
	"fmt"
	"math"
	"os"
	"sort"
)

// getAverageAge will return the average age of all collected persons
func (ages Ages) getAverageAge() Age {

	//check if we have any ages
	if (ages != nil) && (len(ages) > 0) {

		//init the average age
		averageAge := Age(0)

		//add all the ages
		for _, v := range ages {
			averageAge += v
		}

		//divide by the amount of ages, and use the go round call
		tempAge := float64(averageAge) / float64(len(ages))
		return Age(math.Round(tempAge))
	}

	fmt.Fprintln(os.Stdout, "no average ages found")
	return 0
}


// getMedianAge will return the median age of all collected persons
func (ages Ages) getMedianAge() Age {

	//check if we have any ages
	if (ages != nil) && (len(ages) > 0) {

		//sort the ages
		sort.Sort(ages)
		midIndex := len(ages) / 2

		//if there is an odd number of entries, pick the middle number
		if len(ages) % 2 != 0 {
			return ages[midIndex]
		}

		//if there is an even number of entries, pick the two middle numbers, add them, and divide by two
		return (ages[midIndex - 1] + ages[midIndex]) / 2
	}

	fmt.Fprintln(os.Stdout, "no median ages found")
	return 0
}