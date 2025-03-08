package vo

import (
  "errors"
  "strings"
)

type ParentRouteName struct {
  value string
}

func NewParentRouteName(value string) (*ParentRouteName, error) {
  if strings.TrimSpace(value) == "" {
    return nil, errors.New("parent route name cannot be empty")
  }
  return &ParentRouteName{value: value}, nil
}

func (prn *ParentRouteName) Value() string {
  return string(prn.value)
}

