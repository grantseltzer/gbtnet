package whisper

import (
  "net/http"
  "encoding/json"
  "bytes"
  "fmt"
  "time"
)

type Ip struct {
  IpAddress string
  insertionTime time.Time
}

// MOVE TO KEEPER, the post request that this function returns can be added to a queue for whisper to pick up off of to execute
func postIP(ip string) (*http.Request) {
    insertTime := time.Now()
    newIp := &Ip{ip, insertTime}
    jsonNewIp, jsonError := json.Marshal(newIp)
    if (jsonError != nil) {
        panic(jsonError)
    }
    postRequest, postError := http.NewRequest("POST", ip, bytes.NewBuffer(jsonNewIp))
    if (postError != nil) {
        panic(postError)
    }
    return postRequest
}

/***
    postIP (and other post requests) should be part of a system that adds requests
    to a queue in keeper that whisper will pick up to fullfil

    postIP usage:

    func main() {
    ourRequest := postIP("http://localhost:8080/")
    fmt.Println(ourRequest)
    client := &http.Client{}
    resp, err := client.Do(ourRequest)
    if err != nil {
        panic(err)
    }
    fmt.Println("Response: ", resp)
}

***/
