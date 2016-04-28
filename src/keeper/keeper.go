package keeper

import (
  "net/http"
  "common"
)

/*******************************************************************************
  Keeper is a microservice that will run on each block in blockchain that will
  access the data stored on the node and serve requests from the microservices
  in the local node.
*******************************************************************************/

func StartUp() {
// Start keeper microservice
  mux := http.NewServeMux()
  mux.HandleFunc("/newIp", NewIpHandler)
  mux.HandleFunc("/removeIp/", RemoveIpHandler)

  listenError := http.ListenAndServe(":8481", mux)
  common.ListenErrorCheck(listenError, "keeper")
}
