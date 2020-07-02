package api

import (
	"encoding/csv"
	"github.com/alphaaleph/testcs/internal"
	"net/http"
)

//LoadData receives and manipulates CSV data
func (server Server) loadData(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//read all the csv records
	records, err := csv.NewReader(r.Body).ReadAll()
	if err != nil || records == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ingest the data
	internal.WorkQueue <- records
	w.WriteHeader(http.StatusOK)
}

//done will print a report after reading all the files
func (server Server) done(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//write the report when done
	internal.WorkQueue <- true
	w.WriteHeader(http.StatusOK)
}
