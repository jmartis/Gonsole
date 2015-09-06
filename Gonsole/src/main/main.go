package main 

import (
	"net/http"
)

type person struct {
	fName string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First Name: " + p.fName))
}

func main() {

	
	personOne := &person{fName: "Jim"}
	http.ListenAndServe(":8081", personOne)
//	myMux := http.NewServeMux()
//	myMux.HandleFunc("/", personOne)
//	http.ListenAndServe(":8080", myMux)
	
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world"))
}

