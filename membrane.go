/**
  REST API for receiving instructions
**/

package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type Payload struct {
  Stuff Data
}

type Data struct {
  Url Urls
  Instructions Instructions
}

type Urls map[string]string
type Instructions map[string]string


func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func main() {
  http.HandleFunc("/", serveRest)
  http.ListenAndServe("localhost:8000", nil)
}

func getJsonResponse() ([]byte, error) {
  targets := make(map[string]string)
  targets["jin"] = "1:1:1:1"
  targets["target"] = "1:1:1:1"

  instructions := make(map[string]string)
  instructions["download"] = "this"
  instructions["send"] = "that"

  d := Data{targets, instructions}
  p := Payload{d}
  return json.MarshalIndent(p, "", "  ")
}
