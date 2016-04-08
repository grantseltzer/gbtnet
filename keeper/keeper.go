package main

import (
  "fmt"
)

func main() {
  originalString := "BWAHHHHHHHHHHHHHHHHHHHHHHHHHHHH"
  fmt.Println(originalString)

  encryptedString := Encrypt(originalString, "1234567891234567")
  fmt.Println(encryptedString)

  decryptedString := Decrypt(encryptedString, "1234567891234567")
  fmt.Println(decryptedString)
}
