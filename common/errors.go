package common

import "fmt"

func ErrorCheck(err error) {
  if err != nil {
    panic(err)
  }
}

func ListenErrorCheck(err error, projectName string) {
  fmt.Println("Interuptted Listen and Serve:  ", projectName)
}
