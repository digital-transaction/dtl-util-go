// Copyright (c) 2020 Digital Transaction Limited
// All Rights Reserved.
//

package dtl_util

// IsEmptyString return true if input string is zero-length
func IsEmptyString(s string) bool {
	if len(s) <= 0 {
		return true
	}
	return false
}

// IfEmptyString return b if a is empty
func IfEmptyString(a string, b string) string {
	if len(a) > 0 {
		return a
	}
	return b
}

// IsAnyEmptyString return true if any string inside list is zero-length
func IsAnyEmptyString(list ...string) bool {
	for _, s := range list {
		if len(s) <= 0 {
			return true
		}
	}
	return false
}
