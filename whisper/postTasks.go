package whisper

import (
  "net/http"
  "encoding/json"
  "bytes"
  "time"
)

type Ip struct {
  IpAddress string
  insertionTime time.Time
}

func sendNewIp(ip string) {
  insertTime := time.Now()
  newIp := &Ip{ip, insertTime}
  buf, _ := json.Marshal(newIp)
  body := bytes.NewBuffer(buf)
  r, _ := http.Post("http://localhost:8080", "text/json", body)
  // Write r to some log?
  // response, _ := ioutil.ReadAll(r.Body)
}
