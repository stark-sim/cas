package tools

import "strconv"

func StringToInt64(str string) int64 {
	parsedInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return parsedInt
}
