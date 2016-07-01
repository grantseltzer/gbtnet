package main

import (
	"github.com/grantseltzer/Gitaxis/keeper"
)

func main() {
	Authenticate()
	keeper.StartUp()

	/** TODO
	    Check for existing memberlist on startup, ip of one possible node should
	    be somehow packaged during distribution.

	    If one does not exist, a new one should be created. and listeners should
	    be initialized.

	    Have a local registry of active services and ports for quick health checks
	    ---possibly performed by other nodes
	**/
}
