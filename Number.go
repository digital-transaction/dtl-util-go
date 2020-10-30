// Copyright (c) 2020 Digital Transaction Limited
// All Rights Reserved.
//

package dtl_util

import (
	"strconv"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
	return IfThenElseInt(x > y, x, y)
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	return IfThenElseInt(x < y, x, y)
}

// ToIntIfFail converts x from string to int, return def if fail.
func ToIntIfFail(x string, def int) int {
	i, err := strconv.Atoi(x)
	return IfThenElseInt(err == nil, i, def)
}

// ClampInt return value which is >= lowerLimit and <= upperLimit
func ClampInt(value, lowerLimit, upperLimit int) int {
	if value < lowerLimit {
		return lowerLimit
	}
	if value > upperLimit {
		return upperLimit
	}
	return value
}
