/*
Package api manages the REST calls functionality
*/
package api

import (
	"fmt"
	"net/http"
	"os"
)

//Service
type Server struct {
	Port 	int
}

//Run the app
func (server Server) Start() {

	//api setup
	http.HandleFunc("/upload", server.loadData)
	http.HandleFunc("/done", server.done)

	//set the listener
	port := fmt.Sprintf(":%d", server.Port)
	err := http.ListenAndServe(port, nil)

	//if we fail to start get out
	if err != nil {
		fmt.Fprintf(os.Stdout, "Service failed to start: %v", err)
		os.Exit(1)
	}
}
