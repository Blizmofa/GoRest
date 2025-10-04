package utils

import "strconv"

func ParseInt64(parameter string) (int64, error) {
	return strconv.ParseInt(parameter, 10, 64)
}
