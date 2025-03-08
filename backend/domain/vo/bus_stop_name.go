package vo

import (
  "errors"
  "strings"
)

type BusStopName struct {
  value string
}

func NewBusStopName(value string) (*BusStopName, error) {
  if strings.TrimSpace(value) == "" {
    return nil, errors.New("bus stop name cannot be empty")
  }
  return &BusStopName{value: value}, nil
}


func (bsn *BusStopName) Value() string {
  return string(bsn.value)
}

