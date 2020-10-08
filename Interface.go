// Copyright (c) 2020 Digital Transaction Limited
// All Rights Reserved.
//

package dtl_util

import (
	"fmt"
)

// IsEmpty return true if input interface is zero-length
func IsEmpty(tmp interface{}) bool {
	s := fmt.Sprintf("%v", tmp)
	if len(s) <= 0 {
		return true
	}
	return false
}

// IfEmpty return b if a is empty
func IfEmpty(a interface{}, b interface{}) interface{} {
	s := fmt.Sprintf("%v", a)
	return IfThenElse(len(s) > 0, a, b)
}
