/**
  REST API for receiving instructions
**/

package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func getJsonResponse() ([]byte, error) {

}
