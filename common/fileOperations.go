package common

import (
  "os"
  "io/ioutil"
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
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
    if (err != nil) {
      return err
    }

    bytesWritten, err := file.Write([]byte(newLine + "\n"))
    if err == nil && bytesWritten < len(ip) {
    			err = io.ErrShortWrite
		}

    return err
}
