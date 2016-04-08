package main

import (
  "fmt"
)

/*******************************************************************************
  Keeper is a microservice that will run on each block in blockchain that will
  access the data stored on the node and serve requests from the microservices
  in the local node.
*******************************************************************************/

func main() {

  //Testing encrypt files
  originalString := "THIS IS A TEST ARGHHHHh!"
  fmt.Println(originalString)

  encryptedString := Encrypt(originalString, "1234567891234567")
  fmt.Println(encryptedString)

  decryptedString := Decrypt(encryptedString, "1234567891234567")
  fmt.Println(decryptedString)
}
