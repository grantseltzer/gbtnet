// package main
//
// import (
//   "fmt"
//   "github.com/hashicorp/memberlist"
// )
//
// func main() {
//   list, err := memberlist.Create(memberlist.DefaultLocalConfig())
//   n, err := list.Join([]string{"127.0.0.1"})
//   fmt.Println(n)
//   if err != nil {
//       panic("Failed to join cluster: " + err.Error())
//   }
//
// }
