package helper

import (
	"errors"
	"strconv"
)

func StrToInt64(s string) (int64, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.New("invalid uint64 value")
	}
	return val, nil
}

func StringToPointer(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
