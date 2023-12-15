package utils

import (
    "strconv"
    "strings"
	"fmt"
)

func ParseValues(valuesStr string) []uint64 {
	var values []uint64
	for _, s := range strings.Split(valuesStr, ",") {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("failed to parse value: %s", s))
		}
		values = append(values, v)
	}
	return values
}
