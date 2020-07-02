package internal

import (
	"sync"
	"time"
)

var (
	handler 	*Handler
	once      	sync.Once
)

//Handler will manage the incoming data and keep track of it
type Handler struct {
	Ages
	AgeMap
	Start time.Time
	Routines int
	sync.Mutex
}

//Handle service singleton
func Handle() *Handler {
	once.Do(func() {
		handler = &Handler{}
	})
	return handler
}

//NewHandler initializes the Handler
func (handler *Handler) NewHandler() {
	handler.Ages = Ages{}
	handler.AgeMap = make(AgeMap)
	handler.Routines = 0
}

func (handler *Handler) add(person Person) {

	//check that the person has info
	if person == (Person{}) || person.Age == 0 {
		return
	}

	//add the age to the ages slice
	handler.Ages = append(handler.Ages, person.Age)

	//check if we have the age in the age map and add the person
	if persons, ok := handler.AgeMap[person.Age]; ok {
		*persons = append(*persons, person)
		return
	}

	//add a new age and person slice
	newPersons := Persons{person}
	handler.AgeMap[person.Age] = &newPersons
}