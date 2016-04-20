package keeper

import (
  "io"
  "net/http"
  "encoding/json"
)

type ipAddress struct {
  ip string
}

// Add ip address to document
func newIpHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var newIp ipAddress

    err := decoder.Decode(&newIp)
    if (err != nil) {
      panic(err) // Need to fix error handling
    }
    AppendString(newIp,"ips.txt")
}

func removeIpHandler(w http.ResponseWriter, r *http.Request) {
  // remove ip address from document
}

func queueInstructionHandler(w http.ResponseWriter, r *http.Request) {
  // add instruction to queue
}
