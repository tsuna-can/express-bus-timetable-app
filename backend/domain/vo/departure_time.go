package vo

import (
	"errors"
	"strconv"
	"strings"
)

type DepartureTime struct {
	value string
}

func NewDepartureTime(value string) (*DepartureTime, error) {

	if !isValidTimeFormat(value) {
		return nil, errors.New("invalid time format: must be HH:MM")
	}
	return &DepartureTime{value: value}, nil
}

func (d DepartureTime) String() string {
	return d.value
}

func isValidTimeFormat(value string) bool {
	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		return false
	}

	hours, err1 := strconv.Atoi(parts[0])
	minutes, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return false
	}

	if hours < 0 || hours > 23 || minutes < 0 || minutes > 59 {
		return false
	}

	return true
}

