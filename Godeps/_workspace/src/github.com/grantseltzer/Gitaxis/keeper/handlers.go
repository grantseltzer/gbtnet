package keeper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grantseltzer/Gitaxis/common"
)

func newIPHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newIP common.IP

	decodeError := decoder.Decode(&newIP)
	common.ErrorCheck(decodeError)
	fmt.Println(newIP.IPAddress)

	common.AppendString(newIP.IPAddress, "ips.txt")
}

func removeIPHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ipToRemove common.IP

	decodeError := decoder.Decode(&ipToRemove)
	common.ErrorCheck(decodeError)

	common.RemoveString(ipToRemove.IPAddress, "keeper/ips.txt")
}

// QueueInstructionHandler blah
func QueueInstructionHandler(w http.ResponseWriter, r *http.Request) {
	// add instruction to queue
}
