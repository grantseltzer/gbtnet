package keeper

import (
  "io"
  "net/http"
)

func newIp(w http.ResponseWriter, r *http.Request) {
  // Add ip address to document
}

func removeIp(w http.ResponseWriter, r *http.Request) {
  // remove ip address from document
}

func queueInstruction(w http.ResponseWriter, r *http.Request) {
  // add instruction to queue
}
