package main

import (
  "common"
  "fmt"
  "github.com/hashicorp/memberlist"
)

func createCluster() Memberlist {
  
}

// IP Address should be packaged with node on distribution
func joinCluster(ip string) Memberlist {
  memberList, joinError := list.Join(ip)
  if joinError != nil {
    //  handle it
  }
  common.ErrorCheck(joinError)
  return memberList
}

func printMembership(list Memberlist) {
  for _, member := range list.Members() {
    fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
  }
}
