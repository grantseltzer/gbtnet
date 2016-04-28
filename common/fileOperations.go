package common

import (
  "os"
  "io/ioutil"
  "strings"
)

type Document struct {
  Body []byte
}

func (p *Document, fileName string) Save() error {
  return ioutil.WriteFile(fileName, p.Body, 0600)
}

func Load(fileName string) (*Document, err) {
  body, err := ioutil.ReadFile(filename)
  if (err != nil) {
    return nil, err
  }
  return &Document{Body: body}, nil
}

func AppendString(newLine string, fileName string) error {
    file, openError := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
    errorCheck(openError)

    bytesWritten, writeError := file.Write([]byte(newLine + "\n"))
    errorCheck(writeError)

    return err
}

func RemoveString(lineToRemove string, fileName string) error {
  inputFromFile, _ := ioutil.ReadFile(fileName)
  errorCheck(inputError)

  allLines := strings.Split(string(inputFromFile), "\n")

  for i, line := range allLines {
    if strings.Contains(line, lineToRemove) {
      allLines = append(allLines[:i], allLines[i+1:]...)
    }
  }

  outputToFile := strings.Join(allLines, "\n")
  writeError := ioutil.WriteFile(fileName, []byte(outputToFile), 0600)
  return writeError
}
