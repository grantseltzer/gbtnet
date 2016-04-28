package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "bytes"
  "common"
)

// MOVE TO KEEPER, the post request that this function returns can be added to a queue for whisper to pick up off of to execute
func postIP(ip string, targetIp string) (*http.Request) {
    newIp := &common.Ip{ip}
    jsonNewIp, jsonError := json.Marshal(newIp)
    if (jsonError != nil) {
        panic(jsonError)
    }
    postRequest, postError := http.NewRequest("POST", targetIp, bytes.NewBuffer(jsonNewIp))
    if (postError != nil) {
        panic(postError)
    }
    return postRequest
}


// ^ USAGE:
func main() {
    ourRequest := postIP("grantisAwesome", "http://localhost:8481/newIp")
    fmt.Println(ourRequest)
    client := &http.Client{}
    resp, err := client.Do(ourRequest)
    if err != nil {
        panic(err)
    }
    fmt.Println("Response: ", resp)
}
