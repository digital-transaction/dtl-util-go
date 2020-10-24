// Copyright (c) 2020 Digital Transaction Limited
// All Rights Reserved.
//

package dtl_util

import "strings"

func StringListSplit(s, sep string) []string {
	if s == "" {
		return make([]string, 0)
	}
	return strings.Split(s, sep)
}

func StringListContains(list []string, item string) bool {
	return StringListIndexOf(list, item) != -1
}

func StringListIndexOf(list []string, item string) int {
	if list != nil {
		for index, it := range list {
			if it == item {
				return index
			}
		}
	}
	return -1
}

func StringListRemove(list []string, item string) []string {
	return StringListRemoveAt(list, StringListIndexOf(list, item))
}

func StringListRemoveAt(list []string, index int) []string {
	if list != nil && index >= 0 {
		return append(list[:index], list[index+1:]...)
	}
	return list
}

func StringListHaveCommon(list1 []string, list2 []string) bool {
	for _, x := range list1 {
		for _, y := range list2 {
			if x == y {
				return true
			}
		}
	}
	return false
}
