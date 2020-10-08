// Copyright (c) 2020 Digital Transaction Limited
// All Rights Reserved.
//

package dtl_util

func BytesToString(x []byte) string {
	if x == nil {
		return ""
	}
	if len(x) == 0 {
		return ""
	}
	return string(x)
}
