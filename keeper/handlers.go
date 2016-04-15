package keeper

import (
  "io"
  "net/http"
)

func newIpHandler(w http.ResponseWriter, r *http.Request) {
  // Add ip address to document
}

func removeIpHandler(w http.ResponseWriter, r *http.Request) {
  // remove ip address from document
}

func queueInstructionHandler(w http.ResponseWriter, r *http.Request) {
  // add instruction to queue
}
