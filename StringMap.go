// Copyright (c) 2020 Digital Transaction Limited
// All Rights Reserved.
//

package dtl_util

// AnyFieldEmptyInStringMap ...
func AnyFieldEmptyInStringMap(x map[string]string, fields ...string) bool {
	for _, field := range fields {
		if v, ok := x[field]; !ok || IsEmptyString(v) {
			return true
		}
	}
	return false
}

// MergeStringMap ...
func MergeStringMap(x map[string]string, update map[string]string) map[string]string {
	for k, v := range update {
		if !IsEmptyString(v) {
			x[k] = v
		}
	}
	return x
}
