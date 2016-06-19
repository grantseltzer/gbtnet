package common

import (
  "encoding/json"
)

type Payload struct {
  Instructions Instruction
}

type Instruction struct {
  IP             IPS
  InstructionType InstructionTypes
  Specification   Specifications
}

type IPS map[string]string
type InstructionTypes map[string]string
type Specifications map[string]string

func WriteInst(SourceIp, TargetIp, InstructionType, Specification string) ([]byte, error) {

  Ips := make(map[string]string)
  Ips["Source"] = SourceIp
  Ips["Target"] = TargetIp

  InsType := make(map[string]string)
  InsType["InstructionType"] = InstructionType

  Spec := make(map[string]string)
  Spec["Specification"] = Specification

  i := Instruction{Ips, InsType, Spec}
  p := Payload{i}
  return json.Marshal(p)
}

func ReadInst(slice []byte) (Payload) {
  var payload Payload
  err := json.Unmarshal(slice, &payload)

  if err != nil {
    panic(err)
  }

  return payload
}
