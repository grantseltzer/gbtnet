package keeper

import (
  "io"
  "net/http"
  "encoding/json"
)

type ipAddress struct {
  ip string
}

func newIpHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var newIp ipAddress

    decodeError := decoder.Decode(&newIp)
    errorCheck(decodeError)

    appendError := AppendString(newIp,"ips.txt")
    errorCheck(appendError)
}

func removeIpHandler(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var ipToRemove ipAddress

  decodeError := decoder.Decode(&newIp)
  errorCheck(decodeError)

  removeError := RemoveString(ipToRemove, "ips.txt")
  errorCheck(removeError)
}

func queueInstructionHandler(w http.ResponseWriter, r *http.Request) {
  // add instruction to queue
}
