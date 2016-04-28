package keeper

import (
  "net/http"
  "encoding/json"
  "common"
  "fmt"
)

func NewIpHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("minor success\n\n\n")
    decoder := json.NewDecoder(r.Body)
    var newIp common.Ip

    decodeError := decoder.Decode(&newIp)
    common.ErrorCheck(decodeError)
    fmt.Println(newIp.IpAddress)

    common.AppendString(newIp.IpAddress,"ips.txt")
}

func RemoveIpHandler(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var ipToRemove common.Ip

  decodeError := decoder.Decode(&ipToRemove)
  common.ErrorCheck(decodeError)

  common.RemoveString(ipToRemove.IpAddress, "keeper/ips.txt")
}

func QueueInstructionHandler(w http.ResponseWriter, r *http.Request) {
  // add instruction to queue
}
