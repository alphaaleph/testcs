/*
Package internal manages the server logic
*/
package internal

// Ages is a slice that collects all the persons ages and will be used to calculate the average and median
type Ages []Age

// Len is the number of elements in the collection.
func (ages Ages) Len() int { return len(ages) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (ages Ages) Less(i, j int) bool { return ages[i] < ages[j] }

// Swap swaps the elements with indexes i and j.
func (ages Ages) Swap(i, j int) { ages[i], ages[j] = ages[j], ages[i] }
