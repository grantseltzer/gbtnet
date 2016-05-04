package main

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "bufio"
  "os"
)

func main() {
  p := "Password"
  ps, _ := bcrypt.GenerateFromPassword([]byte(p), 10)

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter pw: ")
  text, _ := reader.ReadString('d') // This is what needs to be fixed
  fmt.Println("p:", []byte(p), "text:", []byte(text))
  errorR := bcrypt.CompareHashAndPassword(ps, []byte(text))

  if errorR != nil {
    fmt.Println("no-match")
  } else {
    fmt.Println("match")
  }
}
