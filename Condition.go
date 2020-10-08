// Copyright (c) 2020 Digital Transaction Limited
// All Rights Reserved.
//

package dtl_util

// IfThen evaluates a condition, if true returns the parameters otherwise nil
func IfThen(condition bool, a interface{}) interface{} {
	if condition {
		return a
	}
	return nil
}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// IfNil checks if the value is nil, if true returns the default value otherwise the original
func IfNil(value interface{}, def interface{}) interface{} {
	if value != nil {
		return value
	}
	return def
}

// FirstNonNil returns the first non nil parameter
func FirstNonNil(values ...interface{}) interface{} {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}

// IfThenElseString evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElseString(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}

// IfThenElseBytes evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElseBytes(condition bool, a []byte, b []byte) []byte {
	if condition {
		return a
	}
	return b
}

// IfThenElseInt evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElseInt(condition bool, a int, b int) int {
	if condition {
		return a
	}
	return b
}
