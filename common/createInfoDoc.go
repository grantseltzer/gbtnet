package common

import (
  "os"
)

func createInfoDoc() {
  f, err := os.Create("infodoc.txt")
  if err != nil {
    panic(err)
  }
}
