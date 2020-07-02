/*
Package internal manages the server logic
*/
package internal

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	WorkQueue chan interface{}
)

//Dispatcher will take care of the incoming requests
func Process(waitGroup *sync.WaitGroup, queue int) {

	fmt.Fprintln(os.Stdout, "Starting injest dispatcher")
	defer fmt.Fprintln(os.Stdout, "Stopping injest dispatcher")
	defer waitGroup.Done()

	//time tracker
	startTimer := true

	for {
		//pull an entry from the queue
		task := <- WorkQueue

		switch task := task.(type) {
		case [][]string:
			go func(task [][]string, startTimer bool) {
				Handle().Routines += 1
				if startTimer {
					Handle().Start = time.Now()
					startTimer = false
				}
				IngestCsvData(task)
			}(task, startTimer)
		case bool:
			if task {
				go func(startTimer bool) {
					Report()
					Handle().Routines = 0
					startTimer = true
				}(startTimer)
			}
		}
	}
}