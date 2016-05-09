package common

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
  "bufio"
)

func AppendString(newLine string, fileName string) {
    file, openError := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
    ErrorCheck(openError)
    bytesWritten, writeError := file.Write([]byte(newLine + "\n"))
    fmt.Println(bytesWritten)
    ErrorCheck(writeError)
}

func RemoveString(lineToRemove string, fileName string) {
  inputFromFile, inputError := ioutil.ReadFile(fileName)
  ErrorCheck(inputError)

  allLines := strings.Split(string(inputFromFile), "\n")

  for i, line := range allLines {
    if strings.Contains(line, lineToRemove) {
      allLines = append(allLines[:i], allLines[i+1:]...)
    }
  }

  outputToFile := strings.Join(allLines, "\n")
  writeError := ioutil.WriteFile(fileName, []byte(outputToFile), 0600)
  ErrorCheck(writeError)
}

func ReadFirstLine(fileName string) string {
  infile, inError := os.Open(fileName)
  ErrorCheck(inError)
  defer infile.Close()

  scanner := bufio.NewScanner(infile)
  scanner.Scan()
  text := scanner.Text()
  return text
}
