package gitaxis

import (
	"fmt"

	"github.com/hashicorp/memberlist"
)

// func createCluster() memberlist.Memberlist {
//
// }

// IP Address should be packaged with node on distribution
// func joinCluster(ip string) memberlist.Memberlist {
//   memberList, joinError := list.Join(ip)
//   if joinError != nil {
//     //  handle it
//   }
//   common.ErrorCheck(joinError)
//   return memberList
// }

func printMembership(list memberlist.Memberlist) {
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}
}
