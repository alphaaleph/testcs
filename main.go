package main

import (
	"flag"
	"fmt"
	"github.com/alphaaleph/testcs/api"
	"github.com/alphaaleph/testcs/internal"
	"os"
	"sync"
)

func main() {
	fmt.Fprintln(os.Stdout, "CS SERVICE START")

	//get flags -- this should be read from configuration, but didnt want to add external toml
	port := flag.Int("port", 8282, "application port")
	queue := flag.Int("queue", 10000, "work queue buffer")

	//init the handler which will manage the actual work, and the work queue
	internal.WorkQueue = make(chan interface{}, *queue)
	internal.Handle().NewHandler()

	//init the dispatcher which manages the work queue
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go internal.Process(&waitGroup, *queue)

	//start the server
	server := &api.Server{Port: *port}
	server.Start()

	waitGroup.Wait()
	fmt.Fprintln(os.Stdout, "CS SERVICE STOP")
}

