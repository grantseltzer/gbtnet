package main

import (
  // "common"
  "fmt"
  "github.com/hashicorp/memberlist"
)

func main() {
  list, err := memberlist.Create(memberlist.DefaultLocalConfig())
  if err != nil {
      panic("Failed to create memberlist: " + err.Error())
  }

  var input string
  fmt.Scanln(&input)

  for _, member := range list.Members() {
    fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
  }
}
