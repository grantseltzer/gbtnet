package keeper

import (
	"net/http"

	"github.com/grantseltzer/Gitaxis/common"
)

/*******************************************************************************
  Keeper is a microservice that will run on each block in blockchain that will
  access the data stored on the node and serve requests from the microservices
  in the local node.
*******************************************************************************/

// StartUp boots the keeper service
func StartUp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/newIp", newIPHandler)
	mux.HandleFunc("/removeIp/", removeIPHandler)

	listenError := http.ListenAndServe(":8481", mux)
	common.ListenErrorCheck(listenError, "keeper")
}
