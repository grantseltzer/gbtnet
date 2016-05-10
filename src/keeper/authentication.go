package keeper

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "common"
  "math/rand"
  "time"
)

func Authenticate() {
  fmt.Println("[*] Enter Key:")

  var keyInput string
  fmt.Scanln(&keyInput)
  firstLine := common.ReadFirstLine(".key.hkp")
  compareError := bcrypt.CompareHashAndPassword([]byte(firstLine), []byte(keyInput))

  if compareError != nil {
    denied()
  } else {
    granted()
  }
}

//TODO
func CreateKey() {
  fmt.Println("[*] Enter new key:")
  var keyInput string
  fmt.Scanln(&keyInput)
  generateKey(keyInput)
}

//TODO
func generateKey(password string) {
  hashKey, hashError := bcrypt.GenerateFromPassword([]byte(password), 11)

}

//TODO
func granted() {
    fmt.Println("ACCESS GRANTED")
}

// End Program, check for amount of past attempts.
// If the past attempts are over 10 wipe everything
// If not, increment past attempts
//TODO
func denied() {
    fmt.Println("\nACCESS DENIED")
}
