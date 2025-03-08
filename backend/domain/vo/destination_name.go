package vo

import (
  "errors"
  "strings"
)

type DestinationName struct {
  value string
}

func NewDestinationName(value string) (*DestinationName, error) {
  if strings.TrimSpace(value) == "" {
    return nil, errors.New("destination name cannot be empty")
  }
  return &DestinationName{value: value}, nil
}

func (dn *DestinationName) Value() string {
  return string(dn.value)
}

