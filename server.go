package main

/*
* Main file for fibonacci server
* valid endpoints:
*
* http://localhost:8844/current
* http://localhost:8844/next
* http://localhost:8844/previous
*
* One assumption I made while complying with the steps provided in the instructions
* is that I actually advance the value when calling "Next", so after that call I update the current value too
* in that sense at a specific point, if you call next, you will see the next value but if you call current
* right away you will see the same value in order to keep the sequence on the "next" call
* could have been implemented in a slightly different way but since I did not have more test runs expected
* I assumed it like that.
*/

import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "encoding/json"
)

type errorStruct struct {
  Code int
  Message string
}

//main "global" vars to store the state of the server
var currentValue = 0
var nextValue = 1
var previousValue = 0

//calculates Fibonacci values for next, and assigs all values properly
func calculateFibonacci() int{
  a, b := currentValue, nextValue
  previousValue = currentValue

  currentValue, nextValue = b, a+b

  return b
}

func formatError(w http.ResponseWriter, err error, code int){
  var errorData errorStruct

  errorData.Code = code
  errorData.Message = err.Error()

  log.Printf("An error accured: %v", err)
  w.WriteHeader(code)

  w.Header().Set("Content-Type", "application/json")

  marshalledContent, _ := json.MarshalIndent(errorData, "", "\t")
  w.Write(marshalledContent)
}

//**** Endpoints and Views ***** ///

func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Gabriel Vargas Pex Assignment")
}

func current(w http.ResponseWriter, r *http.Request){

  fmt.Fprintf(w, "Current Value: %d ", currentValue)
}

func next(w http.ResponseWriter, r *http.Request){
  nextVal := calculateFibonacci()

  fmt.Fprintf(w, "Next Value: %d ", nextVal)
}

func previous(w http.ResponseWriter, r *http.Request){

  fmt.Fprintf(w, "Previous Value: %d ", previousValue)
}

//Main function and Router
func main() {

  router := mux.NewRouter()
  router.HandleFunc("/", homeHandler).Methods("GET")

  //Requested End Points
  router.HandleFunc("/current", current).Methods("GET")
  router.HandleFunc("/next", next).Methods("GET")
  router.HandleFunc("/previous", previous).Methods("GET")

  //Listen to this port
  log.Fatal(http.ListenAndServe(":8844", router))
}





