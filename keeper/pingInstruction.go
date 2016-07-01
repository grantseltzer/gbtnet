package keeper

import "fmt"

type pingQueue struct {
	ips []string
}

// this is just a test
func (p pingQueue) pullOffQueue() error {
	for _, ip := range p.ips {
		fmt.Println("PING! ", ip)
	}
	return nil
}
