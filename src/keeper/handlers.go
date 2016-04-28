package keeper

import (
  "net/http"
  "encoding/json"
  "common"
)

type IpAddress struct {
  ip string
}

func NewIpHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var newIp IpAddress

    decodeError := decoder.Decode(&newIp)
    common.ErrorCheck(decodeError)

    common.AppendString(newIp.ip,"ips.txt")
}

func RemoveIpHandler(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var ipToRemove IpAddress

  decodeError := decoder.Decode(&ipToRemove)
  common.ErrorCheck(decodeError)

  removeError := common.RemoveString(ipToRemove.ip, "ips.txt")
  common.ErrorCheck(removeError)
}

func QueueInstructionHandler(w http.ResponseWriter, r *http.Request) {
  // add instruction to queue
}
